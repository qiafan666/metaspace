package middleware

import (
	"context"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/services/api"
	"github.com/jau1jz/cornus/commons"
	"github.com/kataras/iris/v12"
	"net/http"
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

	var parameter []byte
	var apiKey, rand, url string
	apiKey = ctx.Request().Header.Get("api_key")
	rand = ctx.Request().Header.Get("rand")
	url = ctx.Request().RequestURI

	if ctx.Request().Method == http.MethodPost {

		body, err := ctx.Request().GetBody()
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefualtLanguage))
			return
		}
		_, err = body.Read(parameter)
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefualtLanguage))
			return
		}
	} else if ctx.Request().Method == http.MethodGet {
		_, _ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefualtLanguage))
		return
	}

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

	ctx.Values().Set(common.BaseRequest, request.BaseRequest{
		Ctx:       ctx.Values().Get("ctx").(context.Context),
		Language:  language,
		BaseUUID:  uuid,
		BaseEmail: email,
	})

	ctx.Next()
}
