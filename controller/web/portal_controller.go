package web

import (
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/services/web"
	"github.com/jau1jz/cornus/commons"
	"github.com/jau1jz/cornus/commons/utils"
	"github.com/kataras/iris/v12"
)

type PortalWebController struct {
	Ctx               iris.Context
	PortalService     web.PortalService
	GameAssetsService web.GameAssetsService
	MarketService     web.MarketService
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
func (receiver *PortalWebController) PostUserAssets() {
	input := request.GetGameAssets{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostUserAssets"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.GameAssetsService.GetGameAssets(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
func (receiver *PortalWebController) PostSubscribeNewsletterEmail() {
	input := request.SubscribeNewsletterEmail{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostSubscribeNewsletterEmail"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.SubscribeNewsletterEmail(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
func (receiver *PortalWebController) PostTowerStatus() {
	input := request.TowerStats{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostTowerStatus"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.GetTowerStatus(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostSign() {
	input := request.Sign{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostSign"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.GetSign(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostOrderSign() {
	input := request.ShelfSign{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostOrderSign"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.MarketService.GetShelfSignature(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostOrderCreate() {
	input := request.SellShelf{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostOrderCreate"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.MarketService.GetSellShelf(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostOrders() {
	input := request.Orders{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostOrders"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.MarketService.GetOrders(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostUserOrders() {
	input := request.Orders{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostUserOrders"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.MarketService.GetUserOrders(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostOrderCancel() {
	input := request.OrderCancel{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostOrderCancel"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.MarketService.OrderCancel(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) GetHealth() {
	receiver.Ctx.StatusCode(iris.StatusOK)
	return
}
