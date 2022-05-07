package api

import (
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/services/api"
	"github.com/jau1jz/cornus/commons"
	"github.com/jau1jz/cornus/commons/utils"
	"github.com/kataras/iris/v12"
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
