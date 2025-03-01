package converter

import (
	"todoauth/internal/model"
	"todoauth/internal/modeldb"
)

func DatesDBModelToModel(dates *modeldb.Dates) *model.Dates {
	return &model.Dates{
		Created:  dates.Created,
		Modified: dates.Modified,
		Deleted:  dates.Deleted,
	}
}
