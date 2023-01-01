package model

type ValidateToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}
