package router

import (
	"ChallengeCup/controller"
	"ChallengeCup/middleware"

	"github.com/kataras/iris/v12"
)

func InitRoute(r *iris.Application) {
	r.UseGlobal(middleware.Cors())
	v1 := r.Party("/api/v1")
	register_party := r.Party(v1.GetRelPath())
	{
		register_party.Post("/user/register", controller.PostUserRegisterByUserName)
		register_party.Post("/user/register", controller.PostUserRegisterByPhone)
		register_party.Post("/user/register", controller.PostUserRegisterByEmail)
	}
	login_party := r.Party(v1.GetRelPath())
	{
		login_party.Post("/user/login", controller.PostUserLogin)
		login_party.Post("/user/login", controller.PostUserLoginByPhone)
		login_party.Post("/user/login", controller.PostUserLoginByEmail)
	}
	activate_party := r.Party(v1.GetRelPath())
	{
		activate_party.Get("/user/code", controller.GetEmailCode)
		activate_party.Get("/user/code", controller.GetPhoneCode)
		activate_party.Post("/user/activate", controller.PostActivateEmail)
		activate_party.Post("/user/activate", controller.PostActivatePhone)
	}

	v1.Post("/user/user/refresh", controller.PostRefreshToken)

	user_party := r.Party(v1.GetRelPath())
	user_party.Use(middleware.JwtAuthMiddleware)
	{
		user_party.Post("/user/logout", controller.PostUserLogout)

		user_party.Get("/user/info", controller.GetUserInfo)
		user_party.Put("/user/info", controller.PutUserInfo)
		user_party.Put("/user/info/avatar", controller.PutUserAvatar)
		user_party.Put("/user/info/username", controller.PutUserName)
		user_party.Put("/user/info/password", controller.PutUserPassword)
		user_party.Put("/user/info/phone", controller.PutUserPhone)
		user_party.Put("/user/info/email", controller.PutUserEmail)
	}
}
