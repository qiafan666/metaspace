package middleware

import (
	"context"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/jau1jz/cornus/commons"
	"github.com/kataras/iris/v12"
)

func Common(ctx iris.Context) {
	//get language
	language := ctx.Request().Header.Get("Language")
	if language == "" {
		language = commons.DefualtLanguage
	}
	ctx.Values().Set(common.BaseRequest, request.BaseRequest{
		Ctx:      ctx.Values().Get("ctx").(context.Context),
		Language: language,
	})

	ctx.Next()
}
