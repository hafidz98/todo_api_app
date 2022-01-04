package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ActivityGroupsController interface {
	SelectAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SelectById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
