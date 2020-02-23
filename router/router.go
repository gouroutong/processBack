package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"fmt"
	//"github.com/kataras/iris/core/router"
)

func Router() *iris.Application {
	app := iris.Default()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	app.StaticWeb("/api/static", "assets")
	router := app.Party("/api/xprocess", crs).AllowMethods(iris.MethodOptions)
	{
		
	}
	fmt.Println(router)
	return app
}
