package middleware

import (
	"time"

	"ChallengeCup/common"
	"ChallengeCup/dao"
	"ChallengeCup/model"

	"github.com/golang-jwt/jwt/v4"
	iris_jwt "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}

var (
	JwtAuthMiddleware = iris_jwt.New(iris_jwt.Config{
		ValidationKeyGetter: validationKeyGetterFuc,
		SigningMethod:       jwt.SigningMethodHS256,
		Expiration:          true,
		Extractor:           extractor,
		ErrorHandler:        errorHandler,
	}).Serve
	validationKeyGetterFuc = func(token *jwt.Token) (interface{}, error) {
		return common.Secret, nil
	}
	extractor = func(ctx iris.Context) (string, error) {
		auth := ctx.GetHeader("Authorization")
		uuid := ctx.GetHeader("uuid")
		if auth == "" {
			return "", nil
		}
		if auth != dao.RedisClient.Get(ctx, "AccessToken_"+uuid).Val() {
			return "", nil
		}
		return auth, nil
	}
	errorHandler = func(ctx iris.Context, err error) {
		ctx.StopWithJSON(common.AUTH_ERROR, model.Result{
			Code:    common.AUTH_ERROR,
			Message: common.Message(common.AUTH_ERROR),
		})
	}
)

func GenerateToken(username string, password string, id int, issuer string, t time.Duration) (string, error) {
	claims := Claims{
		username,
		password,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t)),
			Issuer:    issuer,
		},
	}
	token := iris_jwt.NewTokenWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(common.Secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetAccessToken(username, pwd string, id int) string {
	token, err := GenerateToken(username, pwd, id, "challengecup", 3*time.Hour*time.Duration(1))
	if err != nil {
		return ""
	}
	return token
}
