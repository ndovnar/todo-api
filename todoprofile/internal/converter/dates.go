package converter

import (
	"todoprofile/internal/dto"
	"todoprofile/internal/model"
	"todoprofile/internal/modeldb"
)

func DatesDBToModel(dates *modeldb.Dates) *model.Dates {
	return &model.Dates{
		Created:  dates.Created,
		Modified: dates.Modified,
		Deleted:  dates.Deleted,
	}
}

func DatesModelToDTOResponse(dates *model.Dates) *dto.DatesResponse {
	return &dto.DatesResponse{
		Created:  dates.Created,
		Modified: dates.Modified,
		Deleted:  dates.Deleted,
	}
}
