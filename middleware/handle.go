package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func Handle(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	ctx.Values().Set("userId", user["userId"])
	ctx.Values().Set("username", user["username"])
	fmt.Println("userId", user["userId"])
	ctx.Next()
}
