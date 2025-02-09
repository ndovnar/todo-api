package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

type getTodosRequest struct {
	Offset int64 `form:"offset,default=0"`
	Limit  int64 `form:"limit,default=50"`
}

type getTodosResponse struct {
	Values []*dto.Todo
	Meta   *httphelper.Meta
}

func (h *Todos) HandleGetTodos(ctx *gin.Context) {
	var req getTodosRequest
	if err := ctx.ShouldBindQuery(req); err != nil {
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

	ctx.JSON(http.StatusOK, &getTodosResponse{
		Values: todos,
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

	ctx.JSON(http.StatusOK, todo)
}

type createTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h *Todos) HandleCreateTodo(ctx *gin.Context) {
	claims := auth.GetClaimsFromContext(ctx)

	var req createTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}
	todo, err := h.todoService.CreateTodo(ctx, &service.CreateTodoParams{
		Todo: &dto.Todo{
			Title:       req.Title,
			Description: req.Description,
			UserID:      claims.UserID,
		},
	})
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

type updateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h *Todos) HandleUpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	claims := auth.GetClaimsFromContext(ctx)

	var req updateTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	todo, err := h.todoService.UpdateTodo(ctx, &service.UpdateTodoParams{
		ID:     id,
		UserID: claims.UserID,
		Todo: &dto.Todo{
			Title:       req.Title,
			Description: req.Description,
		},
	})
	if err != nil {
		if err == db.ErrNotFound {
			ctx.Error(ginhelper.NewHttpError(http.StatusNotFound))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.JSON(http.StatusOK, todo)
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
