package web

import (
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	bizservice "github.com/blockfishio/metaspace-backend/services"
	"github.com/jau1jz/cornus/commons"
	"github.com/jau1jz/cornus/commons/utils"
	"github.com/kataras/iris/v12"
)

type PortalWebController struct {
	Ctx           iris.Context
	PortalService bizservice.PortalService
}

func (receiver *PortalWebController) PostLogin() {
	input := request.UserLogin{}
	if code, msg := utils.ValidateAndBindParameters(&input, receiver.Ctx, "PortalWebController PostLogin"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.Login(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
func (receiver *PortalWebController) PostLoginNonce() {
	input := request.GetNonce{}
	if code, msg := utils.ValidateAndBindParameters(&input, receiver.Ctx, "PortalWebController PostLoginNonce"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.GetNonce(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
func (receiver *PortalWebController) PostRegister() {
	input := request.RegisterUser{}
	if code, msg := utils.ValidateAndBindParameters(&input, receiver.Ctx, "PortalWebController PostRegister"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.Register(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
func (receiver *PortalWebController) PostPasswordUpdate() {
	input := request.PasswordUpdate{}
	if code, msg := utils.ValidateAndBindParameters(&input, receiver.Ctx, "PortalWebController PostPasswordUpdate"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.UpdatePassword(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
