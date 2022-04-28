package api

import (
	"errors"
	"fmt"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"github.com/jau1jz/cornus/commons/utils"
	"sync"
)

// SignService service layer interface
type SignService interface {
	Sign(info inner.SignRequest) (out inner.SignResponse, code commons.ResponseCode, err error)
	VerifySign(info inner.VerifySignRequest) (out inner.VerifySignResponse, code commons.ResponseCode, err error)
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

	if false == function.StringCheck(info.ApiKey, info.Uri, info.Timestamp, info.Parameter, info.Rand) {
		slog.Slog.ErrorF(nil, "SignServiceImp info not be nil")
		return out, 0, nil
	}
	data := fmt.Sprintf("%s%s%s%s%s", info.ApiKey, info.Uri, info.Timestamp, info.Parameter, info.Rand)

	var thirdPartyPublicKey string
	publicKey, err := s.redis.GetPublicKey(nil, info.ApiKey)
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(nil, "SignServiceImp sign GetPublicKey error %s", err.Error())
		return out, 0, err
	} else if err != nil && err.Error() == redis.Nil.Error() {
		slog.Slog.InfoF(nil, "SignServiceImp sign GetPublicKey error %s", err.Error())

		var thirdPartySystem model.ThirdPartySystem
		err = s.dao.First([]string{model.ThirdPartySystemColumns.ThirdPartyPublicKey}, map[string]interface{}{
			model.ThirdPartySystemColumns.APIkey: info.ApiKey,
		}, nil, &thirdPartySystem)

		if err != nil {
			slog.Slog.ErrorF(nil, "SignServiceImp thirdPartySystem First error %s", err.Error())
			return out, 0, err
		}
		thirdPartyPublicKey = thirdPartySystem.ThirdPartyPublicKey

		err = s.redis.SetPublicKey(nil, inner.PublicKey{
			ApiKey:              info.ApiKey,
			ThirdPartyPublicKey: thirdPartySystem.ThirdPartyPublicKey,
		}, 0)
		if err != nil {
			slog.Slog.ErrorF(nil, "SignServiceImp SetPublicKey error %s", err.Error())
			return out, 0, err
		}
	} else {
		thirdPartyPublicKey = publicKey.ThirdPartyPublicKey
	}
	sign, err := utils.Rsa2Sign([]byte(data), []byte(thirdPartyPublicKey), utils.PKCS_1)
	if err != nil {
		slog.Slog.InfoF(nil, "SignServiceImp Rsa2Sign failed")
		return out, common.ThirdPartySignError, errors.New(commons.GetCodeAndMsg(common.ThirdPartySignError, "english"))
	}
	out.Sign = string(sign)
	return
}

func (s SignServiceImp) VerifySign(info inner.VerifySignRequest) (out inner.VerifySignResponse, code commons.ResponseCode, err error) {

	publicKey, err := s.redis.GetPublicKey(nil, info.ApiKey)
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(nil, "SignServiceImp sign VerifySign error %s", err.Error())
		return out, 0, err
	} else if err != nil && err.Error() == redis.Nil.Error() {
		slog.Slog.InfoF(nil, "SignServiceImp sign VerifySign error %s", err.Error())
		return out, 0, err
	}

	data := fmt.Sprintf("%s%s%s%s%s", info.ApiKey, info.Uri, info.Timestamp, info.Parameter, info.Rand)

	err = utils.Rsa2VerifySign(function.Byte2([]byte(data)), []byte(info.Sign), []byte(publicKey.ThirdPartyPublicKey))
	if err != nil {
		out.Result = false
		slog.Slog.InfoF(nil, "SignServiceImp Verify Rsa2Sign failed")
		return out, common.VerifyThirdPartySignError, errors.New(commons.GetCodeAndMsg(common.VerifyThirdPartySignError, "english"))
	}
	out.Result = true
	return
}
