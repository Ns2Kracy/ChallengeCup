package controller

import (
	"time"

	"ChallengeCup/common"
	"ChallengeCup/dao"
	"ChallengeCup/middleware"
	"ChallengeCup/model"
	"ChallengeCup/service"
	"ChallengeCup/service/dbmodel"
	"ChallengeCup/utils/code"
	"ChallengeCup/utils/encrypt"
	uid "ChallengeCup/utils/uuid"
	"ChallengeCup/utils/verify"

	"github.com/kataras/iris/v12"
)

// @Summary 用户使用用户名注册
// @Description 用户使用用户名注册
// @Accept  json
// @Produce  json
// @Param   username	 query    string     true        "用户名"
// @Success 200 {object} model.Result "Success"
// @Router /api/v1/user/register [post]
func PostUserRegisterByUserName(ctx iris.Context) {
	userRequest := &model.UserNameRegister{}
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
	newUser.UUID = uid.GenerateUUID()

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
	phone := ctx.FormValue("phone")
	if !verify.VerifyPhone(phone) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	ValidateCode := code.RandomCode()
	go func() {
		dao.RedisClient.Set(ctx, "phone_code_"+phone, ValidateCode, time.Minute*5)
		err := code.PhoneSendCode(phone, ValidateCode)
		if err != nil {
			ctx.JSON(model.Result{
				Code:    common.CLIENT_ERROR,
				Message: common.Message(common.CLIENT_ERROR),
			})
		}
	}()
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    ValidateCode,
	})
}

func GetEmailCode(ctx iris.Context) {
	// TODO: get email code
}

func PostUserRegisterByPhone(ctx iris.Context) {
	userRequest := &model.PhoneRegister{}
	if userRequest.Code == dao.RedisClient.Get(ctx, "phone_code_"+userRequest.Phone).String() {
		ctx.JSON(model.Result{
			Code:    common.SUCCESS,
			Message: common.Message(common.SUCCESS),
		})
	} else {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.CLIENT_ERROR),
		})
	}
	checkUserExist := service.Service.UserService.IsExistByName(userRequest.Phone)

	if checkUserExist {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_EXIST),
		})
		return
	}

	newUser := dbmodel.UserDBModel{}
	newUser.Phone = userRequest.Phone
	newUser.UUID = uid.GenerateUUID()
	newUser.UserName = userRequest.Phone

	newUser = service.Service.UserService.NewUser(newUser)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PostActivateEmail(ctx iris.Context) {
	// TODO: activate email
}

func PostActivatePhone(ctx iris.Context) {
	// phone := ctx.URLParam("phone")
	// ValidateCode := ctx.URLParam("phone_code")
	// err := code.SendSmsCode(phone, ValidateCode)
	// if err != nil {
	// 	ctx.JSON(model.Result{
	// 		Code:    common.CLIENT_ERROR,
	// 		Message: common.Message(common.CLIENT_ERROR),
	// 	})
	// }
	// if ValidateCode != dao.RedisClient.Get(ctx, "phone_code").String() {
	// 	ctx.JSON(model.Result{
	// 		Code:    common.CLIENT_ERROR,
	// 		Message: common.Message(common.CLIENT_ERROR),
	// 	})
	// }

	// ctx.JSON(model.Result{
	// 	Code:    common.SUCCESS,
	// 	Message: common.Message(common.SUCCESS),
	// })
}

func PostUserLogin(ctx iris.Context) {
	userRequest := model.UserNameRegister{}
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
	token.ExpiresIn = time.Now().Add(expireTime).Unix()
	dao.RedisClient.Set(ctx, "AccessToken_"+user.UUID, token.AccessToken+user.UUID, expireTime)
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
	uid := ctx.GetHeader("uuid")
	dao.RedisClient.Del(ctx, "AccessToken_"+uid)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func GetUserInfo(ctx iris.Context) {
	uid := ctx.Params().Get("uuid")
	user := service.Service.UserService.GetUserByUID(uid)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    user,
	})
}

func PutUserInfo(ctx iris.Context) {
	uid := ctx.GetHeader("uuid")
	updater := dbmodel.UserDBModel{}
	if err := ctx.ReadJSON(&updater); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
	if user.ID == 0 {
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
	uid := ctx.GetHeader("uuid")
	if len(avatar) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
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
	uid := ctx.GetHeader("uuid")
	if len(username) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
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
	uid := ctx.GetHeader("uuid")
	if len(password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
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
	uid := ctx.GetHeader("uuid")
	if len(email) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
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
	uid := ctx.GetHeader("uuid")
	if len(phone) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}
	user := service.Service.UserService.GetUserByUID(uid)
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
