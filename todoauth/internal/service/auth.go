package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"todoauth/internal/model"
	"todoauth/internal/repository"
	"todoauth/internal/token"
	"todoauth/internal/util"
)

type AuthService interface {
	RenewAccessToken(ctx context.Context, refreshToken string) (string, error)
	RenewRefreshToken(ctx context.Context, refreshToken string) (string, string, error)
	Login(ctx context.Context, email, password string) (string, string, error)
	Logout(ctx context.Context, sessionID string) error
}

type authService struct {
	tokenMaker        *token.TokenMaker
	sessionRepository repository.SessionRepository
	userRepository    repository.UserRepository
}

func NewAuthService(tokenMaker *token.TokenMaker, sessionRepository repository.SessionRepository, userRepository repository.UserRepository) AuthService {
	return &authService{
		tokenMaker:        tokenMaker,
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
	}
}

func (s *authService) RenewAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := s.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return "", err
	}

	_, err = s.sessionRepository.GetSessionByID(ctx, claims.ID)
	if err != nil {
		return "", err
	}

	aceesToken, err := s.tokenMaker.CreateAccessToken(claims.UserID, claims.ID)
	if err != nil {
		return "", err
	}

	return aceesToken, nil
}

func (s *authService) RenewRefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	claims, err := s.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	session, err := s.sessionRepository.CreateSession(ctx, &model.Session{
		UserID: claims.UserID,
	})
	if err != nil {
		return "", "", err
	}

	err = s.sessionRepository.DeleteSession(ctx, claims.ID)
	if err != nil {
		fmt.Println(claims.ID)
		fmt.Println(err)
		return "", "", err
	}

	aceesToken, err := s.tokenMaker.CreateAccessToken(claims.UserID, session.ID)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := s.tokenMaker.CreateRefreshToken(claims.UserID, session.ID)
	if err != nil {
		return "", "", err
	}

	return aceesToken, newRefreshToken, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", "", err
	}

	session, err := s.sessionRepository.CreateSession(ctx, &model.Session{
		UserID: user.ID,
	})
	if err != nil {
		return "", "", err
	}

	accessToken, err := s.tokenMaker.CreateAccessToken(user.ID, session.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed to create access token")
		return "", "", err
	}

	refreshToken, err := s.tokenMaker.CreateRefreshToken(user.ID, session.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) Logout(ctx context.Context, sessionID string) error {
	err := s.sessionRepository.DeleteSession(ctx, sessionID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
