package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/qiafan666/metaspace/controller/api"
	"github.com/qiafan666/metaspace/controller/web"
	"github.com/qiafan666/metaspace/middleware"
	api2 "github.com/qiafan666/metaspace/services/api"
	web2 "github.com/qiafan666/metaspace/services/web"
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
				AvatarService:     web2.NewAvatarServiceInstance(),
			})
		})
	//api router
	mvc.Configure(ctx.Party("/metaspace/api"),
		func(application *mvc.Application) {
			application.Router.Use(middleware.CheckSignAuth, middleware.Logger)
			application.Handle(&api.LoginApiController{
				LoginService: api2.NewLoginInstance(),
			})
			application.Handle(&api.PlatformController{
				PlatformService: api2.NewPlatformInstance(),
			})
		})
	//avatar router
	ctx.Get("/avatar/web/token/{id:int}", web2.NewAvatarServiceInstance().Token)
}
