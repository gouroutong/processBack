package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"xProcessBackend/conf"
	"xProcessBackend/router"
	"xProcessBackend/ws"
)

func main() {
	conf.Init()
	configs := conf.GetConfig()
	app := router.Router()
	mvc.Configure(app.Party("/websocket"), ws.ConfigureMVC)
	app.Run(iris.Addr(configs.Server), iris.WithCharset("UTF-8"))
}
