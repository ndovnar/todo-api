package httpapi

import (
	"todolib/auth"
	"todolib/ginhelper"

	"todo/internal/httpapi/handlers"
)

func (a *HTTPAPI) regiesterRoutes() {
	todoHandlers := handlers.NewTodos(a.todoService)

	a.router.Use(ginhelper.ErrorMiddleware())

	authorized := a.router.Group("/")
	authorized.Use(auth.AuthMiddleware(a.publicKey))

	authorized.GET("/todos", todoHandlers.HandleGetTodos)
	authorized.GET("/todos/:id", todoHandlers.HandleGetTodo)
	authorized.POST("/todos", todoHandlers.HandleCreateTodo)
	authorized.PUT("/todos/:id", todoHandlers.HandleUpdateTodo)
	authorized.DELETE("/todos/:id", todoHandlers.HandleDeleteTodo)
}
