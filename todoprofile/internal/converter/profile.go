package converter

import (
	"todoprofile/internal/dto"
	"todoprofile/internal/model"
	"todoprofile/internal/modeldb"
)

func ProfileModelDBToModel(profile *modeldb.Profile) *model.Profile {
	return &model.Profile{
		ID:        profile.ID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		UserID:    profile.UserID,
		Dates:     DatesDBToModel(profile.Dates),
	}
}

func ProfileModelToDTOResponse(profile *model.Profile) *dto.ProfileResponse {
	return &dto.ProfileResponse{
		ID:        profile.ID,
		UserID:    profile.UserID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Dates:     DatesModelToDTOResponse(profile.Dates),
	}
}
