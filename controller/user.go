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
	"ChallengeCup/utils/uid"

	"github.com/kataras/iris/v12"
)

func PostUserRegisterByUserName(ctx iris.Context) {
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
	newUser.UID = uid.GenerateUID()

	newUser = service.Service.UserService.NewUser(newUser)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PostUserRegisterByEmail(ctx iris.Context) {
	// TODO: email register
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

func PostActivateEmail(ctx iris.Context) {
	// TODO: activate email
}

func PostActivatePhone(ctx iris.Context) {
	// TODO: activate phone
	ctx.Next()
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

	if user.UID == 0 {
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
	token.AccessToken = middleware.GetAccessToken(user.UserName, user.Password, user.UID)
	token.RefreshToken = middleware.GetRefreshToken(user.UserName, user.Password, user.UID)
	token.ExpiresIn = time.Now().Add(expireTime).Unix()
	dao.RedisClient.Set(ctx, "AccessToken_"+strconv.Itoa(user.UID), token.AccessToken+strconv.Itoa(user.UID), expireTime)
	dao.RedisClient.Set(ctx, "RefreshToken_"+strconv.Itoa(user.UID), token.RefreshToken+strconv.Itoa(user.UID), expireTime)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    token,
	})
}

func PostUserLoginByPhone(ctx iris.Context) {
	// TODO: phone login
}

func PostUserLoginByEmail(ctx iris.Context) {
	// TODO: email login
}

func PostUserLogout(ctx iris.Context) {
	uid := ctx.GetHeader("uuuuid")
	dao.RedisClient.Del(ctx, "AccessToken_"+uid)
	dao.RedisClient.Del(ctx, "RefreshToken_"+uid)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PostRefreshToken(ctx iris.Context) {
	// TODO: refresh token
}

func GetUserInfo(ctx iris.Context) {
	uid := ctx.Params().Get("uuuuid")
	user := service.Service.UserService.GetUserByUID(uid)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    user,
	})
}

func PutUserInfo(ctx iris.Context) {
	uid := ctx.GetHeader("uuuuid")
	updater := dbmodel.UserDBModel{}
	if err := ctx.ReadJSON(&updater); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
	if user.UID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}
	service.Service.UserService.UpdateUser(updater)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PutUserAvatar(ctx iris.Context) {
	avatar := ctx.FormValue("avatar")
	uid := ctx.GetHeader("uid")
	if len(avatar) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
	if user.UID == 0 {
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
	uid := ctx.GetHeader("uid")
	if len(username) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
	if user.UID == 0 {
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
	uid := ctx.GetHeader("uuuuid")
	if len(password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
	if user.UID == 0 {
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
	uid := ctx.GetHeader("uuuuid")
	if len(email) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
	if user.UID == 0 {
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
	uid := ctx.GetHeader("uuuuid")
	if len(phone) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
	if user.UID == 0 {
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
