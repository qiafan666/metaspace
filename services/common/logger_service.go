package common

import (
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus/commons/utils"
	"sync"
	"time"
)

// LoggerService service layer interface
type LoggerService interface {
	Write(info inner.LogWriteRequest)
}

var LogServiceIns *LogServiceImp
var signServiceInitOnce sync.Once

func NewSignInstance() LoggerService {
	signServiceInitOnce.Do(func() {
		LogServiceIns = &LogServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return LogServiceIns
}

type LogServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (l LogServiceImp) Write(info inner.LogWriteRequest) {
	utils.Go(func() {
		_ = l.dao.Create(&model.RequestLog{
			ThirdPartyID: info.ThirdPartyID,
			UserID:       info.UserID,
			URI:          info.Uri,
			Parameter:    info.Parameter,
			CreatedTime:  time.Now(),
			UpdatedTime:  time.Now(),
		})
	})
}
