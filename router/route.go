package router

import (
	"ChallengeCup/controller"
	"ChallengeCup/middleware"

	"github.com/kataras/iris/v12"
)

func InitRoute(r *iris.Application) {
	r.UseGlobal(middleware.Cors())
	v1 := r.Party("/api/v1")
	v1.Post("/user/register", controller.PostUserRegisterByUserName)
	v1.Post("/user/register", controller.PostUserRegisterByPhone)
	v1.Post("/user/register", controller.PostUserRegisterByEmail)
	v1.Post("/user/login", controller.PostUserLogin)
	v1.Post("/user/user/refresh", controller.PostRefreshToken)
	v1.Post("/user/code", controller.GetEmailCode)
	v1.Post("/user/code", controller.GetPhoneCode)
	v1.Post("/user/activation", controller.PostActivateEmail)
	v1.Post("/user/activation", controller.PostActivatePhone)

	user_group := r.Party(v1.GetRelPath())
	user_group.Use(middleware.JwtAuthMiddleware)
	{
		user_group.Post("/user/logout", controller.PostUserLogout)

		user_group.Get("/user/info", controller.GetUserInfo)
		user_group.Put("/user/info", controller.PutUserInfo)
		user_group.Put("/user/info/avatar", controller.PutUserAvatar)
		user_group.Put("/user/info/username", controller.PutUserName)
		user_group.Put("/user/info/password", controller.PutUserPassword)
		user_group.Put("/user/info/phone", controller.PutUserPhone)
		user_group.Put("/user/info/email", controller.PutUserEmail)
	}
}
