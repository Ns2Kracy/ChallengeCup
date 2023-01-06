package model

type ValidateToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}
