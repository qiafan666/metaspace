package middleware

import (
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/services/api"
	"github.com/jau1jz/cornus/commons"
	"github.com/kataras/iris/v12"
	"sync"
)

var signService api.SignService
var signMidOnce sync.Once

var apiWitheList = map[string]string{
	"/metaspace/api/login/third/code": "",
	"/metaspace/api/assets/add":       "",
}

func CheckSignAuth(ctx iris.Context) {

	signMidOnce.Do(func() {
		signService = api.NewSignInstance()
	})

	verifyResult, code, err := signService.VerifySign(inner.VerifySignRequest{
		Sign:      ctx.Request().Header.Get(common.BaseRequestSign),
		ApiKey:    ctx.Request().Header.Get(common.BaseRequestApiKey),
		Timestamp: ctx.Request().Header.Get(common.BaseRequestTimestamp),
		Rand:      ctx.Request().Header.Get(common.BaseRequestRand),
		Uri:       ctx.Request().RequestURI,
		Parameter: ctx.Values().Get(commons.CtxValueParameter).([]byte),
	})

	if err != nil {
		_, _ = ctx.JSON(commons.BuildFailed(code, commons.DefualtLanguage))
		return
	}

	var uuid, email string
	var userId uint64
	//check white list
	if _, ok := apiWitheList[ctx.Request().RequestURI]; !ok {
		//查询redis

		token := ctx.Request().Header.Get(common.BaseRequestAuthorization)
		tokenUser, _, err := signService.GetTokenUser(ctx, token)
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.ValidateError, commons.DefualtLanguage))
			return
		}
		uuid = tokenUser.Uuid
		email = tokenUser.Email
		userId = tokenUser.UserId
	}

	ctx.Values().Set(common.BaseApiRequest, request.BaseApiRequest{
		BaseUUID:         uuid,
		BaseEmail:        email,
		BaseUserID:       userId,
		BaseThirdPartyId: verifyResult.ThirdPartyId,
	})

	ctx.Next()
}
