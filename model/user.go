package model

type UserNameRegister struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type PhoneRegister struct {
	Phone    string `json:"phone"`
	Code     string `json:"code"`
}

type EmailRegister struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
}
