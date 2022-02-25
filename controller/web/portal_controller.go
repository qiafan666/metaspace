package web

import (
	"context"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	bizservice "github.com/blockfishio/metaspace-backend/services"
	"github.com/dgrijalva/jwt-go"
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
	input.Ctx = receiver.Ctx.Value("ctx").(context.Context)
	if out, code, err := receiver.PortalService.Login(input); err != nil {
		if code == 0 {
			_, _ = receiver.Ctx.JSON(commons.BuildFailed(commons.UnKnowError))
		} else {
			_, _ = receiver.Ctx.JSON(commons.BuildFailed(code))
		}
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out))
	}
}
func (receiver *PortalWebController) PostRegister() {
	input := request.RegisterUser{}
	if code, msg := utils.ValidateAndBindParameters(&input, receiver.Ctx, "PortalWebController PostRegister"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	input.Ctx = receiver.Ctx.Value("ctx").(context.Context)
	if out, code, err := receiver.PortalService.Register(input); err != nil {
		if code == 0 {
			_, _ = receiver.Ctx.JSON(commons.BuildFailed(commons.UnKnowError))
		} else {
			_, _ = receiver.Ctx.JSON(commons.BuildFailed(code))
		}
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out))
	}
}
func (receiver *PortalWebController) PostPasswordUpdate() {
	input := request.PasswordUpdate{}
	if code, msg := utils.ValidateAndBindParameters(&input, receiver.Ctx, "PortalWebController PostPasswordUpdate"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	claims, ok := receiver.Ctx.Value("claims").(jwt.MapClaims)
	if ok != true {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(commons.UnKnowError))
		return
	}
	if input.UUID, ok = claims["uuid"].(string); !ok || input.UUID == "" {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(commons.UnKnowError))
		return
	}
	if input.Account, ok = claims["account"].(string); !ok {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(commons.UnKnowError))
		return
	}
	input.Ctx = receiver.Ctx.Value("ctx").(context.Context)
	if out, code, err := receiver.PortalService.UpdatePassword(input); err != nil {
		if code == 0 {
			_, _ = receiver.Ctx.JSON(commons.BuildFailed(commons.UnKnowError))
		} else {
			_, _ = receiver.Ctx.JSON(commons.BuildFailed(code))
		}
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out))
	}
}
