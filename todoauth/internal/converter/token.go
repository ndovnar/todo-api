package converter

import (
	"todoauth/internal/dto"
	"todoauth/internal/model"
)

func TokenPairModelToDTOResponse(tokenPair *model.TokenPair) *dto.TokenPairResponse {
	return &dto.TokenPairResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}
}

func AccessTokenToDTOResponse(accessToken string) *dto.AccessTokenResponse {
	return &dto.AccessTokenResponse{
		AccessToken: accessToken,
	}
}
