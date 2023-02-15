package common

const (
	SUCCESS      = 200
	SERVER_ERROR = 500
	CLIENT_ERROR = 400
	AUTH_ERROR   = 401

	INVALID_PARAMS        = 10001
	SIMPLE_PASSWORD       = 10002
	USER_EXIST            = 10003
	ERROR_PASSWORD        = 10004
	USER_NOT_EXIST        = 10005
	INVALID_REFRESH_TOKEN = 10006
	EMAIL_EXIST           = 10007
	PHONE_EXIST           = 10008
	CODE_VALIDATION_ERROR = 10009
	INVALID_PHONE         = 10010
)

var message = map[int]string{
	SUCCESS:      "OK",
	SERVER_ERROR: "Server Error",
	CLIENT_ERROR: "Client Error",
	AUTH_ERROR:   "Auth Error",

	INVALID_PARAMS:        "Invalid Params",
	SIMPLE_PASSWORD:       "Password is too simple",
	USER_EXIST:            "User already exist",
	ERROR_PASSWORD:        "Password is wrong",
	USER_NOT_EXIST:        "User not exist",
	INVALID_REFRESH_TOKEN: "Invalid refresh token",
	EMAIL_EXIST:           "Email already exist",
	PHONE_EXIST:           "Phone already exist",
	CODE_VALIDATION_ERROR: "Code validation error",
	INVALID_PHONE:         "Invalid phone number",
}

func Message(code int) string {
	str, ok := message[code]
	if ok {
		return str
	}
	return message[SERVER_ERROR]
}
