package model

type PhoneRegister struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type PhoneCodeParams struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type PhonePasswordParams struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
