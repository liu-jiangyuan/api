package server

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liu-jiangyuan/api/app/api/controller"
	"math/big"
)

func InitRoute() *iris.Application {
	app := iris.New()

	app.Use(recover.New())
	app.Logger().SetLevel("debug")

	app.OnErrorCode(iris.StatusNotFound, func (ctx iris.Context) {
		ctx.StopExecution()
		_, _ = ctx.JSON(struct {
			Code int
			Msg string
		}{404,"Not Found"})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func (ctx iris.Context) {
		ctx.StopExecution()
		_, _ = ctx.JSON(struct {
			Code int
			Msg string
		}{500,"Internal Server Error"})
	})

	api := app.Party("/api",func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})
	{


		//mvc.Configure(api, func(mvcConfigure *mvc.Application) {
		//	mvcConfigure.Register(
		//		big.NewInt(66),
		//	)
		//	mvcConfigure.Handle(new(controller.Index))
		//})

		mvc.Configure(api.Party("/index"), func(mvcConfigure *mvc.Application) {
			mvcConfigure.Register(
				big.NewInt(76),
			)
			mvcConfigure.Handle(new(controller.Index))
		})
	}

	return app
}
