package service

import (
	"context"

	"todoauth/internal/model"
	"todoauth/internal/repository"
	"todoauth/internal/token"
	"todoauth/internal/util"
)

type AuthService interface {
	RenewAccessToken(ctx context.Context, refreshToken string) (string, error)
	RenewRefreshToken(ctx context.Context, refreshToken string) (*model.TokenPair, error)
	Login(ctx context.Context, arg *LoginParams) (*model.TokenPair, error)
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

func (s *authService) RenewRefreshToken(ctx context.Context, refreshToken string) (*model.TokenPair, error) {
	claims, err := s.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return nil, err
	}

	session, err := s.sessionRepository.CreateSession(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	err = s.sessionRepository.DeleteSession(ctx, claims.ID)
	if err != nil {
		return nil, err
	}

	aceessToken, err := s.tokenMaker.CreateAccessToken(claims.UserID, session.ID)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.tokenMaker.CreateRefreshToken(claims.UserID, session.ID)
	if err != nil {
		return nil, err
	}

	tokenPair := &model.TokenPair{
		AccessToken:  aceessToken,
		RefreshToken: newRefreshToken,
	}

	return tokenPair, nil
}

type LoginParams struct {
	Email    string
	Password string
}

func (s *authService) Login(ctx context.Context, arg *LoginParams) (*model.TokenPair, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, arg.Email)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(arg.Password, user.Password)
	if err != nil {
		return nil, err
	}

	session, err := s.sessionRepository.CreateSession(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.tokenMaker.CreateAccessToken(user.ID, session.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenMaker.CreateRefreshToken(user.ID, session.ID)
	if err != nil {
		return nil, err
	}

	tokenPair := &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokenPair, nil
}

func (s *authService) Logout(ctx context.Context, sessionID string) error {
	return s.sessionRepository.DeleteSession(ctx, sessionID)
}
