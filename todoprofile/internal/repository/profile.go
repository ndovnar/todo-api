package repository

import (
	"context"

	"todoprofile/internal/model"
)

type GetProfileParams struct {
	UserID string
}

type UpdateProfileParams struct {
	UserID  string
	Profile *model.Profile
}

type CreateProfileParams struct {
	UserID  string
	Profile *model.Profile
}

type DeleteProfileParams struct {
	UserID string
}

type ProfileRepository interface {
	GetProfile(ctx context.Context, arg *GetProfileParams) (*model.Profile, error)
	CreateProfile(ctx context.Context, arg *CreateProfileParams) (*model.Profile, error)
	UpdateProfile(ctx context.Context, arg *UpdateProfileParams) (*model.Profile, error)
	DeleteProfile(ctx context.Context, arg *DeleteProfileParams) error
}
