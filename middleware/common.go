package middleware

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/qiafan666/fundametality/common"
	"github.com/qiafan666/fundametality/pojo/request"
	"github.com/qiafan666/quickweb/commons"
	"io"
	"net/http"
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
	if ctx.Request().Method == http.MethodPost {

		body, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefualtLanguage))
			return
		}

		ctx.Values().Set(commons.CtxValueParameter, body)
	}
	ctx.Next()
}
