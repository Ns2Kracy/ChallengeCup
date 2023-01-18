package router

import (
	"ChallengeCup/controller"
	"ChallengeCup/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func InitRoute(r *iris.Application) {
	r.UseGlobal(middleware.Cors())
	r.Use(logger.New())
	r.Use(recover.New())
	r.Use(iris.Compression)
	user := r.Party("/user")
	{
		user.Post("/register", controller.PostUserRegisterByUserName)
		user.Post("/register/phone", controller.PostUserRegisterByPhone)
		user.Post("/register/email", controller.PostUserRegisterByEmail)
		user.Post("/login", controller.PostUserLogin)
		user.Post("/login/phone", controller.PostUserLoginByPhone)
		user.Post("/login/email", controller.PostUserLoginByEmail)
		user.Get("/code/email", controller.GetEmailCode)
		user.Get("/code/phone", controller.GetPhoneCode)
		user.Post("/activate/email", controller.PostActivateEmail)
		user.Post("/activate/phone", controller.PostActivatePhone)

		user.Party("/", middleware.JwtAuthMiddleware)
		{
			user.Post("/logout", controller.PostUserLogout)

			user.Get("/info", controller.GetUserInfo)
			user.Put("/info/update", controller.PutUserInfo)
			user.Put("/info/avatar", controller.PutUserAvatar)
			user.Put("/info/username", controller.PutUserName)
			user.Put("/info/password", controller.PutUserPassword)
			user.Put("/info/phone", controller.PutUserPhone)
			user.Put("/info/email", controller.PutUserEmail)
		}
	}
}
