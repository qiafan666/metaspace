package middleware

import (
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/services/common"
	"github.com/jau1jz/cornus/commons"
	"github.com/kataras/iris/v12"
	"sync"
)

var logService common.LoggerService
var logMidOnce sync.Once

var logList = map[string]string{
	"/metaspace/web/login": "",
}

func Logger(ctx iris.Context) {
	logMidOnce.Do(func() {
		logService = common.NewSignInstance()
	})
	var UserID, ThirdPartyID uint64
	baseRequest, _ := function.GetBaseRequest(ctx)
	if base, flag := function.GetBasePortalRequest(ctx); flag == true {
		UserID = base.BaseUserID
	}
	//if base, flag := function.GetBaseApiRequest(ctx); flag == true {
	//	//ThirdPartyID = base.ApiKey
	//}
	//check white list
	if _, ok := logList[ctx.Request().RequestURI]; ok {
		logService.Write(inner.LogWriteRequest{
			Ctx:          baseRequest.Ctx,
			Uri:          ctx.Request().RequestURI,
			UserID:       UserID,
			ThirdPartyID: ThirdPartyID,
			Parameter:    ctx.Values().Get(commons.CtxValueParameter).([]byte),
		})
	}
	ctx.Next()
}
