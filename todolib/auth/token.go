package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
)

func CreateToken(userID, sessionID string, privateKey *rsa.PrivateKey, duration time.Duration) (string, error) {
	claims := newClaims(userID, sessionID, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := jwtToken.SignedString(privateKey)

	return token, err
}

func VerifyToken(tokenString string, publicKey *rsa.PublicKey) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
