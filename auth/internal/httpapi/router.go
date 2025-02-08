package httpapi

import (
	"auth/internal/httpapi/handlers"
	"lib/auth"
	"lib/ginhelper"
)

func (a *HTTPAPI) regiesterRoutes() {
	tokenHandlers := handlers.NewTokens(a.authService)
	userHandlers := handlers.NewUsers(a.userService)

	a.router.Use(ginhelper.ErrorMiddleware())

	authorized := a.router.Group("/")
	authorized.Use(auth.AuthMiddleware(a.publicKey))

	authorized.POST("/logout", tokenHandlers.HandleLogout)
	a.router.POST("/login", tokenHandlers.HandleLogin)
	a.router.POST("/users", userHandlers.CreateUser)
	a.router.POST("/tokens/renew/access", tokenHandlers.HandleRenewAccessToken)
	a.router.POST("/tokens/renew/refresh", tokenHandlers.HandleRenewRefreshToken)
}
