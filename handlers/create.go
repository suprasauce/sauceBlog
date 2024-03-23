package handlers

import (
	"sauceBlog/utils"

	"github.com/kataras/iris/v12"
)

func Create(ctx iris.Context){
	var blog map[string]interface{}
	status := utils.StatusCodeMessage[201]
	ctx.ReadJSON(&blog)
	
	rsp := utils.Response{
		Status: status,
		Data: "",
	}
	ctx.JSON(rsp)
}