package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/qiafan666/fundametality/controller/web"
	"github.com/qiafan666/fundametality/middleware"
)

func RegisterRouter(ctx *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	ctx.Use(middleware.Common)
	//web router
	mvc.Configure(ctx.Party("/metaspace/web", crs).AllowMethods(iris.MethodOptions),
		func(application *mvc.Application) {
			application.Router.Use(middleware.CheckPortalAuth)
			application.Handle(&web.PortalWebController{
				PortalService: web.NewPortalServiceInstance(),
			})
		})
}
