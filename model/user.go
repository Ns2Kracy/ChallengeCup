package model

type Username map[string]string
type Password map[string]string
type Email map[string]string
type Phone map[string]string

type UserNameRegister struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type PhoneRegister struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type EmailRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
