package model

type User struct {
	UserName string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
