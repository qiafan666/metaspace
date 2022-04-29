package middleware

import (
	"github.com/blockfishio/metaspace-backend/services/common"
	"github.com/kataras/iris/v12"
	"sync"
)

var logService common.LoggerService
var logMidOnce sync.Once

var logList = map[string]string{
	"/metaspace/api/test": "",
}

func Logger(ctx iris.Context) {
	logMidOnce.Do(func() {
		logService = common.NewSignInstance()
	})
	//baseRequest := function.GetBaseRequest(ctx)
	//baseApiRequest := function.GetBaseApiRequest(ctx)
	//ctx.Values().Get()
	//check white list
	if _, ok := logList[ctx.Request().RequestURI]; ok {
		//logService.Write(inner.LogWriteRequest{
		//	Uri:          ctx.Request().RequestURI,
		//	UserID:       baseRequest.BaseUserID,
		//	ThirdPartyID: baseApiRequest.BaseUUID,
		//	Parameter:    nil,
		//})
	}
	ctx.Next()
}
