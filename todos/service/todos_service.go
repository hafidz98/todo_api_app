package service

import (
	"context"

	"github.com/hafidz98/todo_api_app/todos/model/web"
)

type TodosService interface {
	SelectAll(context context.Context) []web.TodosResponse
	SelectById(context context.Context, todosId int) web.TodosResponse
	SelectByAgId(context context.Context, agId string) []web.TodosResponse
	Create(context context.Context, request web.TodosCreateRequest) web.TodosResponse
	Update(context context.Context, request web.TodosUpdateRequest) web.TodosResponse
	Delete(context context.Context, todosId int)
}
