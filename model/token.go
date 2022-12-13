package model

type VaildateToken struct {
	AccessToken string
	RefreshToken string
	ExpiresIn int64
}