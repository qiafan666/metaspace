package router

import (
	"github.com/blockfishio/metaspace-backend/controller/web"
	"github.com/blockfishio/metaspace-backend/middleware"
	bizservice "github.com/blockfishio/metaspace-backend/services"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterRouter(ctx *iris.Application) {
	crs := cors.New(cors.Options{
		//AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		//AllowCredentials: true,
		//AllowedHeaders:   []string{"*"},
	})
	mvc.Configure(ctx.Party("/metaspace/web", crs).AllowMethods(iris.MethodOptions),
		func(application *mvc.Application) {
			application.Router.Use(middleware.CheckAuth)
			//application.Router.Use(jwtHandler.Serve)
			application.Handle(&web.PortalWebController{
				PortalService:     bizservice.NewPortalServiceInstance(),
				GameAssetsService: bizservice.NewGameAssetsInstance(),
			})
		})
}
