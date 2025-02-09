package converter

import (
	"todo/internal/dto"
	"todo/internal/model"
)

func TodoModelToDTO(todo *model.Todo) *dto.Todo {
	return &dto.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
		Dates:       DatesModelToDTO(todo.Dates),
	}
}

func TodosModelToDTO(todos []*model.Todo) []*dto.Todo {
	converted := make([]*dto.Todo, len(todos))

	for _, todo := range todos {
		converted = append(converted, TodoModelToDTO(todo))
	}

	return converted
}

func TodoDTOToModel(todo *dto.Todo) *model.Todo {
	return &model.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
		Dates:       DatesDTOtoModel(todo.Dates),
	}
}

func TodosDTOToModel(todos []*dto.Todo) []*model.Todo {
	converted := make([]*model.Todo, len(todos))

	for _, todo := range todos {
		converted = append(converted, TodoDTOToModel(todo))
	}

	return converted
}
