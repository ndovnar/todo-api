package service

import (
	"context"
	"todo/internal/converter"
	"todo/internal/dto"
	"todo/internal/model"
	"todo/internal/repository"
)

type TodoService interface {
	GetTodos(ctx context.Context, arg *GetTodosParams) ([]*dto.Todo, int64, error)
	GetTodo(ctx context.Context, arg *GetTodoParams) (*dto.Todo, error)
	CreateTodo(ctx context.Context, arg *CreateTodoParams) (*dto.Todo, error)
	UpdateTodo(ctx context.Context, arg *UpdateTodoParams) (*dto.Todo, error)
	DeleteTodo(ctx context.Context, arg *DeleteTodoParams) error
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

type GetTodosParams struct {
	Offset int64
	Limit  int64
	UserID string
}

func (s *todoService) GetTodos(ctx context.Context, arg *GetTodosParams) ([]*dto.Todo, int64, error) {
	todos, count, err := s.todoRepository.GetTodos(ctx, &repository.GetTodosParams{
		UserID: arg.UserID,
		Pagination: &model.Pagination{
			Offset: arg.Offset,
			Limit:  arg.Limit,
		},
	})

	return converter.TodosModelToDTO(todos), count, err
}

type GetTodoParams struct {
	ID     string
	UserID string
}

func (s *todoService) GetTodo(ctx context.Context, arg *GetTodoParams) (*dto.Todo, error) {
	todo, err := s.todoRepository.GetTodo(ctx, &repository.GetTodoParams{
		ID:     arg.ID,
		UserID: arg.UserID,
	})

	return converter.TodoModelToDTO(todo), err
}

type CreateTodoParams struct {
	Todo *dto.Todo
}

func (s *todoService) CreateTodo(ctx context.Context, arg *CreateTodoParams) (*dto.Todo, error) {
	todo, err := s.todoRepository.CreateTodo(ctx, &repository.CreateTodoParams{
		Todo: converter.TodoDTOToModel(arg.Todo),
	})

	return converter.TodoModelToDTO(todo), err
}

type UpdateTodoParams struct {
	ID     string
	UserID string
	Todo   *dto.Todo
}

func (s *todoService) UpdateTodo(ctx context.Context, arg *UpdateTodoParams) (*dto.Todo, error) {
	updatedTodo, err := s.todoRepository.UpdateTodo(ctx, &repository.UpdateTodoParams{
		ID:     arg.ID,
		UserID: arg.UserID,
		Todo:   converter.TodoDTOToModel(arg.Todo),
	})

	return converter.TodoModelToDTO(updatedTodo), err
}

type DeleteTodoParams struct {
	ID     string
	UserID string
}

func (s *todoService) DeleteTodo(ctx context.Context, arg *DeleteTodoParams) error {
	return s.todoRepository.DeleteTodo(ctx, &repository.DeleteTodoParams{
		ID:     arg.ID,
		UserID: arg.UserID,
	})
}
