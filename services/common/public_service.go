package common

import (
	"context"
	"github.com/qiafan666/metaspace/dao"
	"github.com/qiafan666/metaspace/model"
	"github.com/qiafan666/metaspace/pojo/inner"
	"github.com/qiafan666/metaspace/redis"
	"github.com/qiafan666/quickweb/commons"
	slog "github.com/qiafan666/quickweb/commons/log"
	"sync"
)

// PublicService service layer interface
type PublicService interface {
	GetUser(ctx context.Context, uuid string) (out inner.User, code commons.ResponseCode, err error)
}

var PublicServiceIns *PublicServiceImp
var PublicServiceInitOnce sync.Once

func NewPublicInstance() PublicService {
	PublicServiceInitOnce.Do(func() {
		PublicServiceIns = &PublicServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return PublicServiceIns
}

type PublicServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (p PublicServiceImp) GetUser(ctx context.Context, uuid string) (out inner.User, code commons.ResponseCode, err error) {
	var user model.User
	err = p.dao.First([]string{model.UserColumns.ID, model.UserColumns.WalletAddress}, map[string]interface{}{
		model.UserColumns.UUID: uuid,
	}, nil, &user)

	if err != nil {
		slog.Slog.ErrorF(ctx, "SignServiceImp GetUser error %s", err.Error())
		return out, 0, err
	}
	out.UserId = user.ID
	out.WalletAddress = user.WalletAddress
	return
}
