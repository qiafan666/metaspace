package web

import (
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"github.com/kataras/iris/v12/context"
	"sync"
)

// AvatarServiceService service layer interface
type AvatarServiceService interface {
	Token(context *context.Context)
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

func NewAvatarServiceInstance() AvatarServiceService {
	avatarServiceInitOnce.Do(func() {
		avatarServiceIns = &avatarServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return avatarServiceIns
}
