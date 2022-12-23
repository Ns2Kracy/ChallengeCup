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
	v1 := r.Party("/api/v1")
	v1.Post("/user/login", controller.PostUserLogin)
	v1.Post("/user/register-by-name", controller.PostUserRegisterByUserNameAndPassword)
	v1.Post("/user/register-by-phone", controller.PostUserRegisterByPhone)
	v1.Post("/user/register-by-email", controller.PostUserRegisterByEmail)
	group := r.Party(v1.GetRelPath())
	group.Use(middleware.JwtAuthMiddleware)
	{
		group.Get("/user/test", func(ctx iris.Context) {
			ctx.JSON(model.Result{
				Code:    common.SUCCESS,
				Message: common.Message(common.SUCCESS),
			})
		})
	}
}
