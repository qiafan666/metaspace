package bizservice

import (
	"encoding/hex"
	"errors"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/contract/assetscontract"
	"github.com/blockfishio/metaspace-backend/contract/marketcontract"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jau1jz/cornus"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"math/big"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// MarketService service layer interface
type MarketService interface {
	GetShelfSignature(info request.ShelfSign) (out response.ShelfSign, code commons.ResponseCode, err error)
	GetSellShelf(info request.SellShelf) (out response.SellShelf, code commons.ResponseCode, err error)
	GetOrders(info request.Orders) (out response.Orders, code commons.ResponseCode, err error)
}

var marketConfig struct {
	Contract struct {
		MarketAddress string `yaml:"market_address"`
	} `yaml:"contract"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&marketConfig)
}

var marketServiceIns *marketServiceImp
var marketServiceInitOnce sync.Once

func NewMarketInstance() *marketServiceImp {
	marketServiceInitOnce.Do(func() {
		marketServiceIns = &marketServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return marketServiceIns
}

type marketServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (m marketServiceImp) GetShelfSignature(info request.ShelfSign) (out response.ShelfSign, code commons.ResponseCode, err error) {

	client, err := ethclient.Dial(portalConfig.Contract.EthClient)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature ethClient Dial error:%s", err.Error())
		return out, 0, err
	}

	address := ethcommon.HexToAddress(marketConfig.Contract.MarketAddress)
	instance, err := marketcontract.NewContracts(address, client)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature NewContracts error:%s", err.Error())
		return out, 0, err
	}
	//_nftAddress
	//tokenId
	var vAssets model.Assets
	err = m.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.TokenId}, map[string]interface{}{
		model.AssetsColumns.Id: info.AssetId,
	}, nil, &vAssets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature assets by AssetId not find error:%s", err.Error())
		return out, 0, err
	}
	tokenId := big.NewInt(vAssets.TokenId)
	//_price
	atoi, err := strconv.Atoi(info.Price)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature assets by AssetId not find error:%s", err.Error())
		return out, 0, err
	}
	price := big.NewInt(int64(atoi))
	//_saltNonce
	saltNonce := big.NewInt(rand.Int63())
	var message [32]byte
	message, err = instance.GetMessageHash(nil, ethcommon.HexToAddress(portalConfig.Contract.Erc721Address), tokenId, ethcommon.HexToAddress(info.PaymentErc20), price, saltNonce)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign GetMessageHash error:%s", err.Error())
		return out, 0, err
	}
	out.SignMessage = hex.EncodeToString(message[:])
	return
}

func (m marketServiceImp) GetSellShelf(info request.SellShelf) (out response.SellShelf, code commons.ResponseCode, err error) {

	//itemId
	var vAssets model.Assets
	err = m.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.TokenId}, map[string]interface{}{
		model.AssetsColumns.Id: info.ItemId,
	}, nil, &vAssets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSellShelf assets by ItemId not find error:%s", err.Error())
		return out, 0, err
	}

	client, err := ethclient.Dial(portalConfig.Contract.EthClient)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature ethClient Dial error:%s", err.Error())
		return out, 0, err
	}

	address := ethcommon.HexToAddress(portalConfig.Contract.Erc721Address)
	instance, err := assetscontract.NewContracts(address, client)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature NewContracts error:%s", err.Error())
		return out, 0, err
	}

	of, err := instance.OwnerOf(nil, big.NewInt(vAssets.TokenId))
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp get userAddress error")
		return out, 0, err
	}

	if err = function.VerifySig(of.String(), info.SignedMessage, info.RawMessage); err != nil {
		slog.Slog.InfoF(info.Ctx, "marketServiceImp GetSellShelf verify error %s", err.Error())
		return out, common.SignatureVerificationError, err
	}

	tx := m.dao.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	//check account exists
	count, err := m.dao.Count(model.Orders{}, map[string]interface{}{
		model.OrdersColumns.Signature: info.SignedMessage,
	}, nil)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSellShelf Count error %s", err.Error())
		return out, 0, err
	}
	if count > 0 {
		slog.Slog.InfoF(info.Ctx, "marketServiceImp GetSellShelf account already exists")
		return out, common.AccountAlreadyExists, errors.New(commons.GetCodeAndMsg(common.AccountAlreadyExists, info.Language))
	}
	orders := model.Orders{
		Seller:      info.BaseRequest.BaseUUID,
		Signature:   info.SignedMessage,
		Status:      1,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	err = tx.WithContext(info.Ctx).Create(&orders)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp orders Create error %s", err.Error())
		return out, 0, err
	}

	ordersDetail := model.OrdersDetail{
		OrderID:     orders.ID,
		NftID:       strconv.FormatInt(vAssets.TokenId, 10),
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	err = tx.WithContext(info.Ctx).Create(&ordersDetail)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp orders detail Create error %s", err.Error())
		return out, 0, err
	}

	return
}

func (m marketServiceImp) GetOrders(info request.Orders) (out response.Orders, code commons.ResponseCode, err error) {

	var orders []model.Orders
	err = m.dao.WithContext(info.Ctx).Find([]string{}, map[string]interface{}{
		model.OrdersColumns.Status: info.Status,
	}, nil, &orders)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp find orders Error: %s", err.Error())
		return out, 0, err
	}
	out.Data = orders
	return
}
