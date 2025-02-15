package httpapi

import (
	"todolib/auth"
	"todolib/ginhelper"

	"todoprofile/internal/httpapi/handlers"
)

func (a *HTTPAPI) regiesterRoutes() {
	profileHandlers := handlers.NewProfiles(a.profileService)

	a.router.Use(ginhelper.ErrorMiddleware())

	authorized := a.router.Group("/")
	authorized.Use(auth.AuthMiddleware(a.publicKey))

	authorized.GET("/profile", profileHandlers.HandleGetProfile)
	authorized.POST("/profile", profileHandlers.HandleCreateProfile)
	authorized.PUT("/profile", profileHandlers.HandleUpdateProfile)
	authorized.DELETE("/profile", profileHandlers.HandleDeleteProfile)
}
