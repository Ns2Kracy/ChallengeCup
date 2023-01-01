package router

import (
	"ChallengeCup/controller"
	"ChallengeCup/middleware"

	"github.com/kataras/iris/v12"
)

func InitRoute(r *iris.Application) {
	r.UseGlobal(middleware.Cors())
	v1 := r.Party("/api/v1")
	iris.RegisterOnInterrupt(middleware.Monitor.Stop)
	v1.Post("/monitor", middleware.Monitor.Stats)
	v1.Get("/monitor", middleware.Monitor.View)
	
	v1.Post("/user/login", controller.PostUserLogin)
	v1.Post("/user/register-by-name", controller.PostUserRegisterByUserNameAndPassword)
	v1.Post("/user/register-by-phone", controller.PostUserRegisterByPhone)
	v1.Post("/user/register-by-email", controller.PostUserRegisterByEmail)
	group := r.Party(v1.GetRelPath())
	group.Use(middleware.JwtAuthMiddleware)
	{
		group.Get("/user/info/{id:int}", controller.GetUserInfoById)
		group.Get("/user/info/{username:string}", controller.GetUserInfoByName)
	}
}
