package main

import (
	"github.com/kataras/iris/v12"
	"sauceBlog/handlers"
	"sauceBlog/logger"
)

func main(){
	ac := logger.MakeAccessLog()
	defer ac.Close()

	app := iris.New()
	app.UseRouter(ac.Handler)

	app.Get("/blogs", handlers.Blogs)

	app.Post("/create", handlers.Create)

	app.Put("/edit", handlers.Edit)

	app.Listen("localhost:8080")
}