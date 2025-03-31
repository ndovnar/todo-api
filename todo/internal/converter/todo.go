package converter

import (
	"todo/internal/dto"
	"todo/internal/model"
	"todo/internal/modeldb"

	"github.com/samber/lo"
)

func TodosDBModelToModel(todos []*modeldb.Todo) []*model.Todo {
	return lo.Map(todos, func(todo *modeldb.Todo, index int) *model.Todo {
		return TodoDBModelToModel(todo)
	})
}

func TodoDBModelToModel(todo *modeldb.Todo) *model.Todo {
	return &model.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		UserID:      todo.UserID,
		IsCompleted: todo.IsCompleted,
		Dates:       DatesDBModelToModel(todo.Dates),
	}
}

func TodosModelToDTOresponse(todos []*model.Todo) []*dto.TodoResponse {
	return lo.Map(todos, func(todo *model.Todo, index int) *dto.TodoResponse {
		return TodoModelToDTOResponse(todo)
	})
}

func TodoModelToDTOResponse(todo *model.Todo) *dto.TodoResponse {
	return &dto.TodoResponse{
		ID:          todo.ID,
		UserID:      todo.UserID,
		Title:       todo.Title,
		Description: todo.Description,
		Dates:       DatesModelToDTOResponse(todo.Dates),
	}
}
