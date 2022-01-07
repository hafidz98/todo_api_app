package todos

import (
	"github.com/hafidz98/todo_api_app/todos/model/domain"
	"github.com/hafidz98/todo_api_app/todos/model/web"
)

func ToResponse(todos domain.Todos) web.TodosResponse {
	return web.TodosResponse{
		ID:              todos.ID,
		ActivityGroupId: todos.ActivityGroupID,
		Title:           todos.Title,
		IsActive:        todos.IsActive,
		Priority:        todos.Priority,
		CreatedAt:       todos.CreatedAt,
		UpdatedAt:       todos.UpdatedAt,
		DeletedAt:       todos.DeletedAt,
	}
}

func ToResponses(todos []domain.Todos) []web.TodosResponse {
	var TodosResponses []web.TodosResponse
	for _, todo := range todos {
		todo := web.TodosResponse{
			ID:              todo.ID,
			ActivityGroupId: todo.ActivityGroupID,
			Title:           todo.Title,
			IsActive:        todo.IsActive,
			Priority:        todo.Priority,
			CreatedAt:       todo.CreatedAt,
			UpdatedAt:       todo.UpdatedAt,
			DeletedAt:       todo.DeletedAt,
		}
		TodosResponses = append(TodosResponses, todo)
	}
	return TodosResponses
}
