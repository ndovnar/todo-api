package converter

import (
	"todoprofile/internal/dto"
	"todoprofile/internal/model"
)

func ProfileModelToDTO(profile *model.Profile) *dto.Profile {
	return &dto.Profile{
		ID:        profile.ID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Dates:     DatesModelToDTO(profile.Dates),
	}
}

func ProfileDtoToModel(profile *dto.Profile) *model.Profile {
	return &model.Profile{
		ID:        profile.ID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Dates:     DatesDTOtoModel(profile.Dates),
	}
}
