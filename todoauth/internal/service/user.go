package service

import (
	"context"

	"todoauth/internal/converter"
	"todoauth/internal/model"
	"todoauth/internal/repository"
	"todoauth/internal/util"
)

type UserService interface {
	CreateUser(ctx context.Context, args *CreateUserParams) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

type CreateUserParams struct {
	Email    string
	Password string
}

func (s *userService) CreateUser(ctx context.Context, args *CreateUserParams) (*model.User, error) {
	hashedPassword, err := util.HashPassword(args.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.CreateUser(ctx, &repository.CreateUserParams{
		Email:    args.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return converter.UserDBModelToModel(user), nil
}
