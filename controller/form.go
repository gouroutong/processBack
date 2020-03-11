package controller

import (
	"github.com/kataras/iris/context"
	"xProcessBackend/model"
	"xProcessBackend/serializer"
)

func NewOrUpdate(ctx context.Context) {
	var (
		formService model.Form
		form        model.Form
		err         error
	)
	ctx.ReadJSON(&formService)
	err = formService.NewOrUpdate(&form)
	ctx.JSON(serializer.GetResponse(form, err))
}
func GetItem(ctx context.Context) {
	var (
		formService model.Form
		form        model.Form
		err         error
	)
	ctx.ReadJSON(&formService)
	err = formService.GetItem(&form)
	ctx.JSON(serializer.GetResponse(form, err))
}

func GetList(ctx context.Context) {
	var (
		formService model.Form
		list        []model.Form
		err         error
	)
	ctx.ReadJSON(&formService)
	err = model.GetList(&list)
	ctx.JSON(serializer.GetResponse(list, err))
}
