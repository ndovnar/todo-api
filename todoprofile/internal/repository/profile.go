package repository

import (
	"context"

	"todoprofile/internal/modeldb"
)

type GetProfileParams struct {
	UserID string
}

type CreateProfileParams struct {
	UserID    string
	FirstName string
	LastName  string
}

type UpdateProfileParams struct {
	UserID    string
	FirstName string
	LastName  string
}

type DeleteProfileParams struct {
	UserID string
}

type ProfileRepository interface {
	GetProfile(ctx context.Context, arg *GetProfileParams) (*modeldb.Profile, error)
	CreateProfile(ctx context.Context, arg *CreateProfileParams) (*modeldb.Profile, error)
	UpdateProfile(ctx context.Context, arg *UpdateProfileParams) (*modeldb.Profile, error)
	DeleteProfile(ctx context.Context, arg *DeleteProfileParams) error
}
