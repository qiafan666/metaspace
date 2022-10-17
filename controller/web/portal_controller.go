package web

import (
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/services/web"
	"github.com/jau1jz/cornus/commons"
	"github.com/jau1jz/cornus/commons/utils"
	"github.com/kataras/iris/v12"
	"net/http"
)

type PortalWebController struct {
	Ctx               iris.Context
	PortalService     web.PortalService
	GameAssetsService web.GameAssetsService
	MarketService     web.MarketService
	AvatarService     web.AvatarService
}

func (receiver *PortalWebController) PostLoginThird() {
	input := request.ThirdPartyLogin{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "LoginApiController PostLoginThird"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.ThirdPartyLogin(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		receiver.Ctx.Redirect(out.Url, http.StatusFound)
	}
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
	if out, code, err := receiver.MarketService.SellShelf(input); err != nil {
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

func (receiver *PortalWebController) PostUserUpdate() {
	input := request.UserUpdate{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostUserUpdate"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.UserUpdate(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostUserHistory() {
	input := request.UserHistory{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostUserUpdate"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.UserHistory(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostExchangePrice() {
	input := request.ExchangePrice{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostExchangePrice"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.ExchangePrice(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostAssetDetail() {
	input := request.AssetDetail{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostAssetDetail"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.AssetDetail(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostGameCurrency() {
	input := request.GameCurrency{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostGameCurrency"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.GameCurrency(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
func (receiver *PortalWebController) PostOrdersGroup() {
	input := request.OrdersGroup{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostOrdersGroup"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.MarketService.GetOrdersGroup(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostOrdersDetail() {
	input := request.OrdersGroupDetail{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostOrdersGroupDetail"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.MarketService.GetOrdersGroupDetail(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostSendCode() {
	input := request.SendCode{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostSendCode"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.SendCode(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostPaperMint() {
	input := request.PaperMint{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostPaperMint"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.PaperMint(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostPaperTransaction() {
	input := request.PaperTransaction{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostPaperTransaction"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.PaperTransaction(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostOrdersOfficial() {
	input := request.OrdersOfficial{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostOrdersOfficial"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.MarketService.GetOrdersOfficial(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostUserAvatar() {
	input := request.Avatar{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostUserAvatar"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.AvatarService.GetAvatar(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostOrderAvatar() {
	input := request.OrderAvatar{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostOrderAvatar"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.MarketService.OrderAvatar(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostAvatarDetail() {
	input := request.AvatarDetail{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostAvatarDetail"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.AvatarService.AvatarDetail(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}

func (receiver *PortalWebController) PostTest() {
	input := request.Test{}
	if code, msg := utils.ValidateAndBindCtxParameters(&input, receiver.Ctx, "PortalWebController PostTest"); code != commons.OK {
		_, _ = receiver.Ctx.JSON(commons.BuildFailedWithMsg(code, msg))
		return
	}
	function.BindBaseRequest(&input, receiver.Ctx)
	if out, code, err := receiver.PortalService.Test(input); err != nil {
		_, _ = receiver.Ctx.JSON(commons.BuildFailed(code, input.Language))
	} else {
		_, _ = receiver.Ctx.JSON(commons.BuildSuccess(out, input.Language))
	}
}
func (receiver *PortalWebController) GetHealth() {
	receiver.Ctx.StatusCode(iris.StatusOK)
	return
}
