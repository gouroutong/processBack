package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
	userHandle "xProcessBackend/controller"
	"xProcessBackend/middleware"
	"xProcessBackend/serializer"

	//"github.com/kataras/iris/core/router"
)

var mySecret = []byte("My Secret")

func Router() *iris.Application {
	app := iris.Default()
	jwtConfig := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
		Extractor: func(ctx context.Context) (s string, e error) {
			authHeader := ctx.GetHeader("Authorization")
			if authHeader == "" {
				return "", nil // No error, just no token
			}
			return authHeader, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	app.StaticWeb("/api/static", "assets")
	router1 := app.Party("/api", crs).
		AllowMethods(iris.MethodOptions)
	{
		router1.PartyFunc("/user", func(user router.Party) {
			user.Post("/new", userHandle.New)
			user.Post("/login", userHandle.Login)
		})
		router1.Use(jwtConfig.Serve, middleware.Handle)
		router1.PartyFunc("/home", func(home router.Party) {
			home.Get("/", func(ctx context.Context) {
				ctx.JSON(serializer.GetResponse("home", nil))
			})
		})
	}
	return app
}
