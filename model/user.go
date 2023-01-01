package model

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type PhoneRegister struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type EmailRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
