package converters

import (
	"todostatistic/internal/dto"
	"todostatistic/internal/model"
	"todostatistic/internal/modeldb"
)

func UserTodoStatisticModelDBToModel(statisticCount *modeldb.UserTodoStatistic) *model.UserTodoStatistic {
	return &model.UserTodoStatistic{
		Month: statisticCount.Month,
		Year:  statisticCount.Year,
		Total: statisticCount.Total,
	}
}

func UserTodoStatisticModelToDTOResponse(statisticCount *model.UserTodoStatistic) *dto.UserTodoStatisticResponse {
	return &dto.UserTodoStatisticResponse{
		Month: statisticCount.Month,
		Year:  statisticCount.Year,
		Total: statisticCount.Total,
	}
}

func TodoStatisticModelDBToModel(statistic *modeldb.TodoStatistic) *model.TodoStatistic {
	return &model.TodoStatistic{
		ID:     statistic.ID,
		TodoID: statistic.TodoID,
		UserID: statistic.TodoID,
		Month:  statistic.Month,
		Year:   statistic.Year,
	}
}
