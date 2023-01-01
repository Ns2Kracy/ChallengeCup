package controller

import (
	"time"

	"ChallengeCup/common"
	"ChallengeCup/middleware"
	"ChallengeCup/model"
	"ChallengeCup/service"
	"ChallengeCup/service/dbmodel"
	"ChallengeCup/utils/encrypt"

	"github.com/kataras/iris/v12"
)

// Description: register by username
// Author: Ns2Kracy
// Updated: 12 13th, 2022 23:05
// Accept: application/json
// Produce: application/json
// Params: username, password (string)
// Router: /user/register-by-name [POST]
func PostUserRegisterByUserNameAndPassword(ctx iris.Context) {
	userRequest := model.User{}
	if err := ctx.ReadJSON(&userRequest); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	if len(userRequest.UserName) == 0 || len(userRequest.Password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	} else if len(userRequest.UserName) < 6 || len(userRequest.Password) < 6 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.SIMPLE_PASSWORD),
		})
		return
	}

	checkUserExist := service.AppService.UserService.GetUserByName(userRequest.UserName)

	if checkUserExist.ID != 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_EXIST),
		})
		return
	}

	newUser := dbmodel.UserDBModel{}
	newUser.UserName = userRequest.UserName
	newUser.Password = encrypt.EncryptPassword(userRequest.Password)

	// 将用户信息存入数据库
	newUser = service.AppService.UserService.NewUser(newUser)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

// Description: register by phone
// Author: Ns2Kracy
// Updated: 12 23rd, 2022 23:56
// Accept: application/json
// Produce: application/json
// Params: username, password (string)
// Router: /user/register-by-phone [POST]
func PostUserRegisterByPhone(ctx iris.Context) {
	// TODO: phone register
}

// Description: register by email
// Author: Ns2Kracy
// Updated: 12 23rd, 2022 23:56
// Accept: application/json
// Produce: application/json
// Params: username, password (string)
// Router: /user/register-by-email [POST]
func PostUserRegisterByEmail(ctx iris.Context) {
	// TODO: email register
}

// Description: login
// Author: Ns2Kracy
// Updated: 12 13th, 2022 23:05
// Accept: application/json
// Produce: application/json
// Params: username | phone | email, password (string)
// Router: /user/login [POST]
func PostUserLogin(ctx iris.Context) {
	userRequest := model.User{}
	if err := ctx.ReadJSON(&userRequest); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	if len(userRequest.UserName) == 0 || len(userRequest.Password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	user := service.AppService.UserService.GetUserByName(userRequest.UserName)

	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.ERROR_PASSWORD),
		})
		return
	}

	if !encrypt.ComparePassword(userRequest.Password, user.Password) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	token := model.VaildateToken{}
	token.AccessToken = middleware.GetAccessToken(user.UserName, user.Password, user.ID)
	token.RefreshToken = middleware.GetRefreshToken(user.UserName, user.Password, user.ID)
	token.ExpiresIn = time.Now().Add(3 * time.Hour * time.Duration(1)).Unix()

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    token,
	})
}
