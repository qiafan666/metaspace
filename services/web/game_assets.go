package web

import (
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/model/join"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"gorm.io/gorm"
	"strconv"
	"time"

	// "gorm.io/gorm"
	"sync"
)

// GameAssetsService service layer interface
type GameAssetsService interface {
	GetGameAssets(info request.GetGameAssets) (out response.GetGameAssets, code commons.ResponseCode, err error)
}

var gameAssetsServiceIns *gameAssetsServiceImp
var gameAssetsServiceInitOnce sync.Once

func NewGameAssetsInstance() GameAssetsService {
	gameAssetsServiceInitOnce.Do(func() {
		gameAssetsServiceIns = &gameAssetsServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return gameAssetsServiceIns
}

type gameAssetsServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (p gameAssetsServiceImp) GetGameAssets(info request.GetGameAssets) (out response.GetGameAssets, code commons.ResponseCode, err error) {

	count, err := p.dao.WithContext(info.Ctx).Count(join.AssetsOrders{}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db = db.Joins("LEFT JOIN orders_detail ON orders_detail.nft_id = assets.token_id").
			Joins("LEFT JOIN orders ON orders.id = orders_detail.order_id").
			Where("assets.uid=?", info.BaseWallet)
		if info.Category != nil {
			db = db.Where("assets.category=?", info.Category)
		}
		if info.IsNft != nil {
			db = db.Where("assets.is_nft=?", info.IsNft)
		}
		if info.Rarity != nil {
			db = db.Where("assets.rarity=?", info.Rarity)
		}
		if info.IsShelf > 0 {
			db = db.Where("assets.is_shelf=?", info.IsShelf)
		}
		return db
	})
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp AssetsOrders count error %s", err.Error())
		return out, 0, err
	}

	var assetsOrders []join.AssetsOrders
	err = p.dao.WithContext(info.Ctx).Find([]string{"assets.is_nft,assets.is_shelf,assets.id,assets.uid,assets.token_id,assets.`name`,assets.nick_name,assets.index_id,assets.image,assets.description,assets.category,assets.rarity,assets.type,assets.mint_signature,assets.updated_at," +
		"orders_detail.price,orders_detail.order_id,orders.start_time,orders.expire_time,orders.`status`,orders.signature,orders.salt_nonce"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db = db.Scopes(Paginate(info.CurrentPage, info.PageCount)).
			Joins("LEFT JOIN orders_detail ON orders_detail.nft_id = assets.token_id").
			Joins("LEFT JOIN orders ON orders.id = orders_detail.order_id").
			Where("assets.uid=?", info.BaseWallet)
		if info.Category != nil {
			db = db.Where("assets.category=?", info.Category)
		}
		if info.IsNft != nil {
			db = db.Where("assets.is_nft=?", info.IsNft)
		}
		if info.Rarity != nil {
			db = db.Where("assets.rarity=?", info.Rarity)
		}
		if info.IsShelf > 0 {
			db = db.Where("assets.is_shelf=?", info.IsShelf)
		}
		db.Order(model.AssetsColumns.UpdatedAt + " desc")
		return db
	}, &assetsOrders)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp find assetsOrders Error: %s", err.Error())
		return out, 0, err
	}

	tx := p.dao.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	for _, vAsset := range assetsOrders {
		//check expireTime
		if !vAsset.ExpireTime.IsZero() && vAsset.ExpireTime.Before(time.Now()) && vAsset.Status == common.OrderStatusActive {
			vAsset.Status = common.OrderStatusExpire

			//update order status
			_, err = tx.WithContext(info.Ctx).Update(model.Orders{
				Status:      common.OrderStatusExpire,
				UpdatedTime: time.Now(),
			}, map[string]interface{}{
				model.OrdersColumns.ID:     vAsset.OrderID,
				model.OrdersColumns.Status: common.OrderStatusActive,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp Update Order Status error %s", err.Error())
				return out, 0, err
			}

			//update assets is_nft
			_, err = tx.WithContext(info.Ctx).Update(model.Assets{
				IsShelf: common.NotShelf,
			}, map[string]interface{}{
				model.AssetsColumns.TokenID: vAsset.TokenId,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp Update assets is_nft error %s", err.Error())
				return out, 0, err
			}

			//add expire history
			newTransactionHistory := model.TransactionHistory{
				WalletAddress: info.BaseWallet,
				TokenID:       vAsset.TokenId,
				Price:         vAsset.Price,
				Status:        common.Expire,
				UpdatedTime:   time.Now(),
				CreatedTime:   time.Now(),
			}
			err = tx.WithContext(info.Ctx).Create(&newTransactionHistory)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp TransactionHistory Create error %s", err.Error())
				return out, 0, err
			}

		}
		var subCategoryString string
		subCategoryString, err = function.GetSubcategoryString(vAsset.Category, vAsset.Type)
		if err != nil {
			category := strconv.FormatInt(vAsset.Category, 10)
			subCategory := strconv.FormatInt(vAsset.Type, 10)
			slog.Slog.ErrorF(info.Ctx, "gameAssetServiceImp SubcategoryString Category:%s,type:%s,Error: %s", category, subCategory, err.Error())
			subCategoryString = "unknown type"
		}

		out.Assets = append(out.Assets, response.AssetBody{
			AssetsId:        vAsset.Id,
			IsNft:           vAsset.IsNft,
			TokenId:         vAsset.TokenId,
			ContractAddress: "0xxxxx",
			ContrainChain:   "BSC",
			Name:            vAsset.Name,
			IndexID:         vAsset.IndexID,
			NickName:        vAsset.NickName,
			Image:           vAsset.Image,
			Description:     vAsset.Description,
			Category:        function.GetCategoryString(vAsset.Category),
			CategoryId:      vAsset.Category,
			Rarity:          function.GetRarityString(vAsset.Rarity),
			RarityId:        vAsset.Rarity,
			MintSignature:   vAsset.MintSignature,
			SubcategoryId:   vAsset.Type,
			Subcategory:     subCategoryString,
			Status:          vAsset.Status,
			Price:           vAsset.Price,
			OrderId:         vAsset.OrderID,
			ExpireTime:      vAsset.ExpireTime,
			Signature:       vAsset.Signature,
			StartTime:       vAsset.StartTime,
			SaltNonce:       vAsset.SaltNonce,
		})

	}
	out.Total = count
	out.CurrentPage = info.CurrentPage
	out.PrePageCount = info.PageCount

	return
}

func Paginate(pageNum int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}
		switch {
		case pageSize > 50:
			pageSize = 50
		case pageSize <= 0:
			pageSize = 8
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
