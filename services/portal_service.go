package bizservice

import (
	"errors"
	"github.com/blockfishio/metaspace-backend/common/function"
	"strings"
	"sync"
	"time"

	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/jau1jz/cornus"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"github.com/jau1jz/cornus/commons/utils"
	"gorm.io/gorm"
)

// PortalService service layer interface
type PortalService interface {
	//Login support email and wallet login api
	Login(info request.UserLogin) (out response.UserLogin, code commons.ResponseCode, err error)
	//GetNonce client get new nonce from server
	GetNonce(info request.GetNonce) (out response.GetNonce, code commons.ResponseCode, err error)
	Register(info request.RegisterUser) (out response.RegisterUser, code commons.ResponseCode, err error)
	UpdatePassword(info request.PasswordUpdate) (out response.PasswordUpdate, code commons.ResponseCode, err error)
	SubscribeNewsletterEmail(info request.SubscribeNewsletterEmail) (out response.SubscribeNewsletterEmail, code commons.ResponseCode, err error)
}

var jwtConfig struct {
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&jwtConfig)
}

var portalServiceIns *portalServiceImp
var portalServiceInitOnce sync.Once

func NewPortalServiceInstance() PortalService {
	portalServiceInitOnce.Do(func() {
		portalServiceIns = &portalServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return portalServiceIns
}

type portalServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (p portalServiceImp) GetNonce(info request.GetNonce) (out response.GetNonce, code commons.ResponseCode, err error) {
	nonce := utils.GenerateUUID()
	err = p.redis.SetNonce(info.Ctx, inner.Nonce{
		Address: info.Address,
		Nonce:   nonce,
	}, time.Minute*5)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp SetNonce error %s", err.Error())
		return out, 0, err
	}
	out.Nonce = nonce
	return
}

func (p portalServiceImp) UpdatePassword(info request.PasswordUpdate) (out response.PasswordUpdate, code commons.ResponseCode, err error) {
	var user model.User
	tx := p.dao.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	err = tx.First([]string{model.UserColumns.Password}, map[string]interface{}{
		model.UserColumns.UUID: info.BaseUUID,
	}, func(db *gorm.DB) *gorm.DB {
		return db.Set("gorm:query_option", "FOR UPDATE")
	}, &user)

	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp UpdatePassword First error %s", err.Error())
		return out, 0, err
	}

	if utils.StringToSha256(info.OldPassword) != user.Password {
		slog.Slog.InfoF(info.Ctx, "portalServiceImp UpdatePassword old_password not equal")
		return out, common.OldPasswordNotEqual, errors.New(commons.GetCodeAndMsg(common.OldPasswordNotEqual, info.Language))
	}

	if info.OldPassword == info.NewPassword {
		slog.Slog.InfoF(info.Ctx, "portalServiceImp UpdatePassword old_password equal new password")
		return out, common.OldPasswordEqualNewPassword, errors.New(commons.GetCodeAndMsg(common.OldPasswordEqualNewPassword, info.Language))
	}

	_, err = tx.WithContext(info.Ctx).Update(model.User{
		Password: utils.StringToSha256(info.NewPassword),
	}, map[string]interface{}{
		model.UserColumns.UUID: info.BaseUUID,
	}, nil)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp UpdatePassword Update error %s", err.Error())
		return out, 0, err
	}
	return
}

func (p portalServiceImp) Register(info request.RegisterUser) (out response.RegisterUser, code commons.ResponseCode, err error) {
	//check account exists
	count, err := p.dao.Count(model.User{}, map[string]interface{}{
		model.UserColumns.Email: info.Email,
	}, nil)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp Register Count error %s", err.Error())
		return out, 0, err
	}
	if count > 0 {
		slog.Slog.InfoF(info.Ctx, "portalServiceImp Register account already exists")
		return out, common.AccountAlreadyExists, errors.New(commons.GetCodeAndMsg(common.AccountAlreadyExists, info.Language))
	}
	user := model.User{
		UUID:     utils.GenerateUUID(),
		Email:    info.Email,
		Password: utils.StringToSha256(info.Password),
	}
	err = p.dao.WithContext(info.Ctx).Create(&user)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp Register Create error %s", err.Error())
		return out, 0, err
	}
	out = response.RegisterUser{}
	return
}
func (p portalServiceImp) Login(info request.UserLogin) (out response.UserLogin, code commons.ResponseCode, err error) {
	var user model.User

	if info.Type == common.LoginTypeWallet {
		////check sign hex add hex prefix
		if strings.HasPrefix(info.Password, "0x") == false {
			info.Password = "0x" + info.Password
		}

		//check sign len
		nonce, err := p.redis.GetNonce(info.Ctx, info.Account)
		if err != nil && err.Error() != redis.Nil.Error() {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp sign GetNonce error %s", err.Error())
			return out, 0, err
		} else if err != nil && err.Error() == redis.Nil.Error() {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp sign GetNonce error %s", err.Error())
			return out, common.NonceExpireOrNull, err
		}
		if err = function.VerifySig(info.Account, info.Password, nonce.Nonce); err != nil && common.DebugFlag == false {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp sign verify error %s", err.Error())
			return out, common.SignatureVerificationError, err
		}
		if err = p.redis.DelNonce(info.Ctx, user.UUID); err != nil {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp DelNonce error %s", err.Error())
			return out, 0, err
		}
		//if wallet address does not register,then register
		err = p.dao.First([]string{model.UserColumns.UUID}, map[string]interface{}{
			model.UserColumns.WalletAddress: info.Account,
		}, nil, &user)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp First error %s", err.Error())
			return out, 0, err
		} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
			//register
			user = model.User{
				UUID:          utils.GenerateUUID(),
				WalletAddress: info.Account,
			}
			if err := p.dao.Create(&user); err != nil {
				slog.Slog.InfoF(info.Ctx, "portalServiceImp Create error %s", err.Error())
				return out, 0, err
			}
		}
	} else {
		var AccountType string
		if info.Type == common.LoginTypeEmail {
			AccountType = model.UserColumns.Email
		}
		err = p.dao.WithContext(info.Ctx).First([]string{model.UserColumns.UUID, model.UserColumns.Email}, map[string]interface{}{
			AccountType:                info.Account,
			model.UserColumns.Password: utils.StringToSha256(info.Password),
		}, nil, &user)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp Login Count error %s", err.Error())
			return out, 0, err
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp Register account or password error")
			return out, common.PasswordOrAccountError, errors.New(commons.GetCodeAndMsg(common.PasswordOrAccountError, info.Language))
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"uuid":  user.UUID,
		"iss":   "metaspace",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})
	signedString, err := token.SignedString([]byte(jwtConfig.JWT.Secret))
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp Login SignedString error %s", err.Error())
		return out, 0, err
	}
	out.JwtToken = signedString
	return
}
func (p portalServiceImp) SubscribeNewsletterEmail(info request.SubscribeNewsletterEmail) (out response.SubscribeNewsletterEmail, code commons.ResponseCode, err error) {

	var subscribeNewsletterEmail model.SubscribeNewsletterEmail
	err = p.dao.First([]string{model.SubscribeNewsletterEmailColumns.Email}, map[string]interface{}{
		model.SubscribeNewsletterEmailColumns.Email: info.Email,
	}, nil, &subscribeNewsletterEmail)

	if errors.Is(err, gorm.ErrRecordNotFound) {

		email := model.SubscribeNewsletterEmail{
			Email:       info.Email,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
			Status:      1,
		}
		err = p.dao.WithContext(info.Ctx).Create(&email)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp SubscribeNewsletterEmail Create error %s", err.Error())
			return out, 0, err
		}

		out = response.SubscribeNewsletterEmail{}
		return
	}

	slog.Slog.ErrorF(info.Ctx, "portalServiceImp SubscribeNewsletterEmail find first email error")
	return out, 0, errors.New("portalServiceImp SubscribeNewsletterEmail find first email")

}
