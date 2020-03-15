package controller

import (
	"github.com/kataras/iris/context"
	"xProcessBackend/model"
	"xProcessBackend/serializer"
)

func NewApply(ctx context.Context) {
	var (
		applyService model.Apply
		apply        model.Apply
		err          error
	)
	ctx.ReadJSON(&applyService)
	err = applyService.NewApply(&apply)
	ctx.JSON(serializer.GetResponse(apply, err))
}

func GetApplyItem(ctx context.Context) {
	var (
		applyService model.Apply
		apply        model.Apply
		err          error
	)
	ctx.ReadJSON(&applyService)
	err = applyService.GetApplyItem(&apply)
	ctx.JSON(serializer.GetResponse(apply, err))
	
}