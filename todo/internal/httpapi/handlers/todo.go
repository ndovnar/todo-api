package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/internal/converter"
	"todo/internal/dto"
	"todo/internal/service"

	"todolib/auth"
	"todolib/db"
	"todolib/ginhelper"
	"todolib/httphelper"
)

type Todos struct {
	todoService service.TodoService
}

func NewTodos(todoService service.TodoService) *Todos {
	return &Todos{
		todoService: todoService,
	}
}

func (h *Todos) HandleGetTodos(ctx *gin.Context) {
	var req dto.GetTodosRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	claims := auth.GetClaimsFromContext(ctx)

	todos, count, err := h.todoService.GetTodos(ctx, &service.GetTodosParams{
		Offset: req.Offset,
		Limit:  req.Limit,
		UserID: claims.UserID,
	})
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, &dto.GetTodosResponse{
		Values: converter.TodosModelToDTOresponse(todos),
		Meta:   httphelper.NewMeta(count),
	})
}

func (h *Todos) HandleGetTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	claims := auth.GetClaimsFromContext(ctx)

	todo, err := h.todoService.GetTodo(ctx, &service.GetTodoParams{
		ID:     id,
		UserID: claims.UserID,
	})
	if err != nil {
		if err == db.ErrNotFound {
			ctx.Error(ginhelper.NewHttpError(http.StatusNotFound))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.JSON(http.StatusOK, converter.TodoModelToDTOResponse(todo))
}

func (h *Todos) HandleCreateTodo(ctx *gin.Context) {
	claims := auth.GetClaimsFromContext(ctx)

	var req dto.CreateTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}
	todo, err := h.todoService.CreateTodo(ctx, &service.CreateTodoParams{
		Title:       req.Title,
		Description: req.Description,
		UserID:      claims.UserID,
	})
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, converter.TodoModelToDTOResponse(todo))
}

func (h *Todos) HandleUpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	claims := auth.GetClaimsFromContext(ctx)

	var req dto.UpdateTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	todo, err := h.todoService.UpdateTodo(ctx, &service.UpdateTodoParams{
		ID:          id,
		UserID:      claims.UserID,
		Title:       req.Title,
		Description: req.Description,
		IsCompleted: req.IsCompleted,
	})
	if err != nil {
		if err == db.ErrNotFound {
			ctx.Error(ginhelper.NewHttpError(http.StatusNotFound))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.JSON(http.StatusOK, converter.TodoModelToDTOResponse(todo))
}

func (h *Todos) HandleDeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	claims := auth.GetClaimsFromContext(ctx)

	err := h.todoService.DeleteTodo(ctx, &service.DeleteTodoParams{
		ID:     id,
		UserID: claims.UserID,
	})
	if err != nil {
		if err == db.ErrNotFound {
			ctx.Error(ginhelper.NewHttpError(http.StatusNotFound))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.Status(http.StatusOK)
}
