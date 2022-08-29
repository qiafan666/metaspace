package web

import (
	"errors"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/contract/eth/eth_assets"
	"github.com/blockfishio/metaspace-backend/contract/eth/eth_market"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/model/join"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/big"
	"math/rand"
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
	GetOrdersGroup(info request.OrdersGroup) (out []response.OrdersGroup, code commons.ResponseCode, err error)
	GetOrdersGroupDetail(info request.OrdersGroupDetail) (out []response.OrdersGroupDetail, code commons.ResponseCode, err error)
}

var marketServiceIns *marketServiceImp
var marketServiceInitOnce sync.Once

func NewMarketInstance() MarketService {
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

	//tokenId
	var vAssets model.Assets
	err = m.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.TokenID, model.AssetsColumns.UID, model.AssetsColumns.Category, model.AssetsColumns.OriginChain}, map[string]interface{}{
		model.AssetsColumns.ID: info.AssetId,
	}, nil, &vAssets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature assets by AssetId not find error:%s", err.Error())
		return out, 0, err
	}

	_, ship, market, assets, client, err := function.JudgeChain(vAssets.OriginChain)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign Chain error")
		return out, common.ChainNetError, errors.New("current network is not supported")
	}

	tokenId := big.NewInt(vAssets.TokenID)
	if vAssets.TokenID > 0 {
		// check is nft

		var address ethcommon.Address
		if vAssets.Category == int64(common.Ship) {
			address = ethcommon.HexToAddress(ship)
		} else {
			address = ethcommon.HexToAddress(assets)
		}

		instanceOwner, errs := eth_assets.NewContracts(address, client)
		if errs != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature address not match error:%s", err.Error())
			return out, 0, errs
		}

		of, errs := instanceOwner.OwnerOf(nil, tokenId)
		if errs != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature get walletAdress error")
			return out, 0, errs
		}
		if strings.ToLower(of.String()) != info.BaseWallet {
			//check assets owner
			if vAssets.UID != strings.ToLower(of.String()) {
				_, errs = m.dao.WithContext(info.Ctx).Update(model.Assets{
					UID:       strings.ToLower(of.String()),
					UpdatedAt: time.Now(),
				}, map[string]interface{}{
					model.AssetsColumns.TokenID: vAssets.TokenID,
				}, nil)
				if errs != nil {
					slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update assets uid error %s", err.Error())
					return out, 0, errs
				}
			}

			slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature find assets walletAdress Mismatch with user error")
			return out, 0, errors.New("marketServiceImp GetShelfSignature find assets walletAdress Mismatch with user error")

		}

	} else {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature error: tokenId not nil")
		return out, 0, err
	}

	//_price
	price, flag := big.NewInt(0).SetString(info.Price, 10)
	if flag == false {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature price setString error")
		return out, commons.ParameterError, err
	}
	//_saltNonce
	saltNonce := big.NewInt(int64(rand.Int31()))

	startTime := time.Now()
	endTime := info.ExpireTime

	address := ethcommon.HexToAddress(market)
	instance, err := eth_market.NewContracts(address, client)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetShelfSignature NewContracts error:%s", err.Error())
		return out, 0, err
	}

	message, err := instance.GetMessageHash(nil, ethcommon.HexToAddress(assets), tokenId, ethcommon.HexToAddress(info.PaymentErc20), price, big.NewInt(startTime.Unix()), big.NewInt(endTime.Unix()), saltNonce)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSign GetMessageHash error:%s", err.Error())
		return out, 0, err
	}
	out.SignMessage = ethcommon.BytesToHash(message[:]).String()
	out.SaltNonce = saltNonce.String()
	err = m.redis.SetRawMessage(info.Ctx, inner.RawMessage{
		RawMessage: out.SignMessage,
		StartTime:  startTime,
		ExpireTime: endTime,
	}, time.Minute*3)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp SetTokenUser error %s", err.Error())
		return out, 0, err
	}
	return out, 0, nil
}

func (m marketServiceImp) SellShelf(info request.SellShelf) (out response.SellShelf, code commons.ResponseCode, err error) {

	signMessage, err := m.redis.GetRawMessage(info.Ctx, info.RawMessage)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetRawMessage error %s", err.Error())
		return
	}

	err = m.redis.DelRawMessage(info.Ctx, info.RawMessage)
	if err != nil {
		slog.Slog.InfoF(info.Ctx, "marketServiceImp DelRawMessage error %s", err.Error())
		return out, 0, err
	}

	if signMessage.ExpireTime.Before(time.Now()) {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp expireTime error:%s", err.Error())
		return out, 0, err
	}
	//itemId
	var vAssets model.Assets
	err = m.dao.WithContext(info.Ctx).First([]string{model.AssetsColumns.TokenID, model.AssetsColumns.Category, model.AssetsColumns.OriginChain}, map[string]interface{}{
		model.AssetsColumns.ID: info.ItemId,
	}, nil, &vAssets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSellShelf assets by ItemId not find error:%s", err.Error())
		return out, 0, err
	}

	_, ship, market, assets, client, err := function.JudgeChain(vAssets.OriginChain)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "portalServiceImp GetSign Chain error")
		return out, common.ChainNetError, errors.New("current network is not supported")
	}

	var address ethcommon.Address
	if vAssets.Category == int64(common.Ship) {
		address = ethcommon.HexToAddress(ship)
	} else {
		address = ethcommon.HexToAddress(assets)
	}

	instance, err := eth_assets.NewContracts(address, client)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSellShelf NewContracts error:%s", err.Error())
		return out, 0, err
	}

	of, err := instance.OwnerOf(nil, big.NewInt(vAssets.TokenID))
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp get userAddress error: %s", err.Error())
		return out, 0, err
	}

	//check
	if info.BaseWallet != strings.ToLower(of.String()) {
		slog.Slog.ErrorF(info.Ctx, "assets address not wallet address")
		return out, common.WalletError, errors.New("inconsistent wallet addresses")
	}

	if strings.HasPrefix(info.SignedMessage, "0x") == false {
		info.SignedMessage = "0x" + info.SignedMessage
	}

	marketAddress := ethcommon.HexToAddress(market)
	marketInstance, err := eth_market.NewContracts(marketAddress, client)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSellShelf marketNewContracts error:%s", err.Error())
		return out, 0, err
	}

	flag, err := marketInstance.UsedSignatures(nil, []byte(info.SignedMessage))
	if flag == true {
		slog.Slog.InfoF(info.Ctx, "marketServiceImp GetSellShelf Signatures already used : %s", err.Error())
		return out, common.UsedSignature, err
	}

	if err = function.VerifySigEthHash(of.String(), info.SignedMessage, info.RawMessage); err != nil {
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
		model.OrdersDetailColumns.NftID: vAssets.TokenID,
	}, nil, &ordersDetail)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetSellShelf get ordersDetail by tokenId error %s", err.Error())
		return out, 0, err

	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == true {
		newOrders := model.Orders{
			Seller:      info.BaseWallet,
			Signature:   info.SignedMessage,
			Status:      common.OrderStatusActive,
			SaltNonce:   info.SaltNonce,
			TotalPrice:  info.Price,
			StartTime:   signMessage.StartTime,
			ExpireTime:  signMessage.ExpireTime,
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
			NftID:       vAssets.TokenID,
			Price:       info.Price,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}

		err = tx.WithContext(info.Ctx).Create(&newOrdersDetail)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp orders detail Create error %s", err.Error())
			return out, 0, err
		}

		//add shelf history
		newTransactionHistory := model.TransactionHistory{
			WalletAddress: info.BaseWallet,
			TokenID:       vAssets.TokenID,
			Price:         info.Price,
			OriginChain:   vAssets.OriginChain,
			Status:        common.Shelf,
			UpdatedTime:   time.Now(),
			CreatedTime:   time.Now(),
		}

		err = tx.WithContext(info.Ctx).Create(&newTransactionHistory)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp TransactionHistory Create error %s", err.Error())
			return out, 0, err
		}

		//update assets is_nft
		_, err = tx.WithContext(info.Ctx).Update(model.Assets{
			IsShelf: common.IsShelf,
		}, map[string]interface{}{
			model.AssetsColumns.TokenID: vAssets.TokenID,
		}, nil)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update assets is_nft error %s", err.Error())
			return out, 0, err
		}

		out.RawMessage = info.RawMessage
		out.SignMessage = info.SignedMessage
	} else {
		//update order status
		_, err = tx.WithContext(info.Ctx).Update(model.Orders{
			Status:      common.OrderStatusActive,
			SaltNonce:   info.SaltNonce,
			Signature:   info.SignedMessage,
			TotalPrice:  info.Price,
			StartTime:   signMessage.StartTime,
			ExpireTime:  signMessage.ExpireTime,
			UpdatedTime: time.Now(),
		}, map[string]interface{}{
			model.OrdersColumns.ID: ordersDetail.OrderID,
		}, nil)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update Order Status error %s", err.Error())
			return out, 0, err
		}
		//update orders_detail status
		_, err = tx.WithContext(info.Ctx).Update(model.OrdersDetail{
			Price:       info.Price,
			UpdatedTime: time.Now(),
		}, map[string]interface{}{
			model.OrdersDetailColumns.NftID: vAssets.TokenID,
		}, nil)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update orders_detail expireTime error %s", err.Error())
			return out, 0, err
		}

		//add shelf history
		newTransactionHistory := model.TransactionHistory{
			WalletAddress: info.BaseWallet,
			TokenID:       vAssets.TokenID,
			Price:         info.Price,
			OriginChain:   vAssets.OriginChain,
			Status:        common.Shelf,
			UpdatedTime:   time.Now(),
			CreatedTime:   time.Now(),
		}

		err = tx.WithContext(info.Ctx).Create(&newTransactionHistory)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp TransactionHistory Create error %s", err.Error())
			return out, 0, err
		}

		//update assets is_nft
		_, err = tx.WithContext(info.Ctx).Update(model.Assets{
			IsShelf: common.IsShelf,
		}, map[string]interface{}{
			model.AssetsColumns.TokenID: vAssets.TokenID,
		}, nil)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update assets is_nft error %s", err.Error())
			return out, 0, err
		}

		out.RawMessage = info.RawMessage
		out.SignMessage = info.SignedMessage
	}
	return out, 0, nil
}

func (m marketServiceImp) GetOrders(info request.Orders) (out response.Orders, code commons.ResponseCode, err error) {
redo:
	count, err := m.dao.WithContext(info.Ctx).Count(join.OrdersDetail{}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db.Joins("INNER JOIN orders_detail ON orders_detail.order_id = orders.id").
			Joins("INNER JOIN assets ON assets.token_id = orders_detail.nft_id").
			Where("orders.status=?", common.OrderStatusActive)
		if info.ChainId > 0 {
			db = db.Where("assets.origin_chain=?", info.ChainId)
		}
		if info.Category != nil {
			db = db.Where("assets.category=?", info.Category)
		}
		if info.Rarity != nil {
			db = db.Where("assets.rarity=?", info.Rarity)
		}

		if len(info.Search) > 0 {
			return db.Where("LOWER(assets.nick_name) Like LOWER(?)", "%"+info.Search+"%")
		}

		return db
	})
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp orders Count error %s", err.Error())
		return out, 0, err
	}

	var ordersDetail []join.OrdersDetail
	err = m.dao.WithContext(info.Ctx).Find([]string{"orders.id,orders.`status`,orders.signature,orders.salt_nonce,orders.buyer,orders.seller,orders.total_price,orders.start_time,orders.expire_time,orders.updated_time,orders_detail.nft_id,orders_detail.price,assets.id as asset_id,assets.description,assets.image,assets.`name`,assets.category,assets.type,assets.rarity,assets.index_id,assets.nick_name,assets.origin_chain"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db.Scopes(Paginate(info.CurrentPage, info.PageCount)).
			Joins("INNER JOIN orders_detail ON orders_detail.order_id = orders.id").
			Joins("INNER JOIN assets ON assets.token_id = orders_detail.nft_id").
			Where("orders.status=?", common.OrderStatusActive)
		//filter
		if info.ChainId > 0 {
			db = db.Where("assets.origin_chain=?", info.ChainId)
		}

		if info.Category != nil {
			db = db.Where("assets.category=?", info.Category)
		}
		if info.Rarity != nil {
			db = db.Where("assets.rarity=?", info.Rarity)
		}

		//search
		if len(info.Search) > 0 {
			return db.Where("LOWER(assets.nick_name) Like LOWER(?)", "%"+info.Search+"%")
		}

		if info.SortTime > 0 && info.SortPrice > 0 {
			return db.Order(model.OrdersColumns.UpdatedTime + " desc")
		}

		if info.SortTime == 0 {
		} else if info.SortTime == 1 {
			return db.Order(model.OrdersColumns.UpdatedTime + " desc")
		} else {
			return db.Order(model.OrdersColumns.UpdatedTime + " asc")
		}

		if info.SortPrice == 0 {
		} else if info.SortPrice == 1 {
			return db.Order("--" + model.OrdersDetailColumns.Price + " desc")
		} else {
			return db.Order("--" + model.OrdersDetailColumns.Price + " asc")
		}

		return db.Order(model.OrdersColumns.UpdatedTime + " desc")
	}, &ordersDetail)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetOrders detail error %s", err.Error())
		return out, 0, err
	}

	tx := m.dao.Tx()

	out.Data = make([]response.OrdersDetail, 0, len(ordersDetail))
	redoFlag := false
	var contractAddress string
	for _, v := range ordersDetail {
		//check expireTime
		if v.ExpireTime.Before(time.Now()) {
			//update order status
			result, err := tx.WithContext(info.Ctx).Update(model.Orders{
				Status: common.OrderStatusExpire,
			}, map[string]interface{}{
				model.OrdersColumns.ID: v.Id,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update Order Status error %s", err.Error())
				tx.Rollback()
				return out, 0, err
			}
			if result == 0 {
				redoFlag = true
				continue
			}
			//update assets is_nft
			_, err = tx.WithContext(info.Ctx).Update(model.Assets{
				IsShelf: common.NotShelf,
			}, map[string]interface{}{
				model.AssetsColumns.TokenID: v.NftID,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update assets is_nft error %s", err.Error())
				tx.Rollback()
				return out, 0, err
			}

			//add expire history
			newTransactionHistory := model.TransactionHistory{
				WalletAddress: v.Seller,
				TokenID:       v.NftID,
				Price:         v.Price,
				OriginChain:   v.OriginChain,
				Status:        common.Expire,
				UpdatedTime:   time.Now(),
				CreatedTime:   time.Now(),
			}
			err = tx.WithContext(info.Ctx).Create(&newTransactionHistory)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "marketServiceImp TransactionHistory Create error %s", err.Error())
				tx.Rollback()
				return out, 0, err
			}
		} else {

			if v.Category == int64(common.Ship) {
				contractAddress = gameConfig.Contract.Ship
			} else {
				contractAddress = gameConfig.Contract.Assets
			}

			out.Data = append(out.Data, response.OrdersDetail{
				AssetId:         v.AssetId,
				Id:              v.Id,
				Seller:          v.Seller,
				Buyer:           v.Buyer,
				Signature:       v.Signature,
				SaltNonce:       v.SaltNonce,
				Status:          v.Status,
				NftID:           v.NftID,
				Category:        v.Category,
				Type:            v.Type,
				Rarity:          v.Rarity,
				Image:           v.Image,
				Name:            v.Name,
				IndexID:         v.IndexID,
				NickName:        v.NickName,
				Description:     v.Description,
				TotalPrice:      v.TotalPrice,
				Price:           v.Price,
				StartTime:       v.StartTime,
				ExpireTime:      v.ExpireTime,
				ContractChain:   v.OriginChain,
				ContractAddress: contractAddress,
			})
		}
	}
	_ = tx.Commit()
	if redoFlag {
		goto redo

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

	if orders.Seller != info.BaseWallet {
		slog.Slog.InfoF(info.Ctx, "marketServiceImp order seller is error")
		return out, common.IdentityError, errors.New(commons.GetCodeAndMsg(common.IdentityError, info.Language))
	}

	if orders.Status == common.OrderStatusCancel {
		slog.Slog.InfoF(info.Ctx, "marketServiceImp order already cancel")
		return out, common.OrderAlreadyCancel, errors.New(commons.GetCodeAndMsg(common.OrderAlreadyCancel, info.Language))
	}

	_, err = tx.WithContext(info.Ctx).Update(model.Orders{
		Status:      common.OrderStatusCancel,
		UpdatedTime: time.Now(),
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

func (m marketServiceImp) GetOrdersGroup(info request.OrdersGroup) (out []response.OrdersGroup, code commons.ResponseCode, err error) {

	var assets []model.Assets
	err = m.dao.WithContext(info.Ctx).Raw("select aa.* from assets as aa right join ( select sku, max(id+0) as max from assets group by sku) as bb on bb.max = aa.id and bb.sku = aa.sku", &assets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp GetOrdersGroup error %s", err.Error())
		return out, 0, err
	}

	out = make([]response.OrdersGroup, 0, len(assets))
	for _, v := range assets {
		out = append(out, response.OrdersGroup{
			Sku:         v.Sku,
			Image:       v.Image,
			Description: v.Description,
			URI:         v.URI,
			URIContent:  v.URIContent,
		})
	}
	return
}

func (m marketServiceImp) GetOrdersGroupDetail(info request.OrdersGroupDetail) (out []response.OrdersGroupDetail, code commons.ResponseCode, err error) {

	var ordersDetail []join.OrdersDetail
	err = m.dao.WithContext(info.Ctx).Find([]string{"orders.id,orders.`status`,orders.signature,orders.salt_nonce,orders.buyer,orders.seller,orders.total_price,orders.start_time,orders.expire_time,orders.updated_time,orders_detail.nft_id,orders_detail.price,assets.id as asset_id,assets.description,assets.image,assets.`name`,assets.category,assets.type,assets.rarity,assets.index_id,assets.nick_name,assets.origin_chain,assets.sku"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db.Joins("INNER JOIN orders_detail ON orders_detail.order_id = orders.id").
			Joins("INNER JOIN assets ON assets.token_id = orders_detail.nft_id").
			Where("orders.status=?", common.OrderStatusActive).
			Where("assets.sku=?", info.Sku).Limit(50)
		//filter
		if info.ChainId > 0 {
			db = db.Where("assets.origin_chain=?", info.ChainId)
		}

		//search
		if info.SortPrice == 0 {
		} else if info.SortPrice == 1 {
			return db.Order("--" + model.OrdersDetailColumns.Price + " desc")
		} else {
			return db.Order("--" + model.OrdersDetailColumns.Price + " asc")
		}

		return db.Order(model.OrdersColumns.UpdatedTime + " desc")
	}, &ordersDetail)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "marketServiceImp  GetOrdersGroupDetail error %s", err.Error())
		return out, 0, err
	}

	tx := m.dao.Tx()

	out = make([]response.OrdersGroupDetail, 0, len(ordersDetail))
	var contractAddress string
	for _, v := range ordersDetail {
		//check expireTime
		if v.ExpireTime.Before(time.Now()) {
			//update order status
			result, err := tx.WithContext(info.Ctx).Update(model.Orders{
				Status: common.OrderStatusExpire,
			}, map[string]interface{}{
				model.OrdersColumns.ID: v.Id,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update Order Status error %s", err.Error())
				tx.Rollback()
				return out, 0, err
			}
			if result == 0 {
				continue
			}
			//update assets is_nft
			_, err = tx.WithContext(info.Ctx).Update(model.Assets{
				IsShelf: common.NotShelf,
			}, map[string]interface{}{
				model.AssetsColumns.TokenID: v.NftID,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "marketServiceImp Update assets is_nft error %s", err.Error())
				tx.Rollback()
				return out, 0, err
			}

			//add expire history
			newTransactionHistory := model.TransactionHistory{
				WalletAddress: v.Seller,
				TokenID:       v.NftID,
				Price:         v.Price,
				OriginChain:   v.OriginChain,
				Status:        common.Expire,
				UpdatedTime:   time.Now(),
				CreatedTime:   time.Now(),
			}
			err = tx.WithContext(info.Ctx).Create(&newTransactionHistory)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "marketServiceImp TransactionHistory Create error %s", err.Error())
				tx.Rollback()
				return out, 0, err
			}
		} else {

			if v.Category == int64(common.Ship) {
				contractAddress = gameConfig.Contract.Ship
			} else {
				contractAddress = gameConfig.Contract.Assets
			}

			out = append(out, response.OrdersGroupDetail{
				AssetId:         v.AssetId,
				Id:              v.Id,
				Seller:          v.Seller,
				Buyer:           v.Buyer,
				Signature:       v.Signature,
				SaltNonce:       v.SaltNonce,
				Status:          v.Status,
				NftID:           v.NftID,
				Category:        v.Category,
				Type:            v.Type,
				Rarity:          v.Rarity,
				Image:           v.Image,
				Name:            v.Name,
				IndexID:         v.IndexID,
				NickName:        v.NickName,
				Description:     v.Description,
				TotalPrice:      v.TotalPrice,
				Price:           v.Price,
				StartTime:       v.StartTime,
				ExpireTime:      v.ExpireTime,
				ContractChain:   v.OriginChain,
				ContractAddress: contractAddress,
			})
		}
	}

	return
}
