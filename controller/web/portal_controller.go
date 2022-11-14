package web

import (
	"github.com/kataras/iris/v12"
	"github.com/qiafan666/fundametality/common/function"
	"github.com/qiafan666/fundametality/pojo/request"
	"github.com/qiafan666/fundametality/services/web"
	"github.com/qiafan666/quickweb/commons"
	"github.com/qiafan666/quickweb/commons/utils"
	"net/http"
)

type PortalWebController struct {
	Ctx           iris.Context
	PortalService web.PortalService
}

func (receiver *PortalWebController) PostLogin() {
	input := request.UserLogin{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostLogin"); code != commons.OK {
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
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostLoginNonce"); code != commons.OK {
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
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostRegister"); code != commons.OK {
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
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostPasswordUpdate"); code != commons.OK {
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

func (receiver *PortalWebController) GetHealth() {
	receiver.Ctx.StatusCode(iris.StatusOK)
	return
}
