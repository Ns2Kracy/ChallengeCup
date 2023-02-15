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
	log "ChallengeCup/utils/logger"
	"ChallengeCup/utils/uuid"
	"ChallengeCup/utils/verify"

	"github.com/kataras/iris/v12"
)

// PostUserLoginByUsername 注册并登录
func PostUserLoginAndRegister(ctx iris.Context) {
	userRequest := &model.PhoneRegister{}

	if err := ctx.ReadJSON(userRequest); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	if verify.VerifyPhone(userRequest.Phone) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PHONE),
		})
		return
	}

	if userRequest.Code != dao.RedisClient.Get(ctx, "PhoneCode_"+userRequest.Phone).Val() {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.CODE_VALIDATION_ERROR),
		})
		return
	}

	if len(userRequest.Password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	} else if len(userRequest.Password) < 8 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.SIMPLE_PASSWORD),
		})
		return
	}

	checkUserExist := service.Service.UserService.CheckPhone(userRequest.Phone)

	if !checkUserExist {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_EXIST),
		})
		return
	}

	newUser := dbmodel.UserDBModel{
		UUID:           uuid.GenerateUUID(),
		UserName:       userRequest.Phone,
		Password:       encrypt.EncryptPassword(userRequest.Password),
		Phone:          userRequest.Phone,
		IsPhoneActived: true,
	}

	newUser = service.Service.UserService.NewUser(newUser)
	dao.RedisClient.Del(ctx, "PhoneCode_"+userRequest.Phone)

	// 过期时间30天
	expireTime := 30 * 24 * time.Hour
	token := model.ValidateToken{
		AccessToken: middleware.GetAccessToken(newUser.UserName, newUser.Password, newUser.ID),
		ExpiresIn:   time.Now().Add(expireTime).Unix(),
	}
	dao.RedisClient.Set(ctx, "AccessToken_"+strconv.Itoa(int(newUser.UUID)), token.AccessToken, expireTime)

	data := map[string]interface{}{
		"token": token,
		"uuid":  newUser.UUID,
	}

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    data,
	})
}

// GetPhoneCode 获取手机验证码
func GetPhoneCode(ctx iris.Context) {
	phone := ctx.FormValue("phone")

	if verify.VerifyPhone(phone) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	ValidateCode := verify.RandomCode()
	dao.RedisClient.Set(ctx, "PhoneCode_"+phone, ValidateCode, time.Minute*5)

	go func() {
		err := verify.PhoneSendCode(phone, ValidateCode)
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

// PostUserLoginByPhonePassword 手机号密码登录
func PostUserLoginByPhonePassword(ctx iris.Context) {
	userRequest := model.PhonePasswordParams{}

	if err := ctx.ReadJSON(&userRequest); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	if len(userRequest.Phone) != 11 || len(userRequest.Password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	user := service.Service.UserService.GetUserByPhone(userRequest.Phone)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}

	log.Infof(user.Password)
	log.Infof(userRequest.Password)

	if !encrypt.ComparePassword(user.Password, userRequest.Password) {
		log.Infof("password is wrong")
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.ERROR_PASSWORD),
		})
		return
	}

	expireTime := 24 * 30 * time.Hour * time.Duration(1)
	token := model.ValidateToken{
		AccessToken: middleware.GetAccessToken(user.UserName, user.Password, user.ID),
		ExpiresIn:   time.Now().Add(expireTime).Unix(),
	}
	dao.RedisClient.Set(ctx, "AccessToken_"+strconv.Itoa(int(user.UUID)), token.AccessToken, expireTime)

	data := map[string]interface{}{
		"token": token,
		"uuid":  user.UUID,
	}

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    data,
	})
}

// PostUserLoginByPhoneCode 手机号验证码登录
func PostUserLoginByPhoneCode(ctx iris.Context) {
	userRequest := model.PhoneCodeParams{}

	if err := ctx.ReadJSON(&userRequest); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	if len(userRequest.Phone) == 0 || len(userRequest.Code) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	user := service.Service.UserService.GetUserByPhone(userRequest.Phone)

	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}

	if userRequest.Code != dao.RedisClient.Get(ctx, "PhoneCode_"+userRequest.Phone).Val() {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.CODE_VALIDATION_ERROR),
		})
		return
	}

	expireTime := 24 * 30 * time.Hour * time.Duration(1)
	token := model.ValidateToken{
		AccessToken: middleware.GetAccessToken(user.UserName, user.Password, user.ID),
		ExpiresIn:   time.Now().Add(expireTime).Unix(),
	}
	dao.RedisClient.Set(ctx, "AccessToken_"+strconv.Itoa(int(user.UUID)), token.AccessToken, expireTime)

	data := map[string]interface{}{
		"token": token,
		"uuid":  user.UUID,
	}

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    data,
	})
}

// PostUserLogout 用户登出
func PostUserLogout(ctx iris.Context) {
	uuid := ctx.GetHeader("uuid")

	dao.RedisClient.Del(ctx, "AccessToken_"+uuid)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

// GetUserInfo 获取用户信息
func GetUserInfo(ctx iris.Context) {
	uuid := ctx.GetHeader("uuid")

	user := service.Service.UserService.GetUserByUUID(uuid)

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

// PutUserInfo 更新用户信息
func PutUserInfo(ctx iris.Context) {
	uuid := ctx.GetHeader("uuid")

	updater := dbmodel.UserDBModel{}
	if err := ctx.ReadJSON(&updater); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	user := service.Service.UserService.GetUserByUUID(uuid)
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
	// TODO
}

// PutUserAvatar 更新用户头像
func PutUserAvatar(ctx iris.Context) {
	uuid := ctx.GetHeader("uuid")

	avatar := ctx.FormValue("avatar")
	if len(avatar) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	user := service.Service.UserService.GetUserByUUID(uuid)
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
	// TODO
}

// PutUserPassword 更新用户名
func PutUsername(ctx iris.Context) {
	uuid := ctx.GetHeader("uuid")

	username := map[string]string{}
	if err := ctx.ReadJSON(&username); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	if len(username) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	user := service.Service.UserService.GetUserByUUID(uuid)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}
	user.UserName = username["username"]
	service.Service.UserService.UpdateUser(user)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

// PutUserPassword 更新密码
func PutUserPassword(ctx iris.Context) {
	uuid := ctx.GetHeader("uuid")

	password := map[string]string{}
	if err := ctx.ReadJSON(&password); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	if len(password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	user := service.Service.UserService.GetUserByUUID(uuid)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}
	user.Password = encrypt.EncryptPassword(password["password"])
	service.Service.UserService.UpdateUser(user)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

// PutUserPhone 更新手机号
func PutUserPhone(ctx iris.Context) {
	uuid := ctx.GetHeader("uuid")

	phone := model.PhoneCodeParams{}

	if err := ctx.ReadJSON(&phone); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	if len(phone.Phone) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		return
	}

	user := service.Service.UserService.GetUserByUUID(uuid)
	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		return
	}

	if service.Service.UserService.CheckPhone(phone.Phone) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.PHONE_EXIST),
		})
		return
	}

	if phone.Code != dao.RedisClient.Get(ctx, "PhoneCode_"+phone.Phone).Val() {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.CODE_VALIDATION_ERROR),
		})
		return
	}

	user.Phone = phone.Phone
	user.IsPhoneActived = true
	user.PhoneActivedAt = time.Now().Unix()
	service.Service.UserService.UpdateUser(user)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}
