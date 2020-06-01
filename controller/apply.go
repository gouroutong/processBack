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
	userId := ctx.Values().Get("userId").(float64)
	if applyService.Id == 0 {
		username := ctx.Values().Get("username").(string)
		applyService.StartUserId = int64(userId)
		applyService.StartUsername = username
	}
	err = applyService.NewApply(&apply, int64(userId))
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

func GetApplyList(ctx context.Context) {
	var (
		applyList []model.Apply
		err       error
	)
	userId := ctx.Values().Get("userId").(float64)
	err = model.GetApplyList(&applyList, int64(userId))
	ctx.JSON(serializer.GetResponse(applyList, err))

}

func GetAuditApplyList(ctx context.Context) {
	var (
		applyList []model.Apply
		err       error
	)
	userId := ctx.Values().Get("userId").(float64)
	err = model.GetAuditApplyList(&applyList, int64(userId))
	ctx.JSON(serializer.GetResponse(applyList, err))

}

func DeleteApplyItem(ctx context.Context) {

	var (
		applyService model.Apply
		apply        model.Apply
		err          error
	)
	ctx.ReadJSON(&applyService)
	err = applyService.DeleteApplyItem(&apply)
	ctx.JSON(serializer.GetResponse(apply, err))
}
