package controller

import (
	"net/http"
	"strconv"

	"github.com/hafidz98/todo_api_app/activity-groups/model/web"
	"github.com/hafidz98/todo_api_app/activity-groups/service"
	"github.com/hafidz98/todo_api_app/api"
	"github.com/hafidz98/todo_api_app/helper"
	"github.com/julienschmidt/httprouter"
)

type ActivityGroupsControllerImpl struct {
	ActivityGroupsService service.ActivityGroupsService
}

func NewActivityGroupsController(activityGroupsService service.ActivityGroupsService) ActivityGroupsController {
	return &ActivityGroupsControllerImpl{
		ActivityGroupsService: activityGroupsService,
	}
}

func (controller *ActivityGroupsControllerImpl) SelectAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	activityGroupResponse := controller.ActivityGroupsService.SelectAll(request.Context())
	apiResonse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activityGroupResponse,
	}

	helper.WriteToResponseBody(writer, apiResonse)
}

func (controller *ActivityGroupsControllerImpl) SelectById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("activityGroupId"))
	helper.PanicIfError(err)

	activityGroupResponse := controller.ActivityGroupsService.SelectById(request.Context(), id)
	apiResonse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activityGroupResponse,
	}

	helper.WriteToResponseBody(writer, apiResonse)
}

func (controller *ActivityGroupsControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	activityGroupCreateRequest := web.ActivityGroupsCreateRequest{}
	helper.ReadFromRequestBody(request, &activityGroupCreateRequest)

	//activityGroupResponse := controller.ActivityGroupsService.Create(request.Context(), activityGroupCreateRequest)
	id := controller.ActivityGroupsService.Create(request.Context(), activityGroupCreateRequest)
	activityGroupResponse := controller.ActivityGroupsService.SelectById(request.Context(), id.ID)
	apiResonse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activityGroupResponse,
	}

	helper.WriteToResponseBody(writer, apiResonse)
}

func (controller *ActivityGroupsControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	activityGroupUpdateRequest := web.ActivityGroupsUpdateRequest{}
	helper.ReadFromRequestBody(request, &activityGroupUpdateRequest)

	id, err := strconv.Atoi(params.ByName("activityGroupId"))
	helper.PanicIfError(err)

	activityGroupUpdateRequest.ID = id
	controller.ActivityGroupsService.Update(request.Context(), activityGroupUpdateRequest)

	activityGroupResponse := controller.ActivityGroupsService.SelectById(request.Context(), id)
	apiResonse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activityGroupResponse,
	}

	helper.WriteToResponseBody(writer, apiResonse)
}

func (controller *ActivityGroupsControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("activityGroupId"))
	helper.PanicIfError(err)

	controller.ActivityGroupsService.Delete(request.Context(), id)
	apiResonse := api.ApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    make(map[string]string),
	}

	helper.WriteToResponseBody(writer, apiResonse)
}
