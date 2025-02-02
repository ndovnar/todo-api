package token

import (
	"crypto/rsa"
	"time"

	"lib/auth"
)

type TokenMaker struct {
	privateKey           *rsa.PrivateKey
	publicKey            *rsa.PublicKey
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewTokenMaker(config *Config, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *TokenMaker {
	return &TokenMaker{
		privateKey:           privateKey,
		publicKey:            publicKey,
		accessTokenDuration:  config.AccessTokenDuration,
		refreshTokenDuration: config.RefreshTokenDuration,
	}
}

func (m *TokenMaker) CreateAccessToken(userID string, sessionID string) (string, error) {
	return auth.CreateToken(userID, sessionID, m.privateKey, m.accessTokenDuration)
}

func (m *TokenMaker) CreateRefreshToken(userID string, sessionID string) (string, error) {
	return auth.CreateToken(userID, sessionID, m.privateKey, m.refreshTokenDuration)
}

func (m *TokenMaker) VerifyToken(tokenString string) (*auth.Claims, error) {
	return auth.VerifyToken(tokenString, m.publicKey)
}
