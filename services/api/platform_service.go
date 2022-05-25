package api

import (
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"strings"
	"sync"
	"time"
)

// PlatformService service layer interface
type PlatformService interface {
	AddAssets(info request.AddAssets) (out response.AddAssets, code commons.ResponseCode, err error)
}

var PlatformServiceIns *PlatformServiceImp
var PlatformServiceInitOnce sync.Once

func NewPlatformInstance() PlatformService {
	PlatformServiceInitOnce.Do(func() {
		PlatformServiceIns = &PlatformServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return PlatformServiceIns
}

type PlatformServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (p PlatformServiceImp) AddAssets(info request.AddAssets) (out response.AddAssets, code commons.ResponseCode, err error) {

	walletAddress := strings.ToLower(info.WalletAddress)

	tx := p.dao.Tx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	newAssets := model.Assets{
		Uid:         walletAddress,
		Category:    info.Category,
		Type:        info.Type,
		Rarity:      info.Rarity,
		Image:       info.Image,
		Uri:         info.Uri,
		UriContent:  info.UriContent,
		Description: info.Description,
		IsNft:       common.NotNft,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = tx.WithContext(info.Ctx).Create(&newAssets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "PlatformServiceImp assets Create error %s", err.Error())
		return out, 0, err
	}

	//update order status
	_, err = tx.WithContext(info.Ctx).Update(model.Assets{
		TokenId: newAssets.Id,
	}, map[string]interface{}{
		model.AssetsColumns.Id: newAssets.Id,
	}, nil)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "PlatformServiceImp Update assets token_id error %s", err.Error())
		return out, 0, err
	}
	return
}
