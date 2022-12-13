package middleware

import (
	"strconv"
	"time"

	"ChallengeCup/common"
	"ChallengeCup/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
)

var secret = make([]byte, 32)


type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}

func ParseToken(tokenString string, isVerify bool) (*Claims, error) {
	return nil, nil
}

func ValidateToken(token string) (*Claims, error) {
	if len(token) == 0 {
		return nil, nil
	}

	claims, err := ParseToken(token, true)

	if err != nil {
		return nil, err
	} else if !claims.VerifyExpiresAt(time.Now(), true) {
		return nil, nil
	}
	return claims, nil
}

func JWT() iris.Handler {
	return func(ctx iris.Context) {
		token := ctx.GetHeader("Authorization")
		if len(token) == 0 {
			ctx.JSON(model.Result{
				Code:    common.AUTH_ERROR,
				Message: common.Message(common.AUTH_ERROR),
			})
			ctx.StopExecution()
			return
		}
		claims, err := ValidateToken(token)
		if err != nil {
			ctx.JSON(model.Result{
				Code:    common.AUTH_ERROR,
				Message: common.Message(common.AUTH_ERROR),
			})
			ctx.StopExecution()
			return
		}
		ctx.Request().Header.Add("id", strconv.Itoa(claims.ID))
		ctx.Next()
	}
}
