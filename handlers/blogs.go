package handlers

import (
	"sauceBlog/db"
	"sauceBlog/utils"

	"github.com/kataras/iris/v12"
)

func Blogs(ctx iris.Context){
	blogs := db.Blogs
	status := utils.StatusCodeMessage[200]

	rsp := utils.Response{
		Status: status,
		Data: blogs,
	}
	ctx.JSON(rsp)
}