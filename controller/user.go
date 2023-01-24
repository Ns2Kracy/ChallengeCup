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
	uuid "ChallengeCup/utils/uuid"
	"ChallengeCup/utils/verify"

	"github.com/kataras/iris/v12"
)

func PostUserRegisterByUserName(ctx iris.Context) {
	userRequest := &model.UserNameRegister{}

	if err := ctx.ReadJSON(userRequest); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Errorf("User %s register failed, error: %s", userRequest.UserName, err)
		return
	}

	if len(userRequest.UserName) == 0 || len(userRequest.Password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Errorf("User %s register failed, the username or password is empty", userRequest.UserName)
		return
	} else if len(userRequest.UserName) < 6 || len(userRequest.Password) < 8 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.SIMPLE_PASSWORD),
		})
		log.Errorf("User %s register failed, the username or password is too simple", userRequest.UserName)
		return
	}

	checkUserExist := service.Service.UserService.IsExistByName(userRequest.UserName)

	if checkUserExist {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_EXIST),
		})
		log.Errorf("User %s register failed, the username has been registered", userRequest.UserName)
		return
	}

	newUser := dbmodel.UserDBModel{
		UserName: userRequest.UserName,
		UUID:     uuid.GenerateUUID(),
		Password: encrypt.EncryptPassword(userRequest.Password),
	}

	newUser = service.Service.UserService.NewUser(newUser)

	log.Infof("User %s register success", newUser.UserName)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PostUserRegisterByPhone(ctx iris.Context) {
	userRequest := &model.PhoneRegister{}

	if err := ctx.ReadJSON(userRequest); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Errorf("register failed, invalid params")
		return
	}

	if !verify.VerifyPhone(userRequest.Phone) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Errorf("register failed, invalid phone number")
		return
	}

	if userRequest.Code != dao.RedisClient.Get(ctx, "phone_code_"+userRequest.Phone).String() {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.CLIENT_ERROR),
		})
		log.Errorf("register failed, invalid phone code")
		return
	}

	if len(userRequest.Password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Infof("register failed, the password is empty")
		return
	} else if len(userRequest.Password) < 6 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.SIMPLE_PASSWORD),
		})
		log.Infof("register failed, the password is too simple")
		return
	}

	checkUserExist := service.Service.UserService.IsExistByName(userRequest.Phone)

	if checkUserExist {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_EXIST),
		})
		log.Errorf("Phone %s register failed, this phone has been registered", userRequest.Phone)
		return
	}

	newUser := dbmodel.UserDBModel{
		UUID:     uuid.GenerateUUID(),
		UserName: userRequest.Phone,
		Password: encrypt.EncryptPassword(userRequest.Password),
		Phone:    userRequest.Phone,
	}

	newUser = service.Service.UserService.NewUser(newUser)

	log.Infof("Phone %s register success", newUser.Phone)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PostUserRegisterByEmail(ctx iris.Context) {
	userRequest := &model.EmailRegister{}

	if err := ctx.ReadJSON(userRequest); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Infof("Email %s register failed, error: %s", userRequest.Email, err)
		return
	}
	if !verify.VerifyEmail(userRequest.Email) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Infof("Email %s register failed, invalid email address", userRequest.Email)
		return
	}

	if userRequest.Code != dao.RedisClient.Get(ctx, "email_code_"+userRequest.Email).String() {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.CLIENT_ERROR),
		})
		log.Infof("Email %s register failed, invalid email code", userRequest.Email)
		return
	}

	if len(userRequest.Password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Infof("Email %s register failed, the password is empty", userRequest.Email)
		return
	} else if len(userRequest.Password) < 6 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.SIMPLE_PASSWORD),
		})
		log.Infof("Email %s register failed, the password is too simple", userRequest.Email)
		return
	}

	checkUserExist := service.Service.UserService.IsExistByName(userRequest.Email)

	if checkUserExist {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_EXIST),
		})
		log.Infof("Email %s register failed, the email has been registered", userRequest.Email)
		return
	}

	newUser := dbmodel.UserDBModel{
		UUID:     uuid.GenerateUUID(),
		UserName: userRequest.Email,
		Password: encrypt.EncryptPassword(userRequest.Password),
		Email:    userRequest.Email,
	}

	newUser = service.Service.UserService.NewUser(newUser)

	log.Infof("Email %s register success", newUser.UserName)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func GetPhoneCode(ctx iris.Context) {
	phone := ctx.FormValue("phone")

	if !verify.VerifyPhone(phone) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Errorf("Phone %s get code failed, invalid phone number", phone)
		return
	}

	ValidateCode := verify.RandomCode()

	go func() {
		dao.RedisClient.Set(ctx, "phone_code_"+phone, ValidateCode, time.Minute*5)
		log.Infof("caching phone code %s to redis", ValidateCode)
		err := verify.PhoneSendCode(phone, ValidateCode)
		if err != nil {
			ctx.JSON(model.Result{
				Code:    common.CLIENT_ERROR,
				Message: common.Message(common.CLIENT_ERROR),
			})
			log.Errorf("sending verification code to phone %s failed, error: %s", phone, err)
		}
	}()

	log.Infof("sending verification code to phone %s success", phone)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    ValidateCode,
	})
}

func GetEmailCode(ctx iris.Context) {
	email := ctx.FormValue("email")

	if !verify.VerifyEmail(email) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Errorf("invalid email address")
		return
	}

	ValidateCode := verify.RandomCode()

	go func() {
		dao.RedisClient.Set(ctx, "email_code_"+email, ValidateCode, time.Minute*5)
		log.Infof("caching phone code %s to redis", ValidateCode)
		err := verify.MailSendCode(email, ValidateCode)
		if err != nil {
			ctx.JSON(model.Result{
				Code:    common.CLIENT_ERROR,
				Message: common.Message(common.CLIENT_ERROR),
			})
			log.Errorf("sending verification code to email %s failed, error: %s", email, err)
		}
	}()

	log.Infof("sending verification code to email %s success", email)

	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    ValidateCode,
	})
}

func PostActivatePhone(ctx iris.Context) {
	// TODO: activate phone
}

func PostActivateEmail(ctx iris.Context) {
	// TODO: activate email
}

func PostUserLogin(ctx iris.Context) {
	userRequest := model.UserNameRegister{}

	if err := ctx.ReadJSON(&userRequest); err != nil {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Errorf("login failed, invalid params")
		return
	}

	if len(userRequest.UserName) == 0 || len(userRequest.Password) == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.INVALID_PARAMS),
		})
		log.Errorf("login failed, username or password is empty")
		return
	}

	user := service.Service.UserService.GetUserByName(userRequest.UserName)

	if user.ID == 0 {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.USER_NOT_EXIST),
		})
		log.Errorf("user %s login failed, user not exist", userRequest.UserName)
		return
	}

	if !encrypt.ComparePassword(userRequest.Password, user.Password) {
		ctx.JSON(model.Result{
			Code:    common.CLIENT_ERROR,
			Message: common.Message(common.ERROR_PASSWORD),
		})
		log.Errorf("user %s login failed, password is wrong", userRequest.UserName)
		return
	}

	expireTime := 24 * time.Hour * time.Duration(1)
	token := model.ValidateToken{
		AccessToken: middleware.GetAccessToken(user.UserName, user.Password, user.ID),
		ExpiresIn:   time.Now().Add(expireTime).Unix(),
	}
	dao.RedisClient.Set(ctx, "AccessToken_"+strconv.Itoa(int(user.UUID)), token.AccessToken, expireTime)
	log.Infof("user %s login success, caching access token to redis", user.UserName)

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
	uuid := ctx.GetHeader("uuid")
	dao.RedisClient.Del(ctx, "AccessToken_"+uuid)
	log.Infof("user logout success, delete access token from redis")
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func GetUserInfo(ctx iris.Context) {
	uuid := ctx.GetHeader("uuid")
	user := service.Service.UserService.GetUserByUUID(uuid)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    user,
	})
}

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
}

func PutUserAvatar(ctx iris.Context) {
	avatar := ctx.FormValue("avatar")
	uuid := ctx.GetHeader("uuid")
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
}

func PutUserName(ctx iris.Context) {
	username := ctx.Params().Get("username")
	uuid := ctx.GetHeader("uuid")
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
	user.UserName = username
	service.Service.UserService.UpdateUser(user)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PutUserPassword(ctx iris.Context) {
	password := ctx.Params().Get("password")
	uuid := ctx.GetHeader("uuid")
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
	user.Password = encrypt.EncryptPassword(password)
	service.Service.UserService.UpdateUser(user)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PutUserEmail(ctx iris.Context) {
	email := ctx.Params().Get("email")
	uuid := ctx.GetHeader("uuid")
	if len(email) == 0 {
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
	user.Email = email
	service.Service.UserService.UpdateUser(user)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}

func PutUserPhone(ctx iris.Context) {
	phone := ctx.Params().Get("phone")
	uuid := ctx.GetHeader("uuid")
	if len(phone) == 0 {
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
	user.Phone = phone
	service.Service.UserService.UpdateUser(user)
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
	})
}
