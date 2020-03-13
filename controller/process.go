package controller

import (
	"github.com/kataras/iris/context"
	"xProcessBackend/model"
	"xProcessBackend/serializer"
)

func NewOrUpdateProcess(ctx context.Context) {
	var (
		processService model.Process
		process        model.Process
		err            error
	)
	ctx.ReadJSON(&processService)
	err = processService.NewOrUpdateProcess(&process)
	ctx.JSON(serializer.GetResponse(process, err))
}
func GetProcessItem(ctx context.Context) {
	var (
		processService model.Process
		process        model.Process
		err            error
	)
	ctx.ReadJSON(&processService)
	err = processService.GetProcessItem(&process)
	ctx.JSON(serializer.GetResponse(process, err))
}

func GetProcessList(ctx context.Context) {
	var (
		processService model.Process
		processList    []model.Process
		err            error
	)
	ctx.ReadJSON(&processService)
	err = model.GetProcessList(&processList)
	ctx.JSON(serializer.GetResponse(processList, err))
}
func DeleteProcessItem(ctx context.Context) {
	var (
		processService model.Process
		process        model.Process
		err            error
	)
	ctx.ReadJSON(&processService)
	err = processService.DeleteProcessItem(&process)
	ctx.JSON(serializer.GetResponse(process, err))
}
