package router

import (
	"ChallengeCup/controller"
	"ChallengeCup/middleware"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	_ "ChallengeCup/docs"
)

func InitRoute(r *iris.Application) {
	r.UseGlobal(middleware.Cors())
	r.Use(logger.New())
	r.Use(recover.New())
	r.Use(iris.Compression)

	swaggerConfig := &swagger.Config{
		URL:         "http://localhost:8080/swagger/swagger.json",
		DeepLinking: true,
	}
	swaggerUI := swagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler)
	r.Get("/swagger", swaggerUI)
	r.Get("/swagger/{any:path}", swaggerUI)

	v1 := r.Party("/api/v1")
	register_party := r.Party(v1.GetRelPath())
	{
		register_party.Post("/register", controller.PostUserRegisterByUserName)
		register_party.Post("/register/phone", controller.PostUserRegisterByPhone)
		register_party.Post("/register/email", controller.PostUserRegisterByEmail)
	}
	login_party := r.Party(v1.GetRelPath())
	{
		login_party.Post("/login", controller.PostUserLogin)
		login_party.Post("/login/phone", controller.PostUserLoginByPhone)
		login_party.Post("/login/email", controller.PostUserLoginByEmail)
	}
	activate_party := r.Party(v1.GetRelPath())
	{
		activate_party.Get("/code/email", controller.GetEmailCode)
		activate_party.Get("/code/phone", controller.GetPhoneCode)
		activate_party.Post("/activate/email", controller.PostActivateEmail)
		activate_party.Post("/activate/phone", controller.PostActivatePhone)
	}

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
