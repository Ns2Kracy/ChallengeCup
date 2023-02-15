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

		user.Post("/login_and_register", controller.PostUserLoginAndRegister)
		user.Post("/login/code", controller.PostUserLoginByPhoneCode)
		user.Post("/login", controller.PostUserLoginByPhonePassword)
		user.Get("/get_code", controller.GetPhoneCode)

		user.Post("/logout", controller.PostUserLogout)

		user.Party("/")
		{
			user.Use(middleware.JwtAuthMiddleware)
			user.Get("/info", controller.GetUserInfo)
			user.Put("/info/update", controller.PutUserInfo)
			user.Put("/info/avatar", controller.PutUserAvatar)
			user.Put("/info/username", controller.PutUsername)
			user.Put("/info/password", controller.PutUserPassword)
			user.Put("/info/phone", controller.PutUserPhone)
		}

		user.Party("/status")
		{
			user.Use(middleware.JwtAuthMiddleware)
			user.Get("/", controller.GetData)
			user.Get("/temperature", controller.GetTemperature)
			user.Get("/heart_rate", controller.GetHeartRate)
			user.Get("/blood_oxygen", controller.GetBloodOxygen)
			user.Get("/get_by_time", controller.GetDataByTime)
			user.Get("/get_by_time/temperature", controller.GetTemperatureByTime)
			user.Get("/get_by_time/heart_rate", controller.GetHeartRateByTime)
			user.Get("/get_by_time/blood_oxygen", controller.GetBloodOxygenByTime)
		}
	}
}
