package main

import (
	"github.com/kataras/iris/v12"
	"github.com/liu-jiangyuan/api/server"
)

func main() {
	app := server.InitRoute()

	_ = app.Run(iris.Addr(":9601"), iris.WithoutServerError(iris.ErrServerClosed))
}
