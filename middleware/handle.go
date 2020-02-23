package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func Handle(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)
	fmt.Println(user)

	ctx.Next()
}
