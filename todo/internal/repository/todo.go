package repository

import (
	"context"
	"todo/internal/modeldb"
)

type GetTodosParams struct {
	UserID string
	Limit  int64
	Offset int64
}

type GetTodoParams struct {
	ID     string
	UserID string
}

type UpdateTodoParams struct {
	ID          string
	UserID      string
	Title       string
	Description string
	IsCompleted bool
}

type CreateTodoParams struct {
	UserID      string
	Title       string
	Description string
}

type DeleteTodoParams struct {
	ID     string
	UserID string
}

type TodoRepository interface {
	GetTodos(ctx context.Context, arg *GetTodosParams) ([]*modeldb.Todo, int64, error)
	GetTodo(ctx context.Context, arg *GetTodoParams) (*modeldb.Todo, error)
	CreateTodo(ctx context.Context, arg *CreateTodoParams) (*modeldb.Todo, error)
	UpdateTodo(ctx context.Context, arg *UpdateTodoParams) (*modeldb.Todo, error)
	DeleteTodo(ctx context.Context, arg *DeleteTodoParams) error
}
