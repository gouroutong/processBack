package userHandle

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"xProcessBackend/model"
	"xProcessBackend/serializer"
)

func New(ctx context.Context) {
	var (
		user model.User
	)
	ctx.ReadJSON(&user)
	ctx.JSON(serializer.GetResponse(iris.Map{
		"user": user,
	}, user.New()))
}

func Login(ctx context.Context) {
	var (
		user     model.User
		userWrap model.UserWrap
		//err      error
	)
	ctx.ReadJSON(&user)
	err := user.Login(&userWrap)
	ctx.JSON(serializer.GetResponse(userWrap, err))
}
