package middleware

import (
	"ChallengeCup/common"
	"ChallengeCup/model"

	"github.com/kataras/iris/v12"
)

// 设置跨域
func Cors() iris.Handler {
	return func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		ctx.Header("Access-Control-Max-Age", "172800")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Content-Type", "application/json;charset=UTF-8")
		if ctx.Method() == "OPTIONS" {
			ctx.JSON(model.Result{
				Code:    common.SUCCESS,
				Message: common.Message(common.SUCCESS),
			})
		}
		ctx.Next()
	}
}
