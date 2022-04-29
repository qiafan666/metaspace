package common

import (
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus/commons"
	"sync"
)

// LoggerService service layer interface
type LoggerService interface {
	Write(info inner.LogWriteRequest) (out inner.LogWriteResponse, code commons.ResponseCode, err error)
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

func (l LogServiceImp) Write(info inner.LogWriteRequest) (out inner.LogWriteResponse, code commons.ResponseCode, err error) {
	//TODO implement me
	panic("implement me")
}
