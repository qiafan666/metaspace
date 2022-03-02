package main

import (
	"github.com/blockfishio/metaspace-backend/common"
	router "github.com/blockfishio/metaspace-backend/controller"
	"github.com/jau1jz/cornus"
	"github.com/jau1jz/cornus/commons"
)

func main() {
	server := cornus.GetCornusInstance()
	server.Default()
	server.RegisterErrorCodeAndMsg(commons.MsgLanguageEnglish, common.EnglishCodeMsg)
	server.StartServer(cornus.DatabaseService, cornus.RedisService)
	router.RegisterRouter(server.App().GetIrisApp())
	server.WaitClose()
}
