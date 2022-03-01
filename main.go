package main

import (
	"github.com/blockfishio/metaspace-backend/common"
	router "github.com/blockfishio/metaspace-backend/controller"
	"github.com/jau1jz/cornus"
)

func main() {
	server := cornus.GetCornusInstance()
	server.Default()
	server.RegisterErrorCodeAndMsg(common.CodeMsg)
	server.StartServer(cornus.DatabaseService, cornus.RedisService)
	router.RegisterRouter(server.App().GetIrisApp())
	server.WaitClose()
}
