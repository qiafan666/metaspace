package web

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
	"github.com/jau1jz/cornus"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	SellShelf(info request.SellShelf) (out response.SellShelf, code commons.ResponseCode, err error)
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

	var user model.User
	err = m.dao.WithContext(info.Ctx).First([]string{model.UserColumns.UUID, model.UserColumns.WalletAddress}, map[string]interface{}{
		model.UserColumns.UUID: info.BasePortalRequest.BaseUUID,
	}, nil, &user)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp failed to fetch UUID. Error: %s", err.Error())
		return out, 0, err
	}

	vWalletAddress := strings.ToLower(user.WalletAddress)

	address := ethcommon.HexToAddress(marketConfig.Contract.MarketAddress)
	instance, err := marketcontract.NewContracts(address, ethClient)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature NewContracts error:%s", err.Error())
		return out, 0, err
	}
	//_nftAddress
	//tokenId
	var vAssets model.Assets
	err = m.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.TokenId, model.AssetsColumns.Uid}, map[string]interface{}{
		model.AssetsColumns.Id: info.AssetId,
	}, nil, &vAssets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature assets by AssetId not find error:%s", err.Error())
		return out, 0, err
	}
	tokenId := big.NewInt(vAssets.TokenId)

	if vAssets.TokenId > 0 {

		// check is nft
		addressOwner := ethcommon.HexToAddress(portalConfig.Contract.Erc721Address)
		instanceOwner, errs := assetscontract.NewContracts(addressOwner, ethClient)
		if errs != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature NewContracts error:%s", err.Error())
			return out, 0, errs
		}

		of, errs := instanceOwner.OwnerOf(nil, tokenId)
		if errs != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature get walletAdress error")
			return out, 0, errs
		}
		if strings.ToLower(of.String()) != vWalletAddress {
			//check assets owner
			if vAssets.Uid != strings.ToLower(of.String()) {
				_, errs = m.dao.WithContext(info.Ctx).Update(model.Assets{
					Uid: strings.ToLower(of.String()),
				}, map[string]interface{}{
					model.AssetsColumns.TokenId: vAssets.TokenId,
				}, nil)
				if errs != nil {
					slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update assets uid error %s", err.Error())
					return out, 0, errs
				}
			}

			slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature find assets walletAdress Mismatch with user error")
			return out, 0, errors.New("marketServiceImp GetShelfSignature find assets walletAdress Mismatch with user error")

		}

		var ordersStatus join.OrdersStatus
		err = m.dao.Find([]string{"orders.id,orders.`status`,orders.signature,orders.seller,orders.buyer,orders_detail.nft_id,orders_detail.expire_time"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
			return db.Joins("LEFT JOIN orders_detail ON orders_detail.order_id = orders.id").Where("orders.status=? and orders_detail.nft_id=?", common.OrderStatusActive, strconv.FormatInt(vAssets.TokenId, 10))
		}, &ordersStatus)
		if err == nil && ordersStatus.Id != 0 {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature get orderDetail error")
			return out, common.OrdersIsShelf, errors.New(commons.GetCodeAndMsg(common.OrdersIsShelf, commons.DefualtLanguage))
		} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature get orderStatus error:%s", err.Error())
			return out, 0, err
		}
	}

	//_price
	atoi, err := strconv.Atoi(info.Price)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature assets by AssetId not find error:%s", err.Error())
		return out, 0, err
	}
	price := big.NewInt(int64(atoi))
	//_saltNonce
	randNum := rand.Int63()
	saltNonce := big.NewInt(randNum)
	var message [32]byte
	message, err = instance.GetMessageHash(nil, ethcommon.HexToAddress(portalConfig.Contract.Erc721Address), tokenId, ethcommon.HexToAddress(info.PaymentErc20), price, saltNonce)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSign GetMessageHash error:%s", err.Error())
		return out, 0, err
	}
	out.SignMessage = hex.EncodeToString(message[:])
	out.SaltNonce = randNum
	return
}

func (m marketServiceImp) SellShelf(info request.SellShelf) (out response.SellShelf, code commons.ResponseCode, err error) {

	var user model.User
	err = m.dao.WithContext(info.Ctx).First([]string{model.UserColumns.UUID, model.UserColumns.WalletAddress}, map[string]interface{}{
		model.UserColumns.UUID: info.BasePortalRequest.BaseUUID,
	}, nil, &user)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp failed to fetch UUID. Error: %s", err.Error())
		return out, 0, err
	}

	vWalletAddress := strings.ToLower(user.WalletAddress)

	if info.ExpireTime.Before(time.Now()) {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp expireTime error:%s", err.Error())
		return out, 0, err
	}
	//itemId
	var vAssets model.Assets
	err = m.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.TokenId}, map[string]interface{}{
		model.AssetsColumns.Id: info.ItemId,
	}, nil, &vAssets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSellShelf assets by ItemId not find error:%s", err.Error())
		return out, 0, err
	}

	address := ethcommon.HexToAddress(portalConfig.Contract.Erc721Address)
	instance, err := assetscontract.NewContracts(address, ethClient)
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

	var ordersDetail model.OrdersDetail
	err = m.dao.First([]string{model.OrdersDetailColumns.OrderID}, map[string]interface{}{
		model.OrdersDetailColumns.NftID: vAssets.TokenId,
	}, nil, &ordersDetail)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSellShelf get ordersDetail by tokenId error %s", err.Error())
		return out, 0, err

	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
		newOrders := model.Orders{
			Seller:      vWalletAddress,
			Signature:   info.SignedMessage,
			Status:      common.OrderStatusActive,
			SaltNonce:   info.SaltNonce,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		err = tx.WithContext(info.Ctx).Create(&newOrders)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp orders Create error %s", err.Error())
			return out, 0, err
		}

		newOrdersDetail := model.OrdersDetail{
			OrderID:     newOrders.ID,
			NftID:       strconv.FormatInt(vAssets.TokenId, 10),
			Price:       info.Price,
			ExpireTime:  info.ExpireTime,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}

		err = tx.WithContext(info.Ctx).Create(&newOrdersDetail)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp orders detail Create error %s", err.Error())
			return out, 0, err
		}
		out.RawMessage = info.RawMessage
		out.SignMessage = info.SignedMessage
	} else {
		//update order status
		_, err = tx.WithContext(info.Ctx).Update(model.Orders{
			Status:    common.OrderStatusActive,
			SaltNonce: info.SaltNonce,
			Signature: info.SignedMessage,
		}, map[string]interface{}{
			model.OrdersColumns.ID: ordersDetail.OrderID,
		}, nil)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update Order Status error %s", err.Error())
			return out, 0, err
		}
		//update orders_detail status
		_, err = tx.WithContext(info.Ctx).Update(model.OrdersDetail{
			Price:      info.Price,
			ExpireTime: info.ExpireTime,
		}, map[string]interface{}{
			model.OrdersDetailColumns.NftID: vAssets.TokenId,
		}, nil)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update orders_detail expireTime error %s", err.Error())
			return out, 0, err
		}
		out.RawMessage = info.RawMessage
		out.SignMessage = info.SignedMessage
	}

	return
}

func (m marketServiceImp) GetOrders(info request.Orders) (out response.Orders, code commons.ResponseCode, err error) {

	var ordersDetail []join.OrdersDetail
	err = m.dao.WithContext(info.Ctx).Find([]string{"orders.id,orders.`status`,orders.signature,orders.salt_nonce,orders.buyer,orders.seller,orders_detail.nft_id,orders_detail.price,orders_detail.expire_time,assets.description,assets.image,assets.`name`,assets.category,assets.type,assets.rarity"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db.Scopes(Paginate(info.CurrentPage, info.PageCount)).
			Joins("LEFT JOIN orders_detail ON orders_detail.order_id = orders.id").
			Joins("LEFT JOIN assets ON assets.token_id = orders_detail.nft_id").
			Where("orders.status=?", common.OrderStatusActive)

		if info.Category != nil {
			db = db.Where("assets.category=?", info.Category)
		}
		if info.Rarity != nil {
			db = db.Where("assets.rarity=?", info.Rarity)
		}
		return db
	}, &ordersDetail)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetOrders detail error %s", err.Error())
		return response.Orders{}, 0, err
	}

	out.Data = make([]response.OrdersDetail, 0, len(ordersDetail))

	for _, v := range ordersDetail {
		if v.Id == 0 {
			continue
		}
		//check expireTime
		if v.ExpireTime.Before(time.Now()) {

			//update order status
			_, err = m.dao.WithContext(info.Ctx).Update(model.Orders{
				Status: common.OrderStatusExpire,
			}, map[string]interface{}{
				model.OrdersColumns.ID: v.Id,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update Order Status error %s", err.Error())
				return out, 0, err
			}

			continue
		}

		out.Data = append(out.Data, response.OrdersDetail{
			Id:            v.Id,
			Seller:        v.Seller,
			Buyer:         v.Buyer,
			Signature:     v.Signature,
			SaltNonce:     v.SaltNonce,
			Status:        v.Status,
			NftID:         v.NftID,
			Category:      v.Category,
			Type:          v.Type,
			Rarity:        v.Rarity,
			Image:         v.Image,
			Name:          v.Name,
			Description:   v.Description,
			Price:         v.Price,
			ExpireTime:    v.ExpireTime,
			ContractChain: "BSC",
		})
	}

	count, err := m.dao.WithContext(info.Ctx).Count(model.Orders{}, map[string]interface{}{
		model.OrdersColumns.Status: common.OrderStatusActive,
	}, nil)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp orders Count error %s", err.Error())
		return out, 0, err
	}

	out.Total = count
	out.CurrentPage = info.CurrentPage
	out.PrePageCount = info.PageCount
	return

}

func (m marketServiceImp) GetUserOrders(info request.Orders) (out response.Orders, code commons.ResponseCode, err error) {

	var ordersDetail []join.OrdersDetail
	err = m.dao.WithContext(info.Ctx).Find([]string{"orders.id,orders.`status`,orders.signature,orders.id,orders.buyer,orders.seller,orders_detail.nft_id,assets.description,assets.image,assets.`name`,assets.category,assets.type,assets.rarity"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		return db.Joins("LEFT JOIN orders_detail ON orders_detail.order_id = orders.id").Joins("LEFT JOIN assets ON assets.token_id = orders_detail.nft_id").Where("orders.status=?", info.Status).Where("orders.seller=? or orders.buyer=?", info.BasePortalRequest.BaseUUID, info.BasePortalRequest.BaseUUID)
	}, &ordersDetail)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetUserOrders detail error %s", err.Error())
		return response.Orders{}, 0, err
	}

	out.Data = make([]response.OrdersDetail, 0, len(ordersDetail))
	for _, v := range ordersDetail {
		if v.Id == 0 {
			continue
		}
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

	err = tx.WithContext(info.Ctx).First([]string{model.OrdersColumns.Status, model.OrdersColumns.Seller}, map[string]interface{}{
		model.OrdersColumns.ID: info.OrderId,
	}, func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Locking{Strength: "UPDATE"})
	}, &orders)

	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp orders First error %s", err.Error())
		return out, common.OrdersNotExist, errors.New(commons.GetCodeAndMsg(common.OrdersNotExist, info.Language))
	}

	if orders.Seller != info.BasePortalRequest.BaseUUID {
		slog.Slog.InfoF(info.Ctx, "marketServiceImp order seller")
		return out, common.IdentityError, errors.New(commons.GetCodeAndMsg(common.IdentityError, info.Language))
	}

	if orders.Status == common.OrderStatusCancel {
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
	out.OrderId = info.OrderId
	return

}
