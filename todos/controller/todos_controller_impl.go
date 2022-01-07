package controller

import (
	"net/http"
	"strconv"

	"github.com/hafidz98/todo_api_app/api"
	"github.com/hafidz98/todo_api_app/helper"
	"github.com/hafidz98/todo_api_app/todos/model/web"
	"github.com/hafidz98/todo_api_app/todos/service"
	"github.com/julienschmidt/httprouter"
)

type TodosControllerImpl struct {
	TodosService service.TodosService
}

func NewTodosController(TodosService service.TodosService) TodosController {
	return &TodosControllerImpl{
		TodosService: TodosService,
	}
}

func (controller *TodosControllerImpl) SelectAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	agId := request.URL.Query().Get("activity_group_id")
	//request.URL.Query().Get()
	var todoResponse []web.TodosResponse
	if agId != "" {
		todoResponse = controller.TodosService.SelectByAgId(request.Context(), agId)
		if todoResponse == nil {
			todoResponse = make([]web.TodosResponse, 0)
		}
	} else {
		todoResponse = controller.TodosService.SelectAll(request.Context())
	}

	apiResonse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todoResponse,
	}

	helper.WriteToResponseBody(writer, apiResonse)
}

func (controller *TodosControllerImpl) SelectById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("todoId"))

	helper.PanicIfError(err)
	//log.Println("id di ctr todos :" + strconv.Itoa(id))

	todoResponse := controller.TodosService.SelectById(request.Context(), id)
	apiResonse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todoResponse,
	}

	helper.WriteToResponseBody(writer, apiResonse)
}

func (controller *TodosControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	todoCreateRequest := web.TodosCreateRequest{}
	helper.ReadFromRequestBody(request, &todoCreateRequest)

	//activityGroupResponse := controller.TodosService.Create(request.Context(), activityGroupCreateRequest)
	id := controller.TodosService.Create(request.Context(), todoCreateRequest)
	todoResponse := controller.TodosService.SelectById(request.Context(), id.ID)
	apiResonse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todoResponse,
	}

	helper.WriteToResponseBody201(writer, apiResonse)
}

func (controller *TodosControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoUpdateRequest := web.TodosUpdateRequest{}
	helper.ReadFromRequestBody(request, &todoUpdateRequest)

	id, err := strconv.Atoi(params.ByName("todoId"))
	helper.PanicIfError(err)

	todoUpdateRequest.ID = id
	todoResponse := controller.TodosService.Update(request.Context(), todoUpdateRequest)

	// todoResponse := controller.TodosService.SelectById(request.Context(), id)
	// log.Println(todoResponse)
	apiResponse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todoResponse,
	}
	// log.Println(apiResponse)
	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *TodosControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("todoId"))
	helper.PanicIfError(err)

	controller.TodosService.Delete(request.Context(), id)
	apiResonse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    make(map[string]string),
	}

	helper.WriteToResponseBody(writer, apiResonse)
}
