package service

import (
	"context"

	"todoauth/internal/model"
	"todoauth/internal/repository"
	"todoauth/internal/util"
)

type UserService interface {
	CreateUser(ctx context.Context, email, password string) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(ctx context.Context, email, password string) (*model.User, error) {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return s.userRepository.CreateUser(ctx, &model.User{
		Email:    email,
		Password: hashedPassword,
	})
}
