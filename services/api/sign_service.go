package api

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/qiafan666/metaspace/common"
	"github.com/qiafan666/metaspace/dao"
	"github.com/qiafan666/metaspace/model"
	"github.com/qiafan666/metaspace/pojo/inner"
	"github.com/qiafan666/metaspace/redis"
	"github.com/qiafan666/quickweb/commons"
	slog "github.com/qiafan666/quickweb/commons/log"
	"github.com/qiafan666/quickweb/commons/utils"
	"strconv"
	"sync"
	"time"
)

// SignService service layer interface
type SignService interface {
	Sign(info inner.SignRequest) (out inner.SignResponse, code commons.ResponseCode, err error)
	VerifySign(info inner.VerifySignRequest) (out inner.VerifySignResponse, code commons.ResponseCode, err error)
	GetTokenUser(ctx context.Context, token string) (out inner.TokenUser, code commons.ResponseCode, err error)
}

var SignServiceIns *SignServiceImp
var signServiceInitOnce sync.Once

func NewSignInstance() SignService {
	signServiceInitOnce.Do(func() {
		SignServiceIns = &SignServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return SignServiceIns
}

type SignServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (s SignServiceImp) Sign(info inner.SignRequest) (out inner.SignResponse, code commons.ResponseCode, err error) {

	ctx := context.Background()
	var thirdPartyPublicKey string
	publicKey, err := s.redis.GetPublicKey(ctx, info.ApiKey)
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(ctx, "SignServiceImp sign GetPublicKey error %s", err.Error())
		return out, 0, err
	} else if err != nil && err.Error() == redis.Nil.Error() {
		slog.Slog.InfoF(ctx, "SignServiceImp sign GetPublicKey error %s", err.Error())

		var thirdPartySystem model.ThirdPartySystem
		err = s.dao.First([]string{model.ThirdPartySystemColumns.ThirdPartyPublicKey, model.ThirdPartySystemColumns.ID}, map[string]interface{}{
			model.ThirdPartySystemColumns.APIkey: info.ApiKey,
		}, nil, &thirdPartySystem)

		if err != nil {
			slog.Slog.ErrorF(ctx, "SignServiceImp thirdPartySystem First error %s", err.Error())
			return out, 0, err
		}
		thirdPartyPublicKey = thirdPartySystem.ThirdPartyPublicKey

		err = s.redis.SetPublicKey(ctx, inner.PublicKey{
			Id:                  thirdPartySystem.ID,
			ApiKey:              info.ApiKey,
			ThirdPartyPublicKey: thirdPartySystem.ThirdPartyPublicKey,
		}, 0)
		if err != nil {
			slog.Slog.ErrorF(ctx, "SignServiceImp SetPublicKey error %s", err.Error())
			return out, 0, err
		}
	} else {
		thirdPartyPublicKey = publicKey.ThirdPartyPublicKey
	}

	data := fmt.Sprintf("%s%s%s%s%s", info.ApiKey, info.Uri, info.Timestamp, info.Parameter, info.Rand)

	bufferString := bytes.NewBufferString(data)
	sign, err := utils.Rsa2Sign(bufferString.Bytes(), []byte(thirdPartyPublicKey), utils.PKCS_8)
	if err != nil {
		slog.Slog.InfoF(ctx, "SignServiceImp Rsa2Sign failed")
		return out, common.ThirdPartySignError, errors.New(commons.GetCodeAndMsg(common.ThirdPartySignError, commons.DefualtLanguage))
	}

	out.Sign = hex.EncodeToString(sign)
	return
}

func (s SignServiceImp) VerifySign(info inner.VerifySignRequest) (out inner.VerifySignResponse, code commons.ResponseCode, err error) {

	//check time
	parseInt, err := strconv.ParseInt(info.Timestamp, 10, 64)
	if err != nil {
		slog.Slog.InfoF(nil, "SignServiceImp VerifySign ParseInt failed")
		return out, 0, err
	}
	tm := time.Unix(parseInt, 0)
	if tm.Add(time.Second*30).Before(time.Now()) && common.DebugFlag == false {
		slog.Slog.InfoF(nil, "SignServiceImp sign VerifySign error %s", err.Error())
		return out, common.VerifyThirdPartySignTimeOut, errors.New(commons.GetCodeAndMsg(common.VerifyThirdPartySignTimeOut, commons.DefualtLanguage))
	}

	//check rand
	ctx := context.Background()
	_, err = s.redis.GetRand(ctx, info.Rand)
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(ctx, "SignServiceImp sign VerifySign error %s", err.Error())
		return out, 0, err
	} else if err != nil && err.Error() == redis.Nil.Error() {

		err = s.redis.SetRand(ctx, inner.Rand{
			Rand: info.Rand,
		}, time.Second*10)
		if err != nil {
			slog.Slog.ErrorF(ctx, "SignServiceImp SetRand error %s", err.Error())
			return out, 0, err
		}
	} else {
		slog.Slog.InfoF(ctx, "SignServiceImp frequent VerifySign error")
		return out, common.FrequentVerifyThirdPartySign, errors.New(commons.GetCodeAndMsg(common.FrequentVerifyThirdPartySign, commons.DefualtLanguage))
	}

	var thirdPartyPublicKey string
	var thirdPartyId uint64
	publicKey, err := s.redis.GetPublicKey(ctx, info.ApiKey)
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(ctx, "SignServiceImp sign VerifySign error %s", err.Error())
		return out, 0, err
	} else if err != nil && err.Error() == redis.Nil.Error() {

		var thirdPartySystem model.ThirdPartySystem
		err = s.dao.First([]string{model.ThirdPartySystemColumns.ThirdPartyPublicKey, model.ThirdPartySystemColumns.ID}, map[string]interface{}{
			model.ThirdPartySystemColumns.APIkey: info.ApiKey,
		}, nil, &thirdPartySystem)

		if err != nil {
			slog.Slog.ErrorF(nil, "SignServiceImp VerifySign thirdPartySystem First error %s", err.Error())
			return out, 0, err
		}
		thirdPartyPublicKey = thirdPartySystem.ThirdPartyPublicKey
		thirdPartyId = publicKey.Id

		err = s.redis.SetPublicKey(ctx, inner.PublicKey{
			Id:                  thirdPartySystem.ID,
			ApiKey:              info.ApiKey,
			ThirdPartyPublicKey: thirdPartySystem.ThirdPartyPublicKey,
		}, 0)
		if err != nil {
			slog.Slog.ErrorF(ctx, "SignServiceImp SetPublicKey error %s", err.Error())
			return out, 0, err
		}
	} else {
		thirdPartyPublicKey = publicKey.ThirdPartyPublicKey
		thirdPartyId = publicKey.Id
	}

	bufferString := bytes.Buffer{}
	bufferString.WriteString(info.ApiKey)
	bufferString.WriteString(info.Uri)
	bufferString.WriteString(info.Timestamp)
	bufferString.Write(info.Parameter)
	bufferString.WriteString(info.Rand)

	decodeString, err := base64.URLEncoding.DecodeString(info.Sign)
	if err != nil {
		slog.Slog.ErrorF(ctx, "SignServiceImp VerifySign Sign DecodeString error %s", err.Error())
		return out, 0, err
	}
	thirdPartyPublicKeyBufferString := bytes.NewBufferString(thirdPartyPublicKey)
	err = utils.Rsa2VerifySign(sha256.Sum256(bufferString.Bytes()), decodeString, thirdPartyPublicKeyBufferString.Bytes())
	if err != nil {
		slog.Slog.InfoF(ctx, "SignServiceImp Verify Rsa2Sign failed %s", err.Error())
		return out, common.VerifyThirdPartySignError, errors.New(commons.GetCodeAndMsg(common.VerifyThirdPartySignError, commons.DefualtLanguage))
	}
	out.ThirdPartyId = thirdPartyId
	return
}

func (s SignServiceImp) GetTokenUser(ctx context.Context, token string) (out inner.TokenUser, code commons.ResponseCode, err error) {

	result, err := s.redis.GetTokenUser(ctx, token)
	if err != nil {
		slog.Slog.ErrorF(nil, "SignServiceImp GetThirdPartyToken error %s", err.Error())
		return
	}
	out = result
	return

}
