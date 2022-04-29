package router

import (
	"github.com/blockfishio/metaspace-backend/controller/api"
	"github.com/blockfishio/metaspace-backend/controller/web"
	"github.com/blockfishio/metaspace-backend/middleware"
	web2 "github.com/blockfishio/metaspace-backend/services/web"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
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
			application.Router.Use(middleware.CheckPortalAuth, middleware.Logger)
			application.Handle(&web.PortalWebController{
				PortalService:     web2.NewPortalServiceInstance(),
				GameAssetsService: web2.NewGameAssetsInstance(),
				MarketService:     web2.NewMarketInstance(),
			})
		})
	//api router
	mvc.Configure(ctx.Party("/metaspace/api"),
		func(application *mvc.Application) {
			application.Router.Use(middleware.CheckSignAuth, middleware.CheckApiAuth, middleware.Logger)
			application.Handle(&api.LoginApiController{})
		})
}
