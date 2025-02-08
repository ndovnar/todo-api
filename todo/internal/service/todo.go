package service

import (
	"context"
	"todo/internal/converter"
	"todo/internal/dto"
	"todo/internal/model"
	"todo/internal/repository"
)

type TodoService interface {
	GetTodos(ctx context.Context, offset, limit int64, userID string) ([]*dto.Todo, int64, error)
	GetTodo(ctx context.Context, id, userID string) (*dto.Todo, error)
	CreateTodo(ctx context.Context, todo *dto.Todo) (*dto.Todo, error)
	UpdateTodo(ctx context.Context, id, userID string, todo *dto.Todo) (*dto.Todo, error)
	DeleteTodo(ctx context.Context, id, userID string) error
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
	UserID     string
	Pagination *model.Pagination
}

func (s *todoService) GetTodos(ctx context.Context, offset int64, limit int64, userID string) ([]*dto.Todo, int64, error) {
	todos, count, err := s.todoRepository.GetTodos(ctx, &repository.GetTodosParams{
		UserID: userID,
		Pagination: &model.Pagination{
			Offset: offset,
			Limit:  limit,
		},
	})

	return converter.TodosModelToDTO(todos), count, err
}

func (s *todoService) GetTodo(ctx context.Context, id, userID string) (*dto.Todo, error) {
	todo, err := s.todoRepository.GetTodo(ctx, &repository.GetTodoParams{
		ID:     id,
		UserID: userID,
	})

	return converter.TodoModelToDTO(todo), err
}

func (s *todoService) CreateTodo(ctx context.Context, todo *dto.Todo) (*dto.Todo, error) {
	Todo := converter.TodoDTOToModel(todo)

	newTodo, err := s.todoRepository.CreateTodo(ctx, &repository.CreateTodoParams{
		Todo: Todo,
	})

	return converter.TodoModelToDTO(newTodo), err
}

func (s *todoService) UpdateTodo(ctx context.Context, id, userID string, todo *dto.Todo) (*dto.Todo, error) {
	updatedTodo, err := s.todoRepository.UpdateTodo(ctx, &repository.UpdateTodoParams{
		ID:     id,
		UserID: userID,
		Todo:   converter.TodoDTOToModel(todo),
	})

	return converter.TodoModelToDTO(updatedTodo), err
}

func (s *todoService) DeleteTodo(ctx context.Context, id, userID string) error {
	return s.todoRepository.DeleteTodo(ctx, &repository.DeleteTodoParams{
		ID:     id,
		UserID: userID,
	})
}
