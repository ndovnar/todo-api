package service

import (
	"context"
	"encoding/json"
	"todo/internal/converter"
	"todo/internal/model"
	"todo/internal/repository"
	"todolib/kafka/events"
	"todolib/kafka/producer"

	"github.com/rs/zerolog/log"
)

type TodoService interface {
	GetTodos(ctx context.Context, arg *GetTodosParams) ([]*model.Todo, int64, error)
	GetTodo(ctx context.Context, arg *GetTodoParams) (*model.Todo, error)
	CreateTodo(ctx context.Context, arg *CreateTodoParams) (*model.Todo, error)
	UpdateTodo(ctx context.Context, arg *UpdateTodoParams) (*model.Todo, error)
	DeleteTodo(ctx context.Context, arg *DeleteTodoParams) error
	notifyTodoCreated(arg *model.Todo)
}

type todoService struct {
	todoRepository  repository.TodoRepository
	todoEventsTopic string
	kafkaProducer   *producer.Producer
}

type TodoServiceParams struct {
	TodoRepository  repository.TodoRepository
	TodoEventsTopic string
	KafkaProducer   *producer.Producer
}

func NewTodoService(arg *TodoServiceParams) TodoService {
	return &todoService{
		todoRepository:  arg.TodoRepository,
		todoEventsTopic: arg.TodoEventsTopic,
		kafkaProducer:   arg.KafkaProducer,
	}
}

type GetTodosParams struct {
	Offset int64
	Limit  int64
	UserID string
}

func (s *todoService) GetTodos(ctx context.Context, arg *GetTodosParams) ([]*model.Todo, int64, error) {
	todos, count, err := s.todoRepository.GetTodos(ctx, &repository.GetTodosParams{
		UserID: arg.UserID,
		Offset: arg.Offset,
		Limit:  arg.Limit,
	})

	if err != nil {
		return nil, 0, err
	}

	return converter.TodosDBModelToModel(todos), count, err
}

type GetTodoParams struct {
	ID     string
	UserID string
}

func (s *todoService) GetTodo(ctx context.Context, arg *GetTodoParams) (*model.Todo, error) {
	todo, err := s.todoRepository.GetTodo(ctx, &repository.GetTodoParams{
		ID:     arg.ID,
		UserID: arg.UserID,
	})

	if err != nil {
		return nil, err
	}

	return converter.TodoDBModelToModel(todo), err
}

type CreateTodoParams struct {
	Title       string
	Description string
	UserID      string
}

func (s *todoService) CreateTodo(ctx context.Context, arg *CreateTodoParams) (*model.Todo, error) {
	newTodo, err := s.todoRepository.CreateTodo(ctx, &repository.CreateTodoParams{
		UserID:      arg.UserID,
		Title:       arg.Title,
		Description: arg.Description,
	})

	if err != nil {
		return nil, err
	}

	todo := converter.TodoDBModelToModel(newTodo)

	s.notifyTodoCreated(todo)

	return todo, nil
}

type UpdateTodoParams struct {
	ID          string
	UserID      string
	Title       string
	IsCompleted bool
	Description string
}

func (s *todoService) UpdateTodo(ctx context.Context, arg *UpdateTodoParams) (*model.Todo, error) {
	updatedTodo, err := s.todoRepository.UpdateTodo(ctx, &repository.UpdateTodoParams{
		ID:          arg.ID,
		UserID:      arg.UserID,
		Title:       arg.Title,
		Description: arg.Description,
		IsCompleted: arg.IsCompleted,
	})

	if err != nil {
		return nil, err
	}

	return converter.TodoDBModelToModel(updatedTodo), err
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

func (s *todoService) notifyTodoCreated(todo *model.Todo) {
	event := events.TodoEvent{
		EventType: events.TodoCreatedEventType,
		UserID:    todo.UserID,
		TodoID:    todo.ID,
	}

	msg, err := json.Marshal(&event)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal event to JSON")
		return
	}

	s.kafkaProducer.Produce(s.todoEventsTopic, nil, msg)
}
