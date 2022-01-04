package exception

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/hafidz98/todo_api_app/api"
	"github.com/hafidz98/todo_api_app/helper"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("content-type", "applicaton/json")
		writer.WriteHeader(http.StatusBadRequest)

		apiResponses := api.ApiResponse{
			Status:  "Bad Request",
			Message: "title cannot be null",
			Data:    make(map[string]string),
		}

		helper.WriteToResponseBody(writer, apiResponses)
		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFound)
	if ok {
		writer.Header().Set("content-type", "applicaton/json")
		writer.WriteHeader(http.StatusNotFound)

		apiResponses := api.ApiResponse{
			Status:  "Not Found",
			Message: "Activity with ID " + exception.Error + " Not Found",
			Data:    make(map[string]string),
		}

		helper.WriteToResponseBody(writer, apiResponses)
		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("content-type", "applicaton/json")
	writer.WriteHeader(http.StatusInternalServerError)

	apiResponses := api.ApiResponse{
		Status:  "Internal Server Error",
		Message: "Server Error",
		Data:    err,
	}

	helper.WriteToResponseBody(writer, apiResponses)
}
