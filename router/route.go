package router

import (
	"ChallengeCup/common"
	"ChallengeCup/controller"
	"ChallengeCup/middleware"
	"ChallengeCup/model"

	"github.com/kataras/iris/v12"
)

func InitRoute(r *iris.Application) {
	r.UseGlobal(middleware.Cors())
	r.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to Challenge Cup</h1>")
	})
	r.Post("/api/user/login", controller.PostUserLogin)
	r.Post("/api/user/register", controller.PostUserRegisterByUserNameAndPassword)
	group := r.Party("/api")
	group.Use(middleware.JWT())
	{
		group.Get("/user", func(ctx iris.Context) {
			ctx.JSON(model.Result{
				Code:    common.SUCCESS,
				Message: common.Message(common.SUCCESS),
			})
		})
	}
}
