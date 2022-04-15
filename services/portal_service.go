package bizservice

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/blockfishio/metaspace-backend/contract"
	"github.com/blockfishio/metaspace-backend/grpc"
	"github.com/blockfishio/metaspace-backend/grpc/proto"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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
	//Login support email and wallet login api
	Login(info request.UserLogin) (out response.UserLogin, code commons.ResponseCode, err error)
	//GetNonce client get new nonce from server
	GetNonce(info request.GetNonce) (out response.GetNonce, code commons.ResponseCode, err error)
	Register(info request.RegisterUser) (out response.RegisterUser, code commons.ResponseCode, err error)
	UpdatePassword(info request.PasswordUpdate) (out response.PasswordUpdate, code commons.ResponseCode, err error)
	SubscribeNewsletterEmail(info request.SubscribeNewsletterEmail) (out response.SubscribeNewsletterEmail, code commons.ResponseCode, err error)
	GetTowerStatus(info request.TowerStats) (out response.TowerStats, code commons.ResponseCode, err error)
	GetSign(info request.Sign) (out response.Sign, code commons.ResponseCode, err error)
}

var portalConfig struct {
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	Contract struct {
		EthClient     string `yaml:"eth_client"`
		NftAddress    string `yaml:"nft_address"`
		Erc721Address string `yaml:"erc721_address"`
	} `yaml:"contract"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&portalConfig)
}

var portalServiceIns *portalServiceImp
var portalServiceInitOnce sync.Once

func NewPortalServiceInstance() *portalServiceImp {
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

	if _, err := strconv.Atoi(info.TowerType); err != nil {
		slog.Slog.ErrorF(info.Ctx, "TowerType error:%s", err.Error())
		return out, 0, err
	}

	if _, err := strconv.Atoi(info.TowerRarity); err != nil {
		slog.Slog.ErrorF(info.Ctx, "TowerRarity error:%s", err.Error())
		return out, 0, err
	}
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

	if TowerTypeConfig, ok1 := TowerTypeConfigs[info.TowerType]; ok1 {
		if RarityConfig, ok2 := RarityConfigs[info.TowerRarity]; ok2 {
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

func (p portalServiceImp) GetSign(info request.Sign) (out response.Sign, code commons.ResponseCode, err error) {

	client, err := ethclient.Dial(portalConfig.Contract.EthClient)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign ethClient Dial error")
		return out, 0, err
	}

	address := ethcommon.HexToAddress(portalConfig.Contract.NftAddress)
	instance, err := contract.NewContracts(address, client)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign NewContracts error")
		return out, 0, err
	}

	var vAssets model.Assets
	err = p.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.Category, model.AssetsColumns.Rarity, model.AssetsColumns.Type}, map[string]interface{}{
		model.AssetsColumns.TokenId: info.Id,
	}, nil, &vAssets)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign error:%s", err.Error())
		return out, 0, err
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign assets not find error")
		return out, common.AssetsNotExist, errors.New(commons.GetCodeAndMsg(common.AssetsNotExist, info.Language))
	} else {
		//_nftAddress
		//_tokenId
		var id int
		id, err = strconv.Atoi(info.Id)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign id format error")
			return out, 0, err
		}
		tokenId := new(big.Int).SetUint64(uint64(id))

		//_ownerAddress
		var user model.User
		err = p.dao.WithContext(info.Ctx).First([]string{model.UserColumns.WalletAddress}, map[string]interface{}{
			model.UserColumns.UUID: info.BaseUUID,
		}, nil, &user)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign User by UUid not find error")
			return out, 0, err
		}

		if false == function.StringCheck(user.WalletAddress, strconv.Itoa(int(vAssets.Category)), strconv.Itoa(int(vAssets.Type)), strconv.Itoa(int(vAssets.Rarity))) {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp asset data not nil error")
			return out, 0, err
		}

		userAddress := ethcommon.HexToAddress(user.WalletAddress)
		//_category
		category := big.NewInt(vAssets.Category)
		//_subcategory
		subCategory := big.NewInt(vAssets.Type)
		//_rarity
		rarity := big.NewInt(vAssets.Rarity)

		var message [32]byte
		message, err = instance.GetMessageHash(nil, ethcommon.HexToAddress(portalConfig.Contract.Erc721Address), tokenId, userAddress, category, subCategory, rarity)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign GetMessageHash error:%s", err.Error())
			return out, 0, err
		}

		// 调用方法
		req := &proto.SigRequest{Mess: fmt.Sprintf(hex.EncodeToString(message[:]))}

		var res *proto.SigResponse

		ctx, cancel := context.WithTimeout(info.BaseRequest.Ctx, common.GrpcTimeoutInSec)
		defer cancel()
		res, err = grpc.SignGrpc.SignClient.Sign(ctx, req)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign Sign error:%s", err)
			return out, 0, err
		}
		out.SignMessage = res.Value
		return
	}

}
