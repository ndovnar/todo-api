package repository

import (
	"context"
	"todostatistic/internal/modeldb"
)

type GetUserTodoStatisticParams struct {
	UserID string
	Month  string
	Year   string
}

type GetTodoStatisticParams struct {
	UserID string
	TodoID string
}

type CreateTodoStatisticParams struct {
	UserID string
	TodoID string
	Month  string
	Year   string
}

type StatisticRepository interface {
	GetUserTodoStatistic(ctx context.Context, arg *GetUserTodoStatisticParams) (*modeldb.UserTodoStatistic, error)
	GetTodoStatistic(ctx context.Context, arg *GetTodoStatisticParams) (*modeldb.TodoStatistic, error)
	CreateTodoStatistic(ctx context.Context, arg *CreateTodoStatisticParams) (*modeldb.TodoStatistic, error)
}
