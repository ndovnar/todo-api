package httpapi

import (
	"todolib/auth"
	"todolib/ginhelper"

	"todostatistic/internal/httpapi/handlers"
)

func (a *HTTPAPI) regiesterRoutes() {
	statisticHandlers := handlers.NewStatistics(a.statisticService)

	a.router.Use(ginhelper.ErrorMiddleware())

	authorized := a.router.Group("/")
	authorized.Use(auth.AuthMiddleware(a.publicKey))

	authorized.GET("/statistic", statisticHandlers.HandleGetStatistic)
}
