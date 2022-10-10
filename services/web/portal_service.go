package web

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blockfishio/metaspace-backend/contract/eth/eth_mint"
	"github.com/blockfishio/metaspace-backend/grpc"
	"github.com/blockfishio/metaspace-backend/grpc/proto"
	"github.com/blockfishio/metaspace-backend/model/join"
	"github.com/blockfishio/metaspace-backend/services/exchange"
	"github.com/blockfishio/metaspace-backend/thirdparty"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/blockfishio/metaspace-backend/common/function"

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
	ThirdPartyLogin(info request.ThirdPartyLogin) (out response.ThirdPartyLogin, code commons.ResponseCode, err error)
	//Login support email and wallet login api
	Login(info request.UserLogin) (out response.UserLogin, code commons.ResponseCode, err error)
	//GetNonce client get new nonce from server
	GetNonce(info request.GetNonce) (out response.GetNonce, code commons.ResponseCode, err error)
	Register(info request.RegisterUser) (out response.RegisterUser, code commons.ResponseCode, err error)
	UpdatePassword(info request.PasswordUpdate) (out response.PasswordUpdate, code commons.ResponseCode, err error)
	SubscribeNewsletterEmail(info request.SubscribeNewsletterEmail) (out response.SubscribeNewsletterEmail, code commons.ResponseCode, err error)
	GetTowerStatus(info request.TowerStats) (out response.TowerStats, code commons.ResponseCode, err error)
	GetSign(info request.Sign) (out response.Sign, code commons.ResponseCode, err error)
	UserUpdate(info request.UserUpdate) (out response.UserUpdate, code commons.ResponseCode, err error)
	UserHistory(info request.UserHistory) (out response.UserHistory, code commons.ResponseCode, err error)
	ExchangePrice(info request.ExchangePrice) (out response.ExchangePrice, code commons.ResponseCode, err error)
	AssetDetail(info request.AssetDetail) (out response.AssetDetail, code commons.ResponseCode, err error)
	GameCurrency(info request.GameCurrency) (out response.GameCurrency, code commons.ResponseCode, err error)
	SendCode(info request.SendCode) (out response.SendCode, code commons.ResponseCode, err error)
	PaperMint(info request.PaperMint) (out response.PaperMint, code commons.ResponseCode, err error)
	PaperTransaction(info request.PaperTransaction) (out response.PaperTransaction, code commons.ResponseCode, err error)
}

var portalConfig struct {
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	ETHContract struct {
		Client string `yaml:"monitor_client"`
		Mint   string `yaml:"mint"`
		Assets string `yaml:"assets"`
		Ship   string `yaml:"ship"`
		Market string `yaml:"market"`
	} `yaml:"eth_contract"`
	BSCContract struct {
		Client string `yaml:"monitor_client"`
		Mint   string `yaml:"mint"`
		Assets string `yaml:"assets"`
		Ship   string `yaml:"ship"`
		Market string `yaml:"market"`
	} `yaml:"bsc_contract"`
	Chain struct {
		ETH uint8 `yaml:"eth"`
		BSC uint8 `yaml:"bsc"`
	} `yaml:"chain"`
	Paper struct {
		Payment struct {
			Url              string `yaml:"url"`
			MintContractId   string `yaml:"mint_contract_id"`
			MarketContractId string `yaml:"market_contract_id"`
			Authorization    string `yaml:"authorization"`
			Currency         string `yaml:"currency"`
		} `yaml:"payment"`
	} `yaml:"paper"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&portalConfig)
}

var portalServiceIns *portalServiceImp
var portalServiceInitOnce sync.Once
var exchangeService exchange.Exchange

func NewPortalServiceInstance() PortalService {

	portalServiceInitOnce.Do(func() {
		portalServiceIns = &portalServiceImp{
			dao:               dao.Instance(),
			redis:             redis.Instance(),
			thirdPartyService: thirdparty.NewThirdPartyService(),
		}
	})

	exchangeService = exchange.NewExchangeService()
	return portalServiceIns
}

type portalServiceImp struct {
	dao               dao.Dao
	redis             redis.Dao
	thirdPartyService thirdparty.Service
}

func (p portalServiceImp) ThirdPartyLogin(info request.ThirdPartyLogin) (out response.ThirdPartyLogin, code commons.ResponseCode, err error) {

	authCode, err := p.redis.GetAuthCode(info.Ctx, info.AuthCode)
	if err != nil && err.Error() != redis.Nil.Error() {
		slog.Slog.InfoF(info.Ctx, "portalServiceImp ThirdPartyLogin error %s", err.Error())
		return out, 0, err
	} else if err != nil && err.Error() == redis.Nil.Error() {
		slog.Slog.InfoF(info.Ctx, "portalServiceImp ThirdPartyLogin auth_code is expired  error %s", err.Error())
		return out, common.AuthCodeAlreadyExpired, errors.New(commons.GetCodeAndMsg(common.AuthCodeAlreadyExpired, commons.DefualtLanguage))
	} else {

		err = p.redis.DelAuthCode(info.Ctx, info.AuthCode)
		if err != nil {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp DelAuthCode error %s", err.Error())
			return out, 0, err
		}

		var user model.User

		vWalletAddress := strings.ToLower(info.Account)

		if info.Type == common.LoginTypeWallet {
			////check sign hex add hex prefix
			if strings.HasPrefix(info.Password, "0x") == false {
				info.Password = "0x" + info.Password
			}

			//check sign len
			nonce, err := p.redis.GetNonce(info.Ctx, vWalletAddress)
			if err != nil && err.Error() != redis.Nil.Error() {
				slog.Slog.InfoF(info.Ctx, "portalServiceImp sign GetNonce error %s", err.Error())
				return out, 0, err
			} else if err != nil && err.Error() == redis.Nil.Error() {
				slog.Slog.InfoF(info.Ctx, "portalServiceImp sign GetNonce error %s", err.Error())
				return out, common.NonceExpireOrNull, err
			}
			if err = function.VerifySig(vWalletAddress, info.Password, nonce.Nonce); err != nil && common.DebugFlag == false {
				slog.Slog.InfoF(info.Ctx, "portalServiceImp sign verify error %s", err.Error())
				return out, common.SignatureVerificationError, err
			}
			if err = p.redis.DelNonce(info.Ctx, user.UUID); err != nil {
				slog.Slog.InfoF(info.Ctx, "portalServiceImp DelNonce error %s", err.Error())
				return out, 0, err
			}
			//if wallet address does not register,then register
			err = p.dao.First([]string{model.UserColumns.UUID}, map[string]interface{}{
				model.UserColumns.WalletAddress: vWalletAddress,
			}, nil, &user)
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
				slog.Slog.ErrorF(info.Ctx, "portalServiceImp First error %s", err.Error())
				return out, 0, err
			} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
				//register
				user = model.User{
					UUID:          utils.GenerateUUID(),
					WalletAddress: vWalletAddress,
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

		var token string
		token = utils.GenerateUUID()
		err = p.redis.SetTokenUser(info.Ctx, inner.TokenUser{
			ThirdPartyPublicId: authCode.ThirdPartyPublicId,
			Token:              utils.GenerateUUID(),
			Uuid:               user.UUID,
			UserId:             user.ID,
			Email:              user.Email,
		}, time.Second*30)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp SetTokenUser error %s", err.Error())
			return out, 0, err
		}
		//del userToken
		err = p.redis.DelUserToken(info.Ctx, strconv.FormatUint(user.ID, 10))
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp DelUserToken error %s", err.Error())
			return out, 0, err
		}
		//set userToken
		err = p.redis.SetUserToken(info.Ctx, inner.UserToken{
			Token:  token,
			UserId: strconv.FormatUint(user.ID, 10),
		})
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp SetUserToken error %s", err.Error())
			return out, 0, err
		}
		out.ExpireTime = time.Now().Add(time.Second * 30)
		out.Token = token
		out.Uuid = user.UUID
		out.WalletAddress = user.WalletAddress
		out.Email = user.Email

	}

	var thirdPartySystem model.ThirdPartySystem
	err = p.dao.First([]string{model.ThirdPartySystemColumns.ThirdPartyPublicKey, model.ThirdPartySystemColumns.CallbackAddress, model.ThirdPartySystemColumns.ID}, map[string]interface{}{
		model.ThirdPartySystemColumns.ID: authCode.ThirdPartyPublicId,
	}, nil, &thirdPartySystem)

	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp ThirdPartyLogin thirdPartySystem First error %s", err.Error())
		return out, 0, err
	}

	out.Url = dataCallBack(out, common.UrlCallbackLogin, thirdPartySystem)

	if len(out.Url) == 0 {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp thirdSign dataCallBack error %s", err.Error())
		return out, 0, err
	}
	return
}

func dataCallBack(out interface{}, enumeration string, thirdPartySystem model.ThirdPartySystem) string {
	marshal, err := json.Marshal(out)
	if err != nil {
		return ""
	}
	publicKeyBuffer := bytes.NewBufferString(thirdPartySystem.ThirdPartyPublicKey)
	encrypt, err := utils.RsaEncrypt(marshal, publicKeyBuffer.Bytes())
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s%s?value=%s", thirdPartySystem.CallbackAddress, enumeration, base64.URLEncoding.EncodeToString(encrypt))
}

func (p portalServiceImp) GetNonce(info request.GetNonce) (out response.GetNonce, code commons.ResponseCode, err error) {
	nonce := utils.GenerateUUID()
	err = p.redis.SetNonce(info.Ctx, inner.Nonce{
		Address: strings.ToLower(info.Address),
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
		Password:  utils.StringToSha256(info.NewPassword),
		UpdatedAt: time.Now(),
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
		UUID:          utils.GenerateUUID(),
		Email:         info.Email,
		Password:      utils.StringToSha256(info.Password),
		AvatarAddress: common.DefaultAvatar,
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

	vWalletAddress := strings.ToLower(info.Account)

	if info.Type == common.LoginTypeWallet {
		////check sign hex add hex prefix
		if strings.HasPrefix(info.Password, "0x") == false {
			info.Password = "0x" + info.Password
		}

		//check sign len
		nonce, err := p.redis.GetNonce(info.Ctx, vWalletAddress)
		if err != nil && err.Error() != redis.Nil.Error() {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp sign GetNonce error %s", err.Error())
			return out, 0, err
		} else if err != nil && err.Error() == redis.Nil.Error() {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp sign GetNonce error %s", err.Error())
			return out, common.NonceExpireOrNull, err
		}
		if err = function.VerifySig(vWalletAddress, info.Password, nonce.Nonce); err != nil && common.DebugFlag == false {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp sign verify error %s", err.Error())
			return out, common.SignatureVerificationError, err
		}
		if err = p.redis.DelNonce(info.Ctx, user.UUID); err != nil {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp DelNonce error %s", err.Error())
			return out, 0, err
		}
		//if wallet address does not register,then register
		err = p.dao.First([]string{model.UserColumns.UUID, model.UserColumns.Email, model.UserColumns.UserName, model.UserColumns.AvatarAddress}, map[string]interface{}{
			model.UserColumns.WalletAddress: vWalletAddress,
		}, nil, &user)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp First error %s", err.Error())
			return out, 0, err
		} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
			//register
			userName := "User" + strconv.FormatInt(time.Now().Unix(), 10)
			user = model.User{
				UUID:          utils.GenerateUUID(),
				WalletAddress: vWalletAddress,
				UserName:      userName,
				AvatarAddress: common.DefaultAvatar,
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

		emailCode, err := p.redis.GetEmailCode(info.Ctx, info.Account)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetEmailCode error %s", err.Error())
			return out, 0, err
		}
		if emailCode.Code != info.Password {
			slog.Slog.InfoF(info.Ctx, "portalServiceImp Register account or password error")
			return out, common.PasswordOrAccountError, errors.New(commons.GetCodeAndMsg(common.PasswordOrAccountError, info.Language))
		}

		err = p.dao.WithContext(info.Ctx).First([]string{model.UserColumns.UUID, model.UserColumns.Email}, map[string]interface{}{
			AccountType: info.Account,
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
	signedString, err := token.SignedString([]byte(portalConfig.JWT.Secret))
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp Login SignedString error %s", err.Error())
		return out, 0, err
	}
	out.JwtToken = signedString
	out.AvatarAddress = user.AvatarAddress
	out.UserName = user.UserName
	return
}
func (p portalServiceImp) SubscribeNewsletterEmail(info request.SubscribeNewsletterEmail) (out response.SubscribeNewsletterEmail, code commons.ResponseCode, err error) {

	var subscribeNewsletterEmail model.SubscribeNewsletterEmail
	err = p.dao.First([]string{model.SubscribeNewsletterEmailColumns.Email}, map[string]interface{}{
		model.SubscribeNewsletterEmailColumns.Email: info.Email,
	}, nil, &subscribeNewsletterEmail)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp SubscribeNewsletterEmail error", err.Error())
		return out, 0, err
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
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
	} else {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp SubscribeNewsletterEmail find first error")
		return out, common.EmailAlreadyExists, errors.New(commons.GetCodeAndMsg(common.EmailAlreadyExists, info.Language))
	}
}

func (p portalServiceImp) GetTowerStatus(info request.TowerStats) (out response.TowerStats, code commons.ResponseCode, err error) {

	var vAssets model.Assets
	err = p.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.Rarity, model.AssetsColumns.Type}, map[string]interface{}{
		model.AssetsColumns.ID: info.Id,
	}, nil, &vAssets)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetTowerStatus error:%s", err.Error())
		return out, 0, err
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetTowerStatus assets not find error")
		return out, common.AssetsNotExist, errors.New(commons.GetCodeAndMsg(common.AssetsNotExist, info.Language))
	} else {

		TowerTypeConfigs := make(map[string]model.TowerConfig)
		TowerTypeConfigs[common.TowerTypeConfigs1] = model.TowerConfig{
			AttackFactors:       [3]float32{20, 2, 0.15},
			AttackSpeedFactors:  [3]float32{100, 140, 240},
			AttackRangeFactors:  [3]float32{300, 350, 400},
			DurabilityFactors:   [3]float32{100, 5, 0.7},
			DefaultAttackPerSec: 0.2,
		}
		TowerTypeConfigs[common.TowerTypeConfigs2] = model.TowerConfig{
			AttackFactors:       [3]float32{80, 8, 0.6},
			AttackSpeedFactors:  [3]float32{100, 140, 240},
			AttackRangeFactors:  [3]float32{400, 500, 500},
			DurabilityFactors:   [3]float32{80, 6, 0.4},
			DefaultAttackPerSec: 1.2,
		}
		TowerTypeConfigs[common.TowerTypeConfigs3] = model.TowerConfig{
			AttackFactors:       [3]float32{300, 200, 3},
			AttackSpeedFactors:  [3]float32{100, 130, 180},
			AttackRangeFactors:  [3]float32{600, 650, 700},
			DurabilityFactors:   [3]float32{30, 3, 0.3},
			DefaultAttackPerSec: 4,
		}
		TowerTypeConfigs[common.TowerTypeConfigs4] = model.TowerConfig{
			AttackFactors:       [3]float32{82, 10, 0.8},
			AttackSpeedFactors:  [3]float32{100, 140, 200},
			AttackRangeFactors:  [3]float32{350, 400, 450},
			DurabilityFactors:   [3]float32{70, 4.5, 0.5},
			DefaultAttackPerSec: 3,
		}
		TowerTypeConfigs[common.TowerTypeConfigs5] = model.TowerConfig{
			AttackFactors:       [3]float32{200, 150, 6},
			AttackSpeedFactors:  [3]float32{100, 135, 200},
			AttackRangeFactors:  [3]float32{200, 250, 300},
			DurabilityFactors:   [3]float32{80, 5, 0.6},
			DefaultAttackPerSec: 4,
		}
		RarityConfigs := make(map[string]model.RarityConfig)
		RarityConfigs[common.RarityConfigs1] = model.RarityConfig{
			AttackFactor:     1,
			DurabilityFactor: 1,
		}
		RarityConfigs[common.RarityConfigs2] = model.RarityConfig{
			AttackFactor:     1.2,
			DurabilityFactor: 1.2,
		}
		RarityConfigs[common.RarityConfigs3] = model.RarityConfig{
			AttackFactor:     1.4,
			DurabilityFactor: 1.4,
		}
		RarityConfigs[common.RarityConfigs4] = model.RarityConfig{
			AttackFactor:     1.65,
			DurabilityFactor: 1.65,
		}
		RarityConfigs[common.RarityConfigs5] = model.RarityConfig{
			AttackFactor:     1.8,
			DurabilityFactor: 1.8,
		}
		RarityConfigs[common.RarityConfigs6] = model.RarityConfig{
			AttackFactor:     0.5,
			DurabilityFactor: 0.5,
		}

		if TowerTypeConfig, ok1 := TowerTypeConfigs[strconv.FormatInt(vAssets.Type, 10)]; ok1 {
			if RarityConfig, ok2 := RarityConfigs[strconv.FormatInt(vAssets.Rarity, 10)]; ok2 {
				out = response.TowerStats{
					Attack:      int((TowerTypeConfig.AttackFactors[0] + TowerTypeConfig.AttackFactors[1] + TowerTypeConfig.AttackFactors[2]) * RarityConfig.AttackFactor),
					FireRate:    TowerTypeConfig.DefaultAttackPerSec * TowerTypeConfig.AttackSpeedFactors[0] / 100,
					AttackRange: int(TowerTypeConfig.AttackRangeFactors[0]),
					Durability:  int((TowerTypeConfig.DurabilityFactors[0] + TowerTypeConfig.DurabilityFactors[1] + TowerTypeConfig.DurabilityFactors[2]) * RarityConfig.DurabilityFactor),
					// Durability: "N/A",
				}
				return
			}
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetTowerStatus RarityConfigs map not find TowerRarity error")
			return out, 0, err
		}
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetTowerStatus TowerTypeConfigs map not find TowerType error")
		return out, 0, err
	}
}

func (p portalServiceImp) GetSign(info request.Sign) (out response.Sign, code commons.ResponseCode, err error) {

	mint, ship, _, assets, _, client, err := function.JudgeChain(info.Chain)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign Chain error")
		return out, common.ChainNetError, errors.New("current network is not supported")
	}

	address := ethCommon.HexToAddress(mint)
	instance, err := eth_mint.NewContracts(address, client)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign NewContracts error")
		return out, 0, err
	}

	var vAssets model.Assets
	err = p.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.Category, model.AssetsColumns.Rarity, model.AssetsColumns.Type, model.AssetsColumns.IsNft}, map[string]interface{}{
		model.AssetsColumns.TokenID: info.TokenId,
	}, nil, &vAssets)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign error:%s", err.Error())
		return out, 0, err
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign assets not find error")
		return out, common.AssetsNotExist, errors.New(commons.GetCodeAndMsg(common.AssetsNotExist, info.Language))
	} else {
		// check is nft
		if vAssets.IsNft == common.IsNft {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign asset is nft,do not repeat signatures ")
			return out, 0, errors.New(commons.GetCodeAndMsg(common.AssetsNotExist, info.Language))
		}
		//_tokenId
		tokenId := new(big.Int).SetInt64(info.TokenId)

		if false == function.StringCheck(strconv.Itoa(int(vAssets.Category)), strconv.Itoa(int(vAssets.Type)), strconv.Itoa(int(vAssets.Rarity))) {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp asset data not nil error")
			return out, 0, err
		}

		userAddress := ethCommon.HexToAddress(info.BaseWallet)
		//_category
		category := big.NewInt(vAssets.Category)
		//_subcategory
		subCategory := big.NewInt(vAssets.Type)
		//_rarity
		rarity := big.NewInt(vAssets.Rarity)

		var message [32]byte

		if vAssets.Category == int64(common.Ship) {
			message, err = instance.GetMessageHash(nil, ethCommon.HexToAddress(ship), tokenId, userAddress, category, subCategory, rarity)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign GetMessageHash error:%s", err.Error())
				return out, 0, err
			}
			out.ContractAddress = ship
		} else {
			message, err = instance.GetMessageHash(nil, ethCommon.HexToAddress(assets), tokenId, userAddress, category, subCategory, rarity)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign GetMessageHash error:%s", err.Error())
				return out, 0, err
			}
			out.ContractAddress = assets
		}

		// 调用方法
		req := &proto.SigRequest{Mess: fmt.Sprintf(hex.EncodeToString(message[:]))}

		var res *proto.SigResponse

		ctx, cancel := context.WithTimeout(info.Ctx, common.GrpcTimeoutIn)

		defer cancel()
		res, err = grpc.SignGrpc.SignClient.Sign(ctx, req)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign Sign error:%s", err)
			return out, 0, err
		}
		out.SignMessage = res.Value
		return out, 0, nil
	}
}

func (p portalServiceImp) UserUpdate(info request.UserUpdate) (out response.UserUpdate, code commons.ResponseCode, err error) {

	if len(info.UserName) > 0 || len(info.AvatarAddress) > 0 {
		_, err = p.dao.WithContext(info.Ctx).Update(model.User{
			UserName:      info.UserName,
			AvatarAddress: info.AvatarAddress,
		}, map[string]interface{}{
			model.UserColumns.UUID: info.BaseUUID,
		}, nil)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "PlatformServiceImp UserUpdate error %s", err.Error())
			return out, 0, err
		}
	} else {
		slog.Slog.ErrorF(info.Ctx, "PlatformServiceImp UserUpdate error")
		return out, 0, err
	}
	return
}

func (p portalServiceImp) UserHistory(info request.UserHistory) (out response.UserHistory, code commons.ResponseCode, err error) {
	switch info.Type {
	case common.TransactionHistory:
		//count
		count, err := p.dao.WithContext(info.Ctx).Count(model.TransactionHistory{}, map[string]interface{}{
			model.TransactionHistoryColumns.WalletAddress: info.BaseWallet,
		}, func(db *gorm.DB) *gorm.DB {
			if info.ChainId > 0 {
				db = db.Where("transaction_history.origin_chain=?", info.ChainId)
			}
			if info.FilterTransaction > 0 {
				db = db.Where("transaction_history.status=?", info.FilterTransaction)
			}
			if info.FilterTime.IsZero() == false {
				db = db.Where("transaction_history.created_time > ?", info.FilterTime)
			}
			return db
		})
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp TransactionHistory Count error %s", err.Error())
			return out, 0, err
		}
		var transactionHistoryAssets []join.TransactionHistoryAssets
		err = p.dao.WithContext(info.Ctx).Find([]string{"transaction_history.wallet_address,transaction_history.token_id,transaction_history.price," +
			"transaction_history.unit,transaction_history.origin_chain,transaction_history.status,transaction_history.created_time,assets.name,assets.index_id,assets.nick_name"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
			db = db.Scopes(Paginate(info.CurrentPage, info.PageCount)).
				Joins("LEFT JOIN assets ON transaction_history.token_id = assets.token_id").
				Where("transaction_history.wallet_address=?", info.BaseWallet)
			if info.ChainId > 0 {
				db = db.Where("transaction_history.origin_chain=?", info.ChainId)
			}
			if info.FilterTransaction > 0 {
				db = db.Where("transaction_history.status=?", info.FilterTransaction)
			}
			if info.FilterTime.IsZero() == false {
				db = db.Where("transaction_history.created_time > ?", info.FilterTime)
			}
			db = db.Order(model.TransactionHistoryColumns.CreatedTime + " desc")
			return db
		}, &transactionHistoryAssets)

		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp find transactionHistoryAssets Error: %s", err.Error())
			return out, 0, err
		}

		out.Data = make([]response.HistoryList, 0, len(transactionHistoryAssets))

		for _, transactionHistoryAsset := range transactionHistoryAssets {
			out.Data = append(out.Data, response.HistoryList{
				WalletAddress: transactionHistoryAsset.WalletAddress,
				TokenID:       transactionHistoryAsset.TokenID,
				Price:         transactionHistoryAsset.Price,
				Unit:          transactionHistoryAsset.Unit,
				Status:        transactionHistoryAsset.Status,
				CreatedTime:   transactionHistoryAsset.CreatedTime,
				Name:          transactionHistoryAsset.Name,
				NickName:      transactionHistoryAsset.NickName,
				IndexID:       transactionHistoryAsset.IndexID,
			})
		}
		out.CurrentPage = info.CurrentPage
		out.PrePageCount = info.PageCount
		out.Total = count
		return out, 0, nil
	case common.MintHistory:
		//count
		count, err := p.dao.WithContext(info.Ctx).Count(model.MintHistory{}, map[string]interface{}{
			model.MintHistoryColumns.WalletAddress: info.BaseWallet,
		}, func(db *gorm.DB) *gorm.DB {
			if info.ChainId > 0 {
				db = db.Where("mint_history.origin_chain=?", info.ChainId)
			}
			if info.FilterTransaction > 0 {
				db = db.Where("mint_history.status=?", info.FilterTransaction)
			}
			if info.FilterTime.IsZero() == false {
				db = db.Where("mint_history.created_time > ?", info.FilterTime)
			}
			return db
		})
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp MintHistory Count error %s", err.Error())
			return out, 0, err
		}
		var mintHistoryAssets []join.MintHistoryAssets
		err = p.dao.WithContext(info.Ctx).Find([]string{"mint_history.wallet_address,mint_history.token_id," +
			"mint_history.origin_chain,mint_history.status,mint_history.created_time,assets.name,assets.nick_name"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
			db = db.Scopes(Paginate(info.CurrentPage, info.PageCount)).
				Joins("LEFT JOIN assets ON mint_history.token_id = assets.token_id").
				Where("mint_history.wallet_address=?", info.BaseWallet)
			if info.ChainId > 0 {
				db = db.Where("mint_history.origin_chain=?", info.ChainId)
			}
			if info.FilterTransaction > 0 {
				db = db.Where("mint_history.status=?", info.FilterTransaction)
			}
			if info.FilterTime.IsZero() == false {
				db = db.Where("mint_history.created_time > ?", info.FilterTime)
			}
			db = db.Order(model.MintHistoryColumns.CreatedTime + " desc")
			return db
		}, &mintHistoryAssets)

		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp find mintHistoryAssets Error: %s", err.Error())
			return out, 0, err
		}

		out.Data = make([]response.HistoryList, 0, len(mintHistoryAssets))

		for _, mintHistoryAsset := range mintHistoryAssets {
			out.Data = append(out.Data, response.HistoryList{
				WalletAddress: mintHistoryAsset.WalletAddress,
				TokenID:       mintHistoryAsset.TokenID,
				Status:        mintHistoryAsset.Status,
				CreatedTime:   mintHistoryAsset.CreatedTime,
				Name:          mintHistoryAsset.Name,
				NickName:      mintHistoryAsset.NickName,
			})
		}
		out.CurrentPage = info.CurrentPage
		out.PrePageCount = info.PageCount
		out.Total = count
		return out, 0, nil
	case common.ListenHistory:
	default:
		slog.Slog.ErrorF(info.Ctx, "PlatformServiceImp UserHistory error:history type not exists")
		return out, common.HistoryError, errors.New("history type not exists")
	}
	return
}

func (p portalServiceImp) ExchangePrice(info request.ExchangePrice) (out response.ExchangePrice, code commons.ResponseCode, err error) {
	price, err := exchangeService.Rate(info.Ctx, info.Quote, info.Base)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp ExchangePrice error:%s", err)
		return out, 0, err
	}
	out.Price = price
	return
}

func (p portalServiceImp) AssetDetail(info request.AssetDetail) (out response.AssetDetail, code commons.ResponseCode, err error) {

	var assetsOrders join.AssetsOrders

	if info.AssetId > 0 {
		err = p.dao.WithContext(info.Ctx).First([]string{"assets.is_nft,assets.id,assets.uid,assets.token_id,assets.`name`,assets.nick_name,assets.index_id," +
			"assets.image,assets.description,assets.category,assets.rarity,assets.type,assets.origin_chain,assets.mint_signature,assets.updated_at," +
			"orders_detail.price,orders_detail.order_id,orders.start_time,orders.expire_time,orders.`status`,orders.signature,orders.salt_nonce,`group`.group_name"},
			map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
				db = db.Joins("LEFT JOIN orders_detail ON orders_detail.nft_id = assets.token_id").
					Joins("LEFT JOIN orders ON orders.id = orders_detail.order_id").
					Joins("LEFT JOIN sku ON assets.category = sku.category and assets.type = sku.type and assets.rarity = sku.rarity").
					Joins("LEFT JOIN `group` ON `group`.sku = sku.sku_name").
					Where("assets.id=?", info.AssetId)

				return db
			}, &assetsOrders)
	} else {
		var asset model.Assets
		err = p.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.ID}, map[string]interface{}{
			model.AssetsColumns.TokenID:     info.TokenId,
			model.AssetsColumns.OriginChain: info.ChainId,
		}, nil, &asset)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp find asset by token_id Error: %s", err.Error())
			return out, common.AssetsNotExist, err
		}

		err = p.dao.WithContext(info.Ctx).First([]string{"assets.is_nft,assets.id,assets.uid,assets.token_id,assets.`name`,assets.nick_name,assets.index_id,assets.image," +
			"assets.description,assets.category,assets.rarity,assets.type,assets.origin_chain,assets.mint_signature,assets.updated_at," +
			"orders_detail.price,orders_detail.order_id,orders.start_time,orders.expire_time,orders.`status`,orders.signature,orders.salt_nonce,`group`.group_name"},
			map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
				db = db.Joins("LEFT JOIN orders_detail ON orders_detail.nft_id = assets.token_id").
					Joins("LEFT JOIN orders ON orders.id = orders_detail.order_id").
					Joins("LEFT JOIN sku ON assets.category = sku.category and assets.type = sku.type and assets.rarity = sku.rarity").
					Joins("LEFT JOIN `group` ON `group`.sku = sku.sku_name").
					Where("assets.id=?", asset.ID)
				return db
			}, &assetsOrders)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp find assetsOrders Error: %s", err.Error())
			return out, common.AssetsNotExist, err
		}

	}

	var subCategoryString string
	subCategoryString, err = function.GetSubcategoryString(assetsOrders.Category, assetsOrders.Type)
	if err != nil {
		category := strconv.FormatInt(assetsOrders.Category, 10)
		subCategory := strconv.FormatInt(assetsOrders.Type, 10)
		slog.Slog.ErrorF(info.Ctx, "gameAssetServiceImp SubcategoryString Category:%s,type:%s,Error: %s", category, subCategory, err.Error())
		subCategoryString = "unknown type"
	}

	contractAddress := gameConfig.Contract.Assets
	if assetsOrders.Category == int64(common.Ship) {
		contractAddress = gameConfig.Contract.Ship
	}
	out = response.AssetDetail{
		AssetId:         assetsOrders.Id,
		WalletAddress:   assetsOrders.Uid,
		IsNft:           assetsOrders.IsNft,
		TokenId:         assetsOrders.TokenId,
		ContractAddress: contractAddress,
		ContrainChain:   assetsOrders.OriginChain,
		Name:            assetsOrders.Name,
		IndexID:         assetsOrders.IndexID,
		NickName:        assetsOrders.NickName,
		Image:           assetsOrders.Image,
		Description:     assetsOrders.Description,
		Category:        function.GetCategoryString(assetsOrders.Category),
		CategoryId:      assetsOrders.Category,
		Rarity:          function.GetRarityString(assetsOrders.Rarity),
		RarityId:        assetsOrders.Rarity,
		MintSignature:   assetsOrders.MintSignature,
		SubcategoryId:   assetsOrders.Type,
		Subcategory:     subCategoryString,
		Status:          assetsOrders.Status,
		Price:           assetsOrders.Price,
		OrderId:         assetsOrders.OrderID,
		ExpireTime:      assetsOrders.ExpireTime,
		Signature:       assetsOrders.Signature,
		StartTime:       assetsOrders.StartTime,
		SaltNonce:       assetsOrders.SaltNonce,
		GroupName:       assetsOrders.GroupName,
	}
	return
}

func (p portalServiceImp) GameCurrency(info request.GameCurrency) (out response.GameCurrency, code commons.ResponseCode, err error) {

	party := thirdparty.BaseThirdParty{
		ThirdPartyID: 1,
	}

	gameCurrencyRequest := thirdparty.GameCurrencyRequest{
		WallAddress: info.BaseWallet,
		Symbol:      info.Symbol,
	}

	baseAssets := thirdparty.BaseNotifyEvent{
		Type:      common.GetGMarsEvent,
		EventData: gameCurrencyRequest,
	}
	var gameCurrencyResponse thirdparty.GameCurrencyResponse

	//connect to game
	_, err = p.thirdPartyService.Request(info.Ctx, thirdparty.UriWalletBalance, party, baseAssets, &gameCurrencyResponse)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "GameCurrency call thirdParty failed: %s", err)
		return out, common.GameCurrencyError, err
	}
	out.Amount = gameCurrencyResponse.Amount
	return
}

func (p portalServiceImp) SendCode(info request.SendCode) (out response.SendCode, code commons.ResponseCode, err error) {

	var user model.User
	err = p.dao.WithContext(info.Ctx).First([]string{model.UserColumns.UserName}, map[string]interface{}{
		model.UserColumns.Email: info.Email,
	}, nil, &user)

	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp SendCode find user error %s", err.Error())
		return out, 0, err
	}

	emailCode := time.Now().UnixNano()
	err = p.redis.SetEmailCode(info.Ctx, inner.EmailCode{
		Code:  strconv.FormatInt(emailCode, 10),
		Email: info.Email,
	}, 3*time.Minute)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp SendCode set redis error %s", err.Error())
		return out, 0, err
	}

	to := function.NewEmailUser(user.UserName, info.Email)
	subject := "Sending with Twilio SendGrid is Fun"
	plainTextContent := "example"
	htmlContent := "<strong>example</strong>"

	responseBody, err := function.SendEmail(to, subject, plainTextContent, htmlContent)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp SendEmail error %s", err.Error())
		return out, 0, err
	}
	slog.Slog.InfoF(info.Ctx, "portalServiceImp SendEmail message %s", responseBody)
	return
}

func (p portalServiceImp) PaperMint(info request.PaperMint) (out response.PaperMint, code commons.ResponseCode, err error) {
	var paperMintRequest request.PaperMintRequest
	paperMintRequest.Quantity = 1
	paperMintRequest.ExpiresInMinutes = 15
	paperMintRequest.UsePaperKey = false
	paperMintRequest.HideApplePayGooglePay = false
	paperMintRequest.ContractId = portalConfig.Paper.Payment.MintContractId
	paperMintRequest.WalletAddress = info.WalletAddress
	paperMintRequest.Email = info.Email

	paperMintRequest.MintMethod.Name = "matchMintPaperBSCDummy"
	paperMintRequest.MintMethod.Payment.Value = "0"
	paperMintRequest.MintMethod.Payment.Currency = portalConfig.Paper.Payment.Currency

	//sign
	mint, ship, _, assets, _, client, err := function.JudgeChain(info.ChainId)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint Chain error")
		return out, common.ChainNetError, errors.New("current network is not supported")
	}

	address := ethCommon.HexToAddress(mint)
	instance, err := eth_mint.NewContracts(address, client)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint NewContracts error")
		return out, 0, err
	}

	var vAssets model.Assets
	err = p.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.Category, model.AssetsColumns.Rarity, model.AssetsColumns.Type, model.AssetsColumns.IsNft}, map[string]interface{}{
		model.AssetsColumns.TokenID: info.TokenId,
	}, nil, &vAssets)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint error:%s", err.Error())
		return out, 0, err
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint assets not find error")
		return out, common.AssetsNotExist, errors.New(commons.GetCodeAndMsg(common.AssetsNotExist, info.Language))
	} else {
		// check is nft
		if vAssets.IsNft == common.IsNft {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign asset is nft,do not repeat signatures ")
			return out, 0, errors.New(commons.GetCodeAndMsg(common.AssetsNotExist, info.Language))
		}

		//_tokenId
		tokenId := new(big.Int).SetInt64(info.TokenId)

		if false == function.StringCheck(strconv.Itoa(int(vAssets.Category)), strconv.Itoa(int(vAssets.Type)), strconv.Itoa(int(vAssets.Rarity))) {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint asset data not nil error")
			return out, 0, err
		}

		userAddress := ethCommon.HexToAddress(info.WalletAddress)
		//_category
		category := big.NewInt(vAssets.Category)
		//_subcategory
		subCategory := big.NewInt(vAssets.Type)
		//_rarity
		rarity := big.NewInt(vAssets.Rarity)

		var message [32]byte
		var nftAddress string
		if vAssets.Category == int64(common.Ship) {
			message, err = instance.GetMessageHash(nil, ethCommon.HexToAddress(ship), tokenId, userAddress, category, subCategory, rarity)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint GetMessageHash error:%s", err.Error())
				return out, 0, err
			}
			nftAddress = ship
		} else {
			message, err = instance.GetMessageHash(nil, ethCommon.HexToAddress(assets), tokenId, userAddress, category, subCategory, rarity)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint GetMessageHash error:%s", err.Error())
				return out, 0, err
			}
			nftAddress = assets
		}

		// 调用方法
		req := &proto.SigRequest{Mess: fmt.Sprintf(hex.EncodeToString(message[:]))}

		var res *proto.SigResponse

		ctx, cancel := context.WithTimeout(info.Ctx, common.GrpcTimeoutIn)

		defer cancel()
		res, err = grpc.SignGrpc.SignClient.Sign(ctx, req)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint Sign error:%s", err)
			return out, 0, err
		}

		paperMintRequest.MintMethod.Args.NftAddress = nftAddress
		paperMintRequest.MintMethod.Args.UserAddress = info.WalletAddress
		paperMintRequest.MintMethod.Args.TokenId = info.TokenId
		paperMintRequest.MintMethod.Args.Category = vAssets.Category
		paperMintRequest.MintMethod.Args.Subcategory = vAssets.Type
		paperMintRequest.MintMethod.Args.Rarity = vAssets.Rarity

		if strings.HasPrefix(res.Value, "0x") == false {
			res.Value = "0x" + res.Value
		}
		paperMintRequest.MintMethod.Args.Signature = res.Value
	}

	marshal, err := json.Marshal(&paperMintRequest)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint Sign error:%s", err)
		return out, 0, err
	}
	reader := bytes.NewReader(marshal)
	req, _ := http.NewRequest("POST", portalConfig.Paper.Payment.Url, reader)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", portalConfig.Paper.Payment.Authorization)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &out)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperMint Unmarshal error:%s", err)
		return out, 0, err
	}
	return
}

func (p portalServiceImp) PaperTransaction(info request.PaperTransaction) (out response.PaperTransaction, code commons.ResponseCode, err error) {
	var paperTransactionRequest request.PaperTransactionRequest
	paperTransactionRequest.Quantity = 1
	paperTransactionRequest.ExpiresInMinutes = 15
	paperTransactionRequest.UsePaperKey = false
	paperTransactionRequest.HideApplePayGooglePay = false
	paperTransactionRequest.ContractId = portalConfig.Paper.Payment.MarketContractId
	paperTransactionRequest.WalletAddress = info.WalletAddress
	paperTransactionRequest.Email = info.Email

	paperTransactionRequest.MintMethod.Name = "matchTransactionPaperBscDummy"
	paperTransactionRequest.MintMethod.Payment.Value = info.Value
	paperTransactionRequest.MintMethod.Payment.Currency = portalConfig.Paper.Payment.Currency

	//sign
	_, _, _, assets, spay, _, err := function.JudgeChain(info.ChainId)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperTransaction Chain error")
		return out, common.ChainNetError, errors.New("current network is not supported")
	}

	//tokenId
	var vAssets model.Assets
	err = p.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.UID}, map[string]interface{}{
		model.AssetsColumns.TokenID: info.TokenId,
	}, nil, &vAssets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperTransaction assets by AssetId not find error:%s", err.Error())
		return out, 0, err
	}

	var orderDetail model.OrdersDetail
	err = p.dao.WithContext(info.Ctx).First([]string{model.OrdersDetailColumns.OrderID, model.OrdersDetailColumns.Price}, map[string]interface{}{
		model.OrdersDetailColumns.NftID: info.TokenId,
	}, nil, &orderDetail)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperTransaction First OrderDetail error:%s", err.Error())
		return out, 0, err
	}

	var order model.Orders
	err = p.dao.WithContext(info.Ctx).First([]string{model.OrdersColumns.SaltNonce, model.OrdersColumns.Signature, model.OrdersColumns.Signature,
		model.OrdersColumns.StartTime, model.OrdersColumns.ExpireTime}, map[string]interface{}{
		model.OrdersColumns.ID: orderDetail.OrderID,
	}, nil, &order)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperTransaction First Order error:%s", err.Error())
		return out, 0, err
	}

	saltNonce, _ := strconv.ParseInt(order.SaltNonce, 10, 64)

	paperTransactionRequest.MintMethod.Args.ToAddress = info.WalletAddress
	paperTransactionRequest.MintMethod.Args.OwnerAddress = vAssets.UID
	paperTransactionRequest.MintMethod.Args.NftAddress = assets
	paperTransactionRequest.MintMethod.Args.PaymentToken = spay
	paperTransactionRequest.MintMethod.Args.TokenId = info.TokenId
	paperTransactionRequest.MintMethod.Args.Price = orderDetail.Price
	paperTransactionRequest.MintMethod.Args.StartTime = order.StartTime.Unix()
	paperTransactionRequest.MintMethod.Args.EndTime = order.ExpireTime.Unix()
	paperTransactionRequest.MintMethod.Args.Price = orderDetail.Price
	paperTransactionRequest.MintMethod.Args.SaltNonce = saltNonce
	paperTransactionRequest.MintMethod.Args.Signature = order.Signature

	marshal, err := json.Marshal(&paperTransactionRequest)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperTransaction Sign error:%s", err)
		return out, 0, err
	}
	reader := bytes.NewReader(marshal)
	req, _ := http.NewRequest("POST", portalConfig.Paper.Payment.Url, reader)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", portalConfig.Paper.Payment.Authorization)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &out)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp PaperTransaction Unmarshal error:%s", err)
		return out, 0, err
	}
	return
}
