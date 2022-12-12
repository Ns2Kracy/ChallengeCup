package router

import (
	"ChallengeCup/middleware"

	"github.com/kataras/iris/v12"
)

func InitRoute(r *iris.Application) {
	r.UseGlobal(middleware.Cors())
	r.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to Challenge Cup</h1>")
	})
	group := r.Party("/api")
	group.Use(middleware.JWT())
	{
		group.Get("/user", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"code": 200, "msg": "OK"})
		})
	}
}
