package controller

import (
	"ChallengeCup/common"
	"ChallengeCup/model"
	"ChallengeCup/service"
	"ChallengeCup/service/dbmodel"
	"ChallengeCup/utils/encrypt"

	"github.com/kataras/iris/v12"
)

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

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PostUserRegisterByPhone(ctx iris.Context) {}

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

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}
