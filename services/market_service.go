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
	"github.com/blockfishio/metaspace-backend/model/join"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jau1jz/cornus"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"gorm.io/gorm"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

// MarketService service layer interface
type MarketService interface {
	GetShelfSignature(info request.ShelfSign) (out response.ShelfSign, code commons.ResponseCode, err error)
	GetSellShelf(info request.SellShelf) (out response.SellShelf, code commons.ResponseCode, err error)
	GetOrders(info request.Orders) (out response.Orders, code commons.ResponseCode, err error)
	GetUserOrders(info request.Orders) (out response.Orders, code commons.ResponseCode, err error)
	OrderCancel(info request.OrderCancel) (out response.OrderCancel, code commons.ResponseCode, err error)
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
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSign GetMessageHash error:%s", err.Error())
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

	if strings.HasPrefix(info.SignedMessage, "0x") == false {
		info.SignedMessage = "0x" + info.SignedMessage
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
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp orders Create error %s", err.Error())
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
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp orders detail Create error %s", err.Error())
		return out, 0, err
	}

	return
}

func (m marketServiceImp) GetOrders(info request.Orders) (out response.Orders, code commons.ResponseCode, err error) {

	var ordersDetail []join.OrdersDetail
	err = m.dao.Find([]string{"orders.id,orders.`status`,orders.signature,orders.id,orders.buyer,orders.seller,orders_detail.nft_id,assets.description,assets.image,assets.`name`,assets.category,assets.type,assets.rarity"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		return db.Joins("LEFT JOIN orders_detail ON orders_detail.order_id = orders.id").Joins("LEFT JOIN assets ON assets.token_id = orders_detail.nft_id").Where("orders.status=?", common.OrderStatusActive)
	}, &ordersDetail)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetOrders detail error %s", err.Error())
		return response.Orders{}, 0, err
	}

	out.Data = make([]response.OrdersDetail, 0, len(ordersDetail))
	for _, v := range ordersDetail {
		out.Data = append(out.Data, response.OrdersDetail{
			Id:          v.Id,
			Seller:      v.Seller,
			Buyer:       v.Buyer,
			Signature:   v.Signature,
			Status:      v.Status,
			NftID:       v.NftID,
			Category:    v.Category,
			Type:        v.Type,
			Rarity:      v.Rarity,
			Image:       v.Image,
			Name:        v.Name,
			Description: v.Description,
		})
	}
	return

}

func (m marketServiceImp) GetUserOrders(info request.Orders) (out response.Orders, code commons.ResponseCode, err error) {

	var ordersDetail []join.OrdersDetail
	err = m.dao.Find([]string{"orders.id,orders.`status`,orders.signature,orders.id,orders.buyer,orders.seller,orders_detail.nft_id,assets.description,assets.image,assets.`name`,assets.category,assets.type,assets.rarity"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		return db.Joins("LEFT JOIN orders_detail ON orders_detail.order_id = orders.id").Joins("LEFT JOIN assets ON assets.token_id = orders_detail.nft_id").Where("orders.status=?", info.Status).Where("orders.seller=? or orders.buyer=?", info.BaseRequest.BaseUUID, info.BaseRequest.BaseUUID)
	}, &ordersDetail)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetUserOrders detail error %s", err.Error())
		return response.Orders{}, 0, err
	}

	out.Data = make([]response.OrdersDetail, 0, len(ordersDetail))
	for _, v := range ordersDetail {
		out.Data = append(out.Data, response.OrdersDetail{
			Id:          v.Id,
			Seller:      v.Seller,
			Buyer:       v.Buyer,
			Signature:   v.Signature,
			Status:      v.Status,
			NftID:       v.NftID,
			Category:    v.Category,
			Type:        v.Type,
			Rarity:      v.Rarity,
			Image:       v.Image,
			Name:        v.Name,
			Description: v.Description,
		})
	}
	return

}

func (m marketServiceImp) OrderCancel(info request.OrderCancel) (out response.OrderCancel, code commons.ResponseCode, err error) {

	tx := m.dao.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var orders model.Orders

	err = tx.First([]string{model.OrdersColumns.Status, model.OrdersColumns.Seller}, map[string]interface{}{
		model.OrdersColumns.ID: info.OrderId,
	}, nil, &orders)

	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp orders First error %s", err.Error())
		return out, common.OrdersNotExist, errors.New(commons.GetCodeAndMsg(common.OrdersNotExist, info.Language))
	}

	if orders.Seller != info.BaseRequest.BaseUUID {
		slog.Slog.InfoF(info.Ctx, "marketServiceImp order seller")
		return out, common.IdentityError, errors.New(commons.GetCodeAndMsg(common.IdentityError, info.Language))
	}

	if orders.Status == 3 {
		slog.Slog.InfoF(info.Ctx, "marketServiceImp order already cancel")
		return out, common.OrderAlreadyCancel, errors.New(commons.GetCodeAndMsg(common.OrderAlreadyCancel, info.Language))
	}

	_, err = tx.WithContext(info.Ctx).Update(model.Orders{
		Status: common.OrderStatusCancel,
	}, map[string]interface{}{
		model.OrdersColumns.ID: info.OrderId,
	}, nil)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp OrderCancel error %s", err.Error())
		return out, 0, err
	}
	return

}
