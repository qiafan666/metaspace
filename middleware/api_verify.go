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

	flag, _, _ := signService.VerifySign(inner.VerifySignRequest{
		Sign:      sign,
		ApiKey:    apiKey,
		Timestamp: timestamp,
		Rand:      rand,
		Uri:       url,
		Parameter: string(parameter),
	})

	if flag.Result != false {
		_, _ = ctx.JSON(commons.BuildFailed(commons.ValidateError, commons.DefualtLanguage))
		return
	}

	ctx.Values().Set(commons.CtxValueParameter, parameter)
	ctx.Values().Set(common.BaseApiRequest, request.BaseApiRequest{
		ApiKey: apiKey,
	})

	ctx.Next()
}

func CheckApiAuth(ctx iris.Context) {
	signMidOnce.Do(func() {
		signService = api.NewSignInstance()
	})

	var language, uuid, email string
	//get language
	language = ctx.Request().Header.Get("Language")
	if language == "" {
		language = commons.DefualtLanguage
	}
	//check white list
	if _, ok := apiWitheList[ctx.Request().RequestURI]; !ok {

	}

	ctx.Values().Set(common.BasePortalRequest, request.BasePortalRequest{
		BaseUUID:  uuid,
		BaseEmail: email,
	})

	ctx.Next()
}
