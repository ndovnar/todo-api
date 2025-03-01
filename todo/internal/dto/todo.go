package dto

import "todolib/httphelper"

type TodoResponse struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	UserID      string         `json:"userId"`
	Description string         `json:"description"`
	IsCompleted bool           `json:"completed"`
	Dates       *DatesResponse `json:"dates"`
}

type GetTodosRequest struct {
	Offset int64 `form:"offset,default=0"`
	Limit  int64 `form:"limit,default=50"`
}

type GetTodosResponse struct {
	Values []*TodoResponse  `json:"values"`
	Meta   *httphelper.Meta `json:"meta"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsCompleted bool   `json:"completed"`
}
