package repository

import (
	"context"

	"todoauth/internal/modeldb"
)

type CreateUserParams struct {
	Email    string
	Password string
}

type UserRepository interface {
	CreateUser(ctx context.Context, arg *CreateUserParams) (*modeldb.User, error)
	GetUserByID(ctx context.Context, id string) (*modeldb.User, error)
	GetUserByEmail(ctx context.Context, email string) (*modeldb.User, error)
}
