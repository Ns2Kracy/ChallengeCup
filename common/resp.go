package common

const (
	SUCCESS      = 200
	SERVER_ERROR = 500
	CLIENT_ERROR = 400
	AUTH_ERROR   = 401

	INVALID_PARAMS  = 10001
	SIMPLE_PASSWORD = 10002
	USER_EXIST      = 10003
	ERROR_PASSWORD  = 10004
	USER_NOT_EXIST  = 10005
)

var message = map[int]string{
	SUCCESS:      "OK",
	SERVER_ERROR: "Server Error",
	CLIENT_ERROR: "Client Error",
	AUTH_ERROR:   "Auth Error",

	INVALID_PARAMS:  "Invalid Params",
	SIMPLE_PASSWORD: "Password is too simple",
	USER_EXIST:      "User already exist",
	ERROR_PASSWORD:  "Password is wrong",
	USER_NOT_EXIST:  "User not exist",
}

func Message(code int) string {
	str, ok := message[code]
	if ok {
		return str
	}
	return message[SERVER_ERROR]
}
