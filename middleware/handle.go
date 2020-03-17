package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func Handle(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userId:= user["userId"]
	ctx.Values().Set("userId", userId)
	fmt.Println("userId", userId)
	ctx.Next()
}
