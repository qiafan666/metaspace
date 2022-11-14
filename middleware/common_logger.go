package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/qiafan666/metaspace/common/function"
	"github.com/qiafan666/metaspace/pojo/inner"
	"github.com/qiafan666/metaspace/services/common"
	"github.com/qiafan666/quickweb/commons"
	"sync"
)

var logService common.LoggerService
var logMidOnce sync.Once

var logList = map[string]string{
	"/metaspace/web/login":      "",
	"/metaspace/api/assets/add": "",
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
	if base, flag := function.GetBaseApiRequest(ctx); flag == true {
		ThirdPartyID = base.BaseThirdPartyId
		UserID = base.BaseUserID
	}
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
