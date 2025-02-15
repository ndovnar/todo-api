package converter

import (
	"todoprofile/internal/dto"
	"todoprofile/internal/model"
)

func DatesModelToDTO(dates *model.Dates) *dto.Dates {
	return &dto.Dates{
		Created:  dates.Created,
		Modified: dates.Modified,
		Deleted:  dates.Deleted,
	}
}

func DatesDTOtoModel(dates *dto.Dates) *model.Dates {
	return &model.Dates{
		Created:  dates.Created,
		Modified: dates.Modified,
		Deleted:  dates.Deleted,
	}
}
