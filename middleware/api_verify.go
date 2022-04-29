package middleware

import (
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/common/function"
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
	"/metaspace/api/login/authcode": "",
	"/metaspace/api/login":          "",
}

func CheckSignAuth(ctx iris.Context) {

	signMidOnce.Do(func() {
		signService = api.NewSignInstance()
	})

	sign := function.Byte2([]byte(ctx.Request().Header.Get("sign")))

	timestamp, err := strconv.ParseInt(ctx.Request().Header.Get("timestamp"), 10, 64)
	if err != nil {
		_, _ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefualtLanguage))
		return
	}

	parameter := ctx.Request().Header.Get(commons.CtxValueParameter)
	var apiKey, rand, url string
	apiKey = ctx.Request().Header.Get("api_key")
	rand = ctx.Request().Header.Get("rand")
	url = ctx.Request().RequestURI

	result, _, _ := signService.VerifySign(inner.VerifySignRequest{
		Sign:      sign,
		ApiKey:    apiKey,
		Timestamp: timestamp,
		Rand:      rand,
		Uri:       url,
		Parameter: parameter,
	})

	if result.Flag != false {
		_, _ = ctx.JSON(commons.BuildFailed(commons.ValidateError, commons.DefualtLanguage))
		return
	}

	ctx.Values().Set(commons.CtxValueParameter, parameter)

	var uuid, email string
	var userId uint64
	//check white list
	if _, ok := apiWitheList[ctx.Request().RequestURI]; !ok {
		//查询redis
		thirdPartyToken, err := signService.GetThirdPartyToken(ctx, result.ThirdPartyId)
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.ValidateError, commons.DefualtLanguage))
			return
		}

		tokenUser, err := signService.GetTokenUser(ctx, thirdPartyToken.Token)
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.ValidateError, commons.DefualtLanguage))
			return
		}
		uuid = tokenUser.Uuid
		email = tokenUser.Email
		userId = tokenUser.UserId
	}

	ctx.Values().Set(common.BaseApiRequest, request.BaseApiRequest{
		BaseUUID:     uuid,
		BaseEmail:    email,
		BaseUserID:   userId,
		ThirdPartyId: result.ThirdPartyId,
	})

	ctx.Next()
}
