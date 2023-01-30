package model

type (
	Username map[string]string
	Password map[string]string
	Email    map[string]string
	Phone    map[string]string
)

type UsernameRegister struct {
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

type PhoneCodeParams struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type EmailCodeParams struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type PhonePasswordParams struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type EmailPasswordParams struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

