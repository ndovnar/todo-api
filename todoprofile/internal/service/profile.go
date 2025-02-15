package service

import (
	"context"

	"todoprofile/internal/model"
	"todoprofile/internal/repository"
)

type ProfileService interface {
	GetProfile(ctx context.Context, arg *GetProfileParams) (*model.Profile, error)
	CreateProfile(ctx context.Context, arg *CreateProfileParams) (*model.Profile, error)
	UpdateProfile(ctx context.Context, arg *UpdateProfileParams) (*model.Profile, error)
	DeleteProfile(ctx context.Context, arg *DeleteProfileParams) error
}

type profileService struct {
	profileRepository repository.ProfileRepository
}

func NewProfileService(profileRepository repository.ProfileRepository) ProfileService {
	return &profileService{
		profileRepository: profileRepository,
	}
}

type GetProfileParams struct {
	ID     string
	UserID string
}

func (s *profileService) GetProfile(ctx context.Context, arg *GetProfileParams) (*model.Profile, error) {
	profile, err := s.profileRepository.GetProfile(ctx, &repository.GetProfileParams{
		UserID: arg.UserID,
	})

	return profile, err
}

type CreateProfileParams struct {
	UserID    string
	FirstName string
	LastName  string
}

func (s *profileService) CreateProfile(ctx context.Context, arg *CreateProfileParams) (*model.Profile, error) {
	profile, err := s.profileRepository.CreateProfile(ctx, &repository.CreateProfileParams{
		Profile: &model.Profile{
			UserID:    arg.UserID,
			FirstName: arg.FirstName,
			LastName:  arg.LastName,
		},
	})

	return profile, err
}

type UpdateProfileParams struct {
	UserID    string
	FirstName string
	LastName  string
}

func (s *profileService) UpdateProfile(ctx context.Context, arg *UpdateProfileParams) (*model.Profile, error) {
	updatedProfile, err := s.profileRepository.UpdateProfile(ctx, &repository.UpdateProfileParams{
		UserID: arg.UserID,
		Profile: &model.Profile{
			FirstName: arg.FirstName,
			LastName:  arg.LastName,
		},
	})

	return updatedProfile, err
}

type DeleteProfileParams struct {
	UserID string
}

func (s *profileService) DeleteProfile(ctx context.Context, arg *DeleteProfileParams) error {
	return s.profileRepository.DeleteProfile(ctx, &repository.DeleteProfileParams{
		UserID: arg.UserID,
	})
}
