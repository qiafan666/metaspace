package web

import (
	"errors"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/model/join"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
	"sync"
	"time"
)

// AvatarService service layer interface
type AvatarService interface {
	Token(context *context.Context)
	GetAvatar(info request.Avatar) (out response.Avatar, code commons.ResponseCode, err error)
	AvatarDetail(info request.AvatarDetail) (out response.AvatarDetail, code commons.ResponseCode, err error)
}

var avatarConfig struct {
	ETHContract struct {
		Client string `yaml:"monitor_client"`
		Mint   string `yaml:"mint"`
		Assets string `yaml:"assets"`
		Ship   string `yaml:"ship"`
		Market string `yaml:"market"`
		Avatar string `yaml:"avatar"`
	} `yaml:"eth_contract"`
	BSCContract struct {
		Client string `yaml:"monitor_client"`
		Mint   string `yaml:"mint"`
		Assets string `yaml:"assets"`
		Ship   string `yaml:"ship"`
		Market string `yaml:"market"`
	} `yaml:"bsc_contract"`
	Chain struct {
		ETH uint64 `yaml:"eth"`
		BSC uint64 `yaml:"bsc"`
	} `yaml:"chain"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&avatarConfig)
}

var avatarServiceIns *avatarServiceImp
var avatarServiceInitOnce sync.Once

type avatarServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (a avatarServiceImp) Token(context *context.Context) {

	avatarID, err := context.Params().GetInt("id")
	if err != nil {
		_, _ = context.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefualtLanguage))
		slog.Slog.ErrorF(context, "avatarServiceImp getInt error %s", err.Error())
		return
	} else {
		var avatar model.Avatar
		err = a.dao.First([]string{model.AvatarColumns.Content}, map[string]interface{}{
			model.AvatarColumns.AvatarID: avatarID,
		}, nil, &avatar)
		if err != nil {
			slog.Slog.ErrorF(context, "avatarServiceImp get avatar error %s", err.Error())
			_, _ = context.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefualtLanguage))
			slog.Slog.ErrorF(context, "avatarServiceImp getInt error %s", err.Error())
			return
		}

		_, _ = context.Text(string(avatar.Content))
	}
}

func (a avatarServiceImp) GetAvatar(info request.Avatar) (out response.Avatar, code commons.ResponseCode, err error) {

	count, err := a.dao.WithContext(info.Ctx).Count(join.AvatarOrders{}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db = db.Joins("Left JOIN orders_detail ON orders_detail.nft_id = avatar.avatar_id and orders_detail.market_type=? ", common.Avatar).
			Joins("Left JOIN orders ON orders.id = orders_detail.order_id").
			Where("avatar.owner=?", info.BaseWallet)
		return db
	})
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "avatarServiceImp AvatarOrders count error %s", err.Error())
		return out, 0, err
	}

	var avatarOrders []join.AvatarOrders
	err = a.dao.WithContext(info.Ctx).Find([]string{"avatar.id,avatar.avatar_id,avatar.is_shelf,avatar.content,avatar.owner," +
		"orders_detail.price,orders_detail.order_id,orders.start_time,orders.expire_time,orders.updated_time,orders.`status`," +
		"orders.signature,orders.salt_nonce"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db = db.Scopes(Paginate(info.CurrentPage, info.PageCount)).
			Joins("Left JOIN orders_detail ON orders_detail.nft_id = avatar.avatar_id and orders_detail.market_type=? ", common.Avatar).
			Joins("Left JOIN orders ON orders.id = orders_detail.order_id").
			Where("avatar.owner=?", info.BaseWallet)
		db.Order(model.OrdersColumns.UpdatedTime + " desc")
		return db
	}, &avatarOrders)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "avatarServiceImp find avatarOrders Error: %s", err.Error())
		return out, 0, err
	}

	tx := a.dao.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	out.Data = make([]response.AvatarBody, 0, len(avatarOrders))

	for _, avatarOrder := range avatarOrders {
		//check expireTime
		if !avatarOrder.ExpireTime.IsZero() && avatarOrder.ExpireTime.Before(time.Now()) && avatarOrder.Status == common.OrderStatusActive {
			avatarOrder.Status = common.OrderStatusExpire

			//update order status
			_, err = tx.WithContext(info.Ctx).Update(model.Orders{
				Status:      common.OrderStatusExpire,
				UpdatedTime: time.Now(),
			}, map[string]interface{}{
				model.OrdersColumns.ID:     avatarOrder.OrderID,
				model.OrdersColumns.Status: common.OrderStatusActive,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "avatarServiceImp Update Order Status error %s", err.Error())
				return out, 0, err
			}

			//update avatar is_shelf
			_, err = tx.WithContext(info.Ctx).Update(model.Avatar{
				IsShelf: common.NotShelf,
			}, map[string]interface{}{
				model.AvatarColumns.AvatarID: avatarOrder.AvatarID,
			}, nil)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "avatarServiceImp Update avatar is_shelf error %s", err.Error())
				return out, 0, err
			}

			//add expire history
			newTransactionHistory := model.TransactionHistory{
				WalletAddress: info.BaseWallet,
				TokenID:       avatarOrder.AvatarID,
				Price:         avatarOrder.Price,
				OriginChain:   avatarConfig.Chain.ETH,
				Status:        common.Expire,
				MarketType:    common.Avatar,
				UpdatedTime:   time.Now(),
				CreatedTime:   time.Now(),
			}
			err = tx.WithContext(info.Ctx).Create(&newTransactionHistory)
			if err != nil {
				slog.Slog.ErrorF(info.Ctx, "avatarServiceImp TransactionHistory Create error %s", err.Error())
				return out, 0, err
			}

		}

		out.Data = append(out.Data, response.AvatarBody{
			Id:            avatarOrder.ID,
			Owner:         avatarOrder.Owner,
			AvatarID:      avatarOrder.AvatarID,
			IsShelf:       avatarOrder.IsShelf,
			Status:        avatarOrder.Status,
			Content:       string(avatarOrder.Content),
			Price:         avatarOrder.Price,
			OrderId:       avatarOrder.OrderID,
			ExpireTime:    avatarOrder.ExpireTime,
			Signature:     avatarOrder.Signature,
			StartTime:     avatarOrder.StartTime,
			SaltNonce:     avatarOrder.SaltNonce,
			ContractChain: avatarConfig.Chain.ETH,
		})

	}
	out.Total = count
	out.CurrentPage = info.CurrentPage
	out.PrePageCount = info.PageCount

	return
}

func (a avatarServiceImp) AvatarDetail(info request.AvatarDetail) (out response.AvatarDetail, code commons.ResponseCode, err error) {

	var avatarOrder join.AvatarOrders
	err = a.dao.WithContext(info.Ctx).First([]string{"avatar.id,avatar.avatar_id,avatar.is_shelf,avatar.content,avatar.owner," +
		"orders_detail.price,orders_detail.order_id,orders.start_time,orders.expire_time,orders.updated_time,orders.`status`," +
		"orders.signature,orders.salt_nonce"}, map[string]interface{}{}, func(db *gorm.DB) *gorm.DB {
		db = db.Joins("Left JOIN orders_detail ON orders_detail.nft_id = avatar.avatar_id and orders_detail.market_type=? ", common.Avatar).
			Joins("Left JOIN orders ON orders.id = orders_detail.order_id").
			Where("avatar.owner=?", info.BaseWallet).
			Where("avatar.owner != nil").
			Where("avatar.avatar_id=?", info.TokenId)
		return db
	}, &avatarOrder)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Slog.ErrorF(info.Ctx, "avatarServiceImp find avatarOrders Error: %s", err.Error())
		return out, common.AssetsNotExist, errors.New(commons.GetCodeAndMsg(common.AssetsNotExist, info.Language))
	} else if err != nil {
		slog.Slog.ErrorF(info.Ctx, "avatarServiceImp find avatarOrders Error: %s", err.Error())
		return out, 0, err
	}

	out = response.AvatarDetail{
		AssetId:       avatarOrder.ID,
		Owner:         avatarOrder.Owner,
		AvatarID:      avatarOrder.AvatarID,
		Status:        avatarOrder.Status,
		Content:       string(avatarOrder.Content),
		Price:         avatarOrder.Price,
		ExpireTime:    avatarOrder.ExpireTime,
		Signature:     avatarOrder.Signature,
		StartTime:     avatarOrder.StartTime,
		SaltNonce:     avatarOrder.SaltNonce,
		ContractChain: avatarConfig.Chain.ETH,
		IsNft:         common.IsNft,
	}
	return
}

func NewAvatarServiceInstance() AvatarService {
	avatarServiceInitOnce.Do(func() {
		avatarServiceIns = &avatarServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return avatarServiceIns
}
