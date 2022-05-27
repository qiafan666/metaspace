package common

import (
	"context"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"sync"
)

// PublicService service layer interface
type PublicService interface {
	GetUserId(ctx context.Context, uuid string) (out inner.UserId, code commons.ResponseCode, err error)
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

func (p PublicServiceImp) GetUserId(ctx context.Context, uuid string) (out inner.UserId, code commons.ResponseCode, err error) {
	var user model.User
	err = p.dao.First([]string{model.UserColumns.ID}, map[string]interface{}{
		model.UserColumns.UUID: uuid,
	}, nil, &user)

	if err != nil {
		slog.Slog.ErrorF(ctx, "SignServiceImp GetUserId error %s", err.Error())
		return out, 0, err
	}
	out.UserId = user.ID
	return
}
