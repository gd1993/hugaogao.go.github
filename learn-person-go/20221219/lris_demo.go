package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Use(myMiddleWare)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
	app.Run(iris.Addr(":8080"))

}

func myMiddleWare(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
