package repository

import (
	"context"
	"todo/internal/model"
)

type GetTodosParams struct {
	UserID     string
	Pagination *model.Pagination
}

type GetTodoParams struct {
	ID string
	UserID string
}

type UpdateTodoParams struct {
	ID string
	Todo   *model.Todo
	UserID string
}

type CreateTodoParams struct {
	Todo *model.Todo
}

type DeleteTodoParams struct {
	ID string
	UserID string
}

type TodoRepository interface {
	GetTodos(ctx context.Context, arg *GetTodosParams) ([]*model.Todo, int64, error)
	GetTodo(ctx context.Context, arg *GetTodoParams) (*model.Todo, error)
	CreateTodo(ctx context.Context, arg *CreateTodoParams) (*model.Todo, error)
	UpdateTodo(ctx context.Context, arg *UpdateTodoParams) (*model.Todo, error)
	DeleteTodo(ctx context.Context, arg *DeleteTodoParams) error
}
