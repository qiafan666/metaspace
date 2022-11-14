package main

import (
	"github.com/qiafan666/metaspace/common"
	router "github.com/qiafan666/metaspace/controller"
	"github.com/qiafan666/quickweb"
	"github.com/qiafan666/quickweb/commons"
)

func main() {
	server := cornus.GetCornusInstance()
	server.Default()
	server.RegisterErrorCodeAndMsg(commons.MsgLanguageEnglish, common.EnglishCodeMsg)
	server.StartServer(cornus.DatabaseService, cornus.RedisService)
	router.RegisterRouter(server.App().GetIrisApp())
	server.WaitClose()
}
