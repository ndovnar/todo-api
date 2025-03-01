package converter

import (
	"todo/internal/dto"
	"todo/internal/model"
	"todo/internal/modeldb"
)

func DatesDBModelToModel(dates *modeldb.Dates) *model.Dates {
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
