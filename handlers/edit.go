package handlers

import (
	"sauceBlog/utils"

	"github.com/kataras/iris/v12"
)

func Edit(ctx iris.Context){
	var blog map[string]interface{}
	status := utils.StatusCodeMessage[204]
	ctx.ReadJSON(&blog)

	rsp := utils.Response{
		Status: status,
		Data: "",
	}
	ctx.JSON(rsp)
}