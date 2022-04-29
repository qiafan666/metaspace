package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"github.com/jau1jz/cornus/commons/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"sync"
	"time"
)

// SignService service layer interface
type SignService interface {
	Sign(info inner.SignRequest) (out inner.SignResponse, code commons.ResponseCode, err error)
	VerifySign(info inner.VerifySignRequest) (out inner.VerifySignResponse, code commons.ResponseCode, err error)
	CreateAuthCode(info request.CreateAuthCode) (out response.CreateAuthCode, code commons.ResponseCode, err error)
	ThirdPartyLogin(info request.ThirdPartyLogin) (out response.ThirdPartyLogin, code commons.ResponseCode, err error)
	GetThirdPartyToken(ctx context.Context, thirdPartyId uint64) (out inner.ThirdPartyToken, err error)
	GetTokenUser(ctx context.Context, token string) (out inner.TokenUser, err error)
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

	var thirdPartyPublicKey string
	publicKey, err := s.redis.GetPublicKey(nil, info.ApiKey)
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(nil, "SignServiceImp sign GetPublicKey error %s", err.Error())
		return out, 0, err
	} else if err != nil && err.Error() == redis.Nil.Error() {
		slog.Slog.InfoF(nil, "SignServiceImp sign GetPublicKey error %s", err.Error())

		var thirdPartySystem model.ThirdPartySystem
		err = s.dao.First([]string{model.ThirdPartySystemColumns.ThirdPartyPublicKey, model.ThirdPartySystemColumns.ID}, map[string]interface{}{
			model.ThirdPartySystemColumns.APIkey: info.ApiKey,
		}, nil, &thirdPartySystem)

		if err != nil {
			slog.Slog.ErrorF(nil, "SignServiceImp thirdPartySystem First error %s", err.Error())
			return out, 0, err
		}
		thirdPartyPublicKey = thirdPartySystem.ThirdPartyPublicKey

		err = s.redis.SetPublicKey(nil, inner.PublicKey{
			Id:                  thirdPartySystem.ID,
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

	data := fmt.Sprintf("%s%s%s%s%s", info.ApiKey, info.Uri, info.Timestamp, info.Parameter, info.Rand)
	bufferString := bytes.NewBufferString(data)
	sign, err := utils.Rsa2Sign(bufferString.Bytes(), []byte(thirdPartyPublicKey), utils.PKCS_8)
	if err != nil {
		slog.Slog.InfoF(nil, "SignServiceImp Rsa2Sign failed")
		return out, common.ThirdPartySignError, errors.New(commons.GetCodeAndMsg(common.ThirdPartySignError, commons.DefualtLanguage))
	}
	out.Sign = string(sign)
	return
}

func (s SignServiceImp) VerifySign(info inner.VerifySignRequest) (out inner.VerifySignResponse, code commons.ResponseCode, err error) {

	//check time
	tm := time.Unix(info.Timestamp, 0)
	nowTime := time.Now()
	if nowTime.Sub(tm) > time.Second*30 {
		slog.Slog.InfoF(nil, "SignServiceImp sign VerifySign error %s", err.Error())
		return out, common.VerifyThirdPartySignTimeOut, errors.New(commons.GetCodeAndMsg(common.VerifyThirdPartySignTimeOut, commons.DefualtLanguage))
	}

	//check rand
	_, err = s.redis.GetRand(nil, info.ApiKey)
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(nil, "SignServiceImp sign VerifySign error %s", err.Error())
		return out, 0, err
	} else if err != nil && err.Error() == redis.Nil.Error() {

		err = s.redis.SetRand(nil, inner.Rand{
			ApiKey: info.ApiKey,
			Rand:   info.Rand,
		}, time.Second*30)
		if err != nil {
			slog.Slog.ErrorF(nil, "SignServiceImp SetRand error %s", err.Error())
			return out, 0, err
		}
	} else {
		slog.Slog.InfoF(nil, "SignServiceImp frequent VerifySign error %s", nil)
		return out, common.FrequentVerifyThirdPartySign, errors.New(commons.GetCodeAndMsg(common.FrequentVerifyThirdPartySign, commons.DefualtLanguage))
	}

	var thirdPartyPublicKey string
	var third_party_id uint64
	publicKey, err := s.redis.GetPublicKey(nil, info.ApiKey)
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(nil, "SignServiceImp sign VerifySign error %s", err.Error())
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
		third_party_id = publicKey.Id
	} else {
		thirdPartyPublicKey = publicKey.ThirdPartyPublicKey
		third_party_id = publicKey.Id
	}

	data := fmt.Sprintf("%s%s%s%s%s", info.ApiKey, info.Uri, info.Timestamp, info.Parameter, info.Rand)

	bufferString := bytes.NewBufferString(data)
	err = utils.Rsa2VerifySign(info.Sign, bufferString.Bytes(), []byte(thirdPartyPublicKey))
	if err != nil {
		out.Flag = false
		slog.Slog.InfoF(nil, "SignServiceImp Verify Rsa2Sign failed")
		return out, common.VerifyThirdPartySignError, errors.New(commons.GetCodeAndMsg(common.VerifyThirdPartySignError, commons.DefualtLanguage))
	}
	out.ThirdPartyId = third_party_id
	out.Flag = true
	return
}

func (s SignServiceImp) CreateAuthCode(info request.CreateAuthCode) (out response.CreateAuthCode, code commons.ResponseCode, err error) {

	uuid := utils.GenerateUUID()
	err = s.redis.SetAuthCode(info.Ctx, inner.AuthCode{
		ThirdPartyPublicId: strconv.FormatUint(info.ThirdPartyId, 10),
		AuthCode:           uuid,
	}, time.Minute*3)
	if err != nil {
		slog.Slog.ErrorF(nil, "SignServiceImp SetRand error %s", err.Error())
		return out, 0, err
	}

	out.AuthCode = uuid
	return

}

func (s SignServiceImp) ThirdPartyLogin(info request.ThirdPartyLogin) (out response.ThirdPartyLogin, code commons.ResponseCode, err error) {

	_, err = s.redis.GetAuthCode(info.Ctx, strconv.FormatUint(info.ThirdPartyId, 10))
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(info.Ctx, "SignServiceImp ThirdPartyLogin error %s", err.Error())
		return out, 0, err
	} else if err != nil && err.Error() == redis.Nil.Error() {

		slog.Slog.InfoF(info.Ctx, "SignServiceImp ThirdPartyLogin auth_code is expired  error %s", err.Error())
		return out, common.AuthCodeAlreadyExpired, errors.New(commons.GetCodeAndMsg(common.AuthCodeAlreadyExpired, commons.DefualtLanguage))
	} else {

		var user model.User

		vWalletAddress := strings.ToLower(info.Account)

		if info.Type == common.LoginTypeWallet {
			////check sign hex add hex prefix
			if strings.HasPrefix(info.Password, "0x") == false {
				info.Password = "0x" + info.Password
			}

			//check sign len
			nonce, err := s.redis.GetNonce(info.Ctx, vWalletAddress)
			if err != nil && err.Error() != redis.Nil.Error() {
				slog.Slog.InfoF(info.Ctx, "SignServiceImp sign GetNonce error %s", err.Error())
				return out, 0, err
			} else if err != nil && err.Error() == redis.Nil.Error() {
				slog.Slog.InfoF(info.Ctx, "SignServiceImp sign GetNonce error %s", err.Error())
				return out, common.NonceExpireOrNull, err
			}
			if err = function.VerifySig(vWalletAddress, info.Password, nonce.Nonce); err != nil && common.DebugFlag == false {
				slog.Slog.InfoF(info.Ctx, "SignServiceImp sign verify error %s", err.Error())
				return out, common.SignatureVerificationError, err
			}
			if err = s.redis.DelNonce(info.Ctx, user.UUID); err != nil {
				slog.Slog.InfoF(info.Ctx, "SignServiceImp DelNonce error %s", err.Error())
				return out, 0, err
			}
			//if wallet address does not register,then register
			err = s.dao.First([]string{model.UserColumns.UUID}, map[string]interface{}{
				model.UserColumns.WalletAddress: vWalletAddress,
			}, nil, &user)
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
				slog.Slog.ErrorF(info.Ctx, "SignServiceImp First error %s", err.Error())
				return out, 0, err
			} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
				//register
				user = model.User{
					UUID:          utils.GenerateUUID(),
					WalletAddress: vWalletAddress,
				}
				if err := s.dao.Create(&user); err != nil {
					slog.Slog.InfoF(info.Ctx, "SignServiceImp Create error %s", err.Error())
					return out, 0, err
				}
			}
		} else {
			var AccountType string
			if info.Type == common.LoginTypeEmail {
				AccountType = model.UserColumns.Email
			}
			err = s.dao.WithContext(info.Ctx).First([]string{model.UserColumns.UUID, model.UserColumns.Email}, map[string]interface{}{
				AccountType:                info.Account,
				model.UserColumns.Password: utils.StringToSha256(info.Password),
			}, nil, &user)

			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Slog.ErrorF(info.Ctx, "SignServiceImp Login Count error %s", err.Error())
				return out, 0, err
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Slog.InfoF(info.Ctx, "SignServiceImp Register account or password error")
				return out, common.PasswordOrAccountError, errors.New(commons.GetCodeAndMsg(common.PasswordOrAccountError, info.Language))
			}
		}

		token := utils.GenerateUUID()
		err = s.redis.SetThirdPartyToken(info.Ctx, inner.ThirdPartyToken{
			ThirdPartyPublicId: strconv.FormatUint(info.ThirdPartyId, 10),
			Token:              token,
		}, time.Second*30)
		if err != nil {
			slog.Slog.ErrorF(nil, "SignServiceImp SetThirdPartyToken error %s", err.Error())
			return out, 0, err
		}

		err = s.redis.SetTokenUser(info.Ctx, inner.TokenUser{
			Token:  token,
			Uuid:   user.UUID,
			UserId: user.ID,
			Email:  user.Email,
		}, time.Second*30)
		if err != nil {
			slog.Slog.ErrorF(nil, "SignServiceImp SetTokenUser error %s", err.Error())
			return out, 0, err
		}
		out.JwtToken = token
	}

	return
}

func (s SignServiceImp) GetThirdPartyToken(ctx context.Context, thirdPartyId uint64) (out inner.ThirdPartyToken, err error) {

	result, err := s.redis.GetThirdPartyToken(ctx, strconv.FormatUint(thirdPartyId, 10))
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.ErrorF(nil, "SignServiceImp GetThirdPartyToken error %s", err.Error())
		return
	}
	out = result
	return

}

func (s SignServiceImp) GetTokenUser(ctx context.Context, token string) (out inner.TokenUser, err error) {

	result, err := s.redis.GetTokenUser(ctx, token)
	if err != nil {
		slog.Slog.ErrorF(nil, "SignServiceImp GetThirdPartyToken error %s", err.Error())
		return
	}
	out = result
	return

}
