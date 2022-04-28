package middleware

import (
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/services/api"
	"github.com/jau1jz/cornus/commons"
	"github.com/kataras/iris/v12"
	"sync"
)

var signService api.SignService
var signMidOnce sync.Once

var signWitheList = map[string]string{
	"/metaspace/api/test": "",
}

func CheckApiAuth(ctx iris.Context) {
	signMidOnce.Do(func() {
		signService = api.NewSignInstance()
	})

	var language string
	//get language
	language = ctx.Request().Header.Get("Language")
	if language == "" {
		language = commons.DefualtLanguage
	}

	//check white list
	if _, ok := signWitheList[ctx.Request().RequestURI]; !ok {
		//find parameter in request
		sign, _, _ := signService.VerifySign(inner.VerifySignRequest{
			Sign:      ctx.Request().Header.Get("sign"),
			ApiKey:    ctx.Request().Header.Get("api-key"),
			Timestamp: ctx.Request().Header.Get("timestamp"),
			Rand:      ctx.Request().Header.Get("rand"),
			Uri:       ctx.Request().RequestURI,
			Parameter: "",
		})
		if sign.Result != false {
			return
		}
	}
	ctx.Next()
}
