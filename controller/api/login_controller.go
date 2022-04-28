package api

import (
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/jau1jz/cornus/commons"
	"github.com/jau1jz/cornus/commons/utils"
	"github.com/kataras/iris/v12"
)

type LoginApiController struct {
	Ctx iris.Context
}

func (receiver *LoginApiController) PostLogin() {
	input := request.UserLogin{}
	if code, msg := utils.ValidateAndBindParameters(&input, receiver.Ctx, "LoginApiController PostLogin"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	//if out, code, err := receiver.PortalService.Login(input); err != nil {
	//	_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	//} else {
	//	_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	//}
}
