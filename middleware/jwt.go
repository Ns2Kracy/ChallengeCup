package middleware

import (
	"time"

	"ChallengeCup/model"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
)

var secret = make([]byte, 32)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, password string, id int, t time.Duration) (string, error) {
	claims := Claims{
		username,
		password,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

func ParseToken(token string, valid bool) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok {
			if valid && tokenClaims.Valid {
				return claims, nil
			} else if !valid {
				return claims, nil
			}
		}
	}
	return nil, err
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
				Code:    401,
				Message: "权限不足",
				Data:    nil,
			})
			ctx.StopExecution()
			return
		}
		claims, err := ValidateToken(token)
		if err != nil {
			ctx.JSON(model.Result{
				Code:    401,
				Message: "权限不足",
				Data:    nil,
			})
			ctx.StopExecution()
			return
		}
		ctx.Values().Set("claims", claims)
		ctx.Next()
	}
}
