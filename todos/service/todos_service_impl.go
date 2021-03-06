package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/hafidz98/todo_api_app/exception"
	"github.com/hafidz98/todo_api_app/helper"
	"github.com/hafidz98/todo_api_app/todos"
	"github.com/hafidz98/todo_api_app/todos/model/domain"
	"github.com/hafidz98/todo_api_app/todos/model/web"
	"github.com/hafidz98/todo_api_app/todos/repository"
)

type TodosServiceImpl struct {
	TodosRepository repository.TodosRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewTodosService(todoRepository repository.TodosRepository, DB *sql.DB, validate *validator.Validate) TodosService {
	return &TodosServiceImpl{
		TodosRepository: todoRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *TodosServiceImpl) SelectAll(context context.Context) []web.TodosResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todosService := service.TodosRepository.SelectAll(context, tx)
	return todos.ToResponses(todosService)
}

func (service *TodosServiceImpl) SelectById(context context.Context, todoId int) web.TodosResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todosResult, err := service.TodosRepository.SelectById(context, tx, todoId)

	if err != nil {
		panic(exception.NewNotFound(strconv.Itoa(todoId), "Todo"))
	}
	//log.Println(todosResult)
	return todos.ToResponse(todosResult)
}

func (service *TodosServiceImpl) SelectByAgId(context context.Context, agId string) []web.TodosResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todosResult := service.TodosRepository.SelectByAgId(context, tx, agId)

	return todos.ToResponses(todosResult)
}

func (service *TodosServiceImpl) Create(context context.Context, request web.TodosCreateRequest) web.TodosResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//id, _ := request.ActivityGroupId.Int64()
	todosReq := domain.Todos{
		ActivityGroupID: request.ActivityGroupId,
		Title:           request.Title,
	}
	//log.Println(todosReq)

	todosReq = service.TodosRepository.Create(context, tx, todosReq)
	return todos.ToResponse(todosReq)
}

func (service *TodosServiceImpl) Update(context context.Context, request web.TodosUpdateRequest) web.TodosResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodosRepository.SelectById(context, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFound(strconv.Itoa(request.ID), "Todo"))
	}

	if request.Title != "" {
		todo.Title = request.Title
	}
	todo.IsActive = request.IsActive

	todo = service.TodosRepository.Update(context, tx, todo)
	// log.Println("Di Service")
	// log.Println(todo)
	return todos.ToResponse(todo)
}

func (service *TodosServiceImpl) Delete(context context.Context, todoId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodosRepository.SelectById(context, tx, todoId)
	if err != nil {
		panic(exception.NewNotFound(strconv.Itoa(todoId), "Todo"))
	}

	service.TodosRepository.Delete(context, tx, todo)
}
