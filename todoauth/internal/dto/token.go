package dto

type RenewTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type TokenPairResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
