package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todolib/auth"
	"todolib/ginhelper"

	"todostatistic/internal/converters"
	"todostatistic/internal/service"
)

type Statistics struct {
	statisticService service.StatisticService
}

func NewStatistics(statisticService service.StatisticService) *Statistics {
	return &Statistics{
		statisticService: statisticService,
	}
}

func (h *Statistics) HandleGetStatistic(ctx *gin.Context) {
	claims := auth.GetClaimsFromContext(ctx)

	statisticCount, err := h.statisticService.GetUserTodoStatistic(ctx, claims.UserID)
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, converters.UserTodoStatisticModelToDTOResponse(statisticCount))
}
