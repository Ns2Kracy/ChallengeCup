package controller

import (
	"strconv"
	"time"

	"ChallengeCup/common"
	"ChallengeCup/dao"
	"ChallengeCup/middleware"
	"ChallengeCup/model"
	"ChallengeCup/service"
	"ChallengeCup/service/dbmodel"
	"ChallengeCup/utils/encrypt"

	"github.com/kataras/iris/v12"
)

func PostUserRegisterByUserNameAndPassword(ctx iris.Context) {
	userRequest := &model.User{}
	if err := ctx.ReadJSON(userRequest); err != nil {
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

	checkUserExist := service.Service.UserService.IsExistByName(userRequest.UserName)

	if checkUserExist {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_EXIST),
		})
		return
	}

	newUser := dbmodel.UserDBModel{}
	newUser.UserName = userRequest.UserName
	newUser.Password = encrypt.EncryptPassword(userRequest.Password)

	newUser = service.Service.UserService.NewUser(newUser)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func GetPhoneCode(ctx iris.Context) {
	// TODO: get phone code
}

func GetEmailCode(ctx iris.Context) {
	// TODO: get email code
}

func PostUserRegisterByPhone(ctx iris.Context) {
	// TODO: phone register
}

func PostUserRegisterByEmail(ctx iris.Context) {
	// TODO: email register
}

func PostActivateEmail(ctx iris.Context) {
	// TODO: activate email
}

func PostActivatePhone(ctx iris.Context) {
	// TODO: activate phone
}

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

	user := service.Service.UserService.GetUserByName(userRequest.UserName)

	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}

	if !encrypt.ComparePassword(userRequest.Password, user.Password) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.ERROR_PASSWORD),
		})
		return
	}

	token := model.ValidateToken{}
	expireTime := 3 * time.Hour * time.Duration(1)
	token.AccessToken = middleware.GetAccessToken(user.UserName, user.Password, user.ID)
	token.RefreshToken = middleware.GetRefreshToken(user.UserName, user.Password, user.ID)
	token.ExpiresIn = time.Now().Add(expireTime).Unix()
	dao.RedisClient.Set("access_token_"+strconv.Itoa(user.ID), token.AccessToken, expireTime)
	dao.RedisClient.Set("refresh_token_"+strconv.Itoa(user.ID), token.RefreshToken, expireTime*24*7)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    token,
	})
}

func PostUserLogout(ctx iris.Context) {
	id := ctx.GetHeader("id")
	dao.RedisClient.Del("access_token_" + id)
	dao.RedisClient.Del("refresh_token_" + id)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func GetUserInfoById(ctx iris.Context) {
	id := ctx.GetHeader("id")
	user := service.Service.UserService.GetUserById(id)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    user,
	})
}

func GetUserInfoByName(ctx iris.Context) {
	username := ctx.Params().Get("username")
	if len(username) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByName(username)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    user,
	})
}

func PutUserAvatar(ctx iris.Context) {
	avatar := ctx.FormValue("avatar")
	id := ctx.GetHeader("id")
	if len(avatar) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserById(id)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}
	user.Avatar = avatar
	service.Service.UserService.UpdateUser(user)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PutUserName(ctx iris.Context) {
	username := ctx.FormValue("username")
	id := ctx.GetHeader("id")
	if len(username) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserById(id)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}
	user.UserName = username
	service.Service.UserService.UpdateUser(user)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PutUserPassword(ctx iris.Context) {
	password := ctx.FormValue("password")
	id := ctx.GetHeader("id")
	if len(password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserById(id)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}
	user.Password = encrypt.EncryptPassword(password)
	service.Service.UserService.UpdateUser(user)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PutUserEmail(ctx iris.Context) {
	email := ctx.FormValue("email")
	id := ctx.GetHeader("id")
	if len(email) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserById(id)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}
	user.Email = email
	service.Service.UserService.UpdateUser(user)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PutUserPhone(ctx iris.Context) {
	phone := ctx.FormValue("phone")
	id := ctx.GetHeader("id")
	if len(phone) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserById(id)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}
	user.Phone = phone
	service.Service.UserService.UpdateUser(user)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}