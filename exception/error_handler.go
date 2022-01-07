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

//func msgForTag(validator.FieldError)
// func (fld reflect.StructField) msgVal() string {
// 	msg, _ := fld.Tag.Lookup("json")
// 	return msg
// }

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	//validator := new()
	//val := validator.New()

	exception, ok := err.(validator.ValidationErrors)
	//log.Println(exception)

	var msg string
	// excep := exception[0]
	// msg := excep.Field()

	for _, excep := range exception {
		msg = excep.Field()
		// f, _ := reflect.TypeOf(args).Elem().FieldByName(msg)
		// msg, _ = f.Tag.Lookup("json")
	}

	if ok {
		writer.Header().Set("content-type", "applicaton/json")
		writer.WriteHeader(http.StatusBadRequest)

		apiResponses := api.ApiResponse{
			Status:  "Bad Request",
			Message: msg + " cannot be null",
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
			Message: exception.From + " with ID " + exception.Error + " Not Found",
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
