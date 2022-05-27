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
	"strings"
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

	var user model.User
	err = p.dao.WithContext(info.Ctx).First([]string{model.UserColumns.UUID, model.UserColumns.WalletAddress}, map[string]interface{}{
		model.UserColumns.UUID: info.BasePortalRequest.BaseUUID,
	}, nil, &user)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp failed to fetch UUID. Error: %s", err.Error())
		return out, 0, err
	}

	vWalletAddress := strings.ToLower(user.WalletAddress)

	count, err := p.dao.WithContext(info.Ctx).Count(model.Assets{}, map[string]interface{}{
		model.AssetsColumns.Uid: vWalletAddress,
	}, func(db *gorm.DB) *gorm.DB {
		if info.Category != nil {
			db = db.Where("assets.category=?", info.Category)
		}
		if info.IsNft != nil {
			db = db.Where("assets.is_nft=?", info.IsNft)
		}
		if info.Rarity != nil {
			db = db.Where("assets.rarity=?", info.Rarity)
		}
		return db
	})
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp assets Count error %s", err.Error())
		return out, 0, err
	}

	var assetsOrders []join.AssetsOrders
	err = p.dao.WithContext(info.Ctx).Find([]string{"assets.is_nft,assets.id,assets.uid,assets.token_id,assets.`name`,assets.image,assets.description,assets.category,assets.rarity,assets.type,assets.mint_signature," +
		"orders_detail.price,orders_detail.order_id,orders_detail.expire_time,orders.`status`,orders.signature"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db = db.Scopes(Paginate(info.CurrentPage, info.PageCount)).
			Joins("LEFT JOIN orders_detail ON orders_detail.nft_id = assets.token_id").
			Joins("LEFT JOIN orders ON orders.id = orders_detail.order_id").
			Where("assets.uid=?", vWalletAddress)
		if info.Category != nil {
			db = db.Where("assets.category=?", info.Category)
		}
		if info.IsNft != nil {
			db = db.Where("assets.is_nft=?", info.IsNft)
		}
		if info.Rarity != nil {
			db = db.Where("assets.rarity=?", info.Rarity)
		}
		return db
	}, &assetsOrders)

	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp find assetsOrders Error: %s", err.Error())
		return out, 0, err
	}
	for _, vAsset := range assetsOrders {
		if vAsset.Id == 0 {
			continue
		}
		//check expireTime
		if !vAsset.ExpireTime.IsZero() && vAsset.ExpireTime.Before(time.Now()) {
			vAsset.Status = common.OrderStatusExpire

			//update order status
			_, err = p.dao.WithContext(info.Ctx).Update(model.Orders{
				Status: common.OrderStatusExpire,
			}, map[string]interface{}{
				model.OrdersColumns.ID:     vAsset.OrderID,
				model.OrdersColumns.Status: common.OrderStatusActive,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp Update Order Status error %s", err.Error())
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
			TokenId:         strconv.FormatInt(vAsset.Id, 10),
			ContractAddress: "0xxxxx",
			ContrainChain:   "BSC",
			Name:            vAsset.Name,
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
			OrderId:         int64(vAsset.OrderID),
			ExpireTime:      vAsset.ExpireTime,
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
