package api

import (
	"github.com/kataras/iris/v12"
	"github.com/qiafan666/metaspace/common/function"
	"github.com/qiafan666/metaspace/pojo/request"
	"github.com/qiafan666/metaspace/services/api"
	"github.com/qiafan666/quickweb/commons"
	"github.com/qiafan666/quickweb/commons/utils"
)

type LoginApiController struct {
	Ctx          iris.Context
	LoginService api.LoginService
}

func (receiver *LoginApiController) PostLoginThirdCode() {
	input := request.CreateAuthCode{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "LoginApiController PostLoginThirdCode"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.LoginService.CreateAuthCode(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
