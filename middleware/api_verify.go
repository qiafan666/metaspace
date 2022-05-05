package middleware

import (
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/services/api"
	"github.com/jau1jz/cornus/commons"
	"github.com/kataras/iris/v12"
	"strconv"
	"sync"
)

var signService api.SignService
var signMidOnce sync.Once

var apiWitheList = map[string]string{
	"/metaspace/api/login/third/code": "",
	"/metaspace/api/login":            "",
}

func CheckSignAuth(ctx iris.Context) {

	signMidOnce.Do(func() {
		signService = api.NewSignInstance()
	})

	verifyResult, _, _ := signService.VerifySign(inner.VerifySignRequest{
		Sign:      ctx.Request().Header.Get("sign"),
		ApiKey:    ctx.Request().Header.Get("api_key"),
		Timestamp: ctx.Request().Header.Get("timestamp"),
		Rand:      ctx.Request().Header.Get("rand"),
		Uri:       ctx.Request().Header.Get("url"),
		Parameter: ctx.Values().Get(commons.CtxValueParameter).([]byte),
	})

	if verifyResult.Flag != false {
		_, _ = ctx.JSON(commons.BuildFailed(commons.ValidateError, commons.DefualtLanguage))
		return
	}

	var uuid, email string
	var userId uint64
	//check white list
	if _, ok := apiWitheList[ctx.Request().RequestURI]; !ok {
		//查询redis
		_, err := signService.GetThirdPartyToken(ctx, verifyResult.ThirdPartyId)
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.ValidateError, commons.DefualtLanguage))
			return
		}

		token := ctx.Request().Header.Get("Authorization")
		tokenUser, err := signService.GetTokenUser(ctx, token, strconv.FormatUint(verifyResult.ThirdPartyId, 10))
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
