package main

import (
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	agcontroller "github.com/hafidz98/todo_api_app/activity-groups/controller"
	agrepository "github.com/hafidz98/todo_api_app/activity-groups/repository"
	agservice "github.com/hafidz98/todo_api_app/activity-groups/service"
	"github.com/hafidz98/todo_api_app/app"
	"github.com/hafidz98/todo_api_app/exception"
	"github.com/hafidz98/todo_api_app/helper"
	tdcontroller "github.com/hafidz98/todo_api_app/todos/controller"
	tdrepository "github.com/hafidz98/todo_api_app/todos/repository"
	tdservice "github.com/hafidz98/todo_api_app/todos/service"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

/*
	TODO:
	-asdfljasdflj
*/

func main() {
	log.Println("Todos API APP start") //TODO: aaaaaa

	DB := app.NewDB()
	validate := validator.New()
	var f string
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		f = strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		return f
	})

	activityGroupsRepository := agrepository.NewActivityGroupsRepository()
	activityGroupsService := agservice.NewActivityGroupsService(activityGroupsRepository, DB, validate)
	activityGroupsController := agcontroller.NewActivityGroupsController(activityGroupsService)

	todosRepository := tdrepository.NewTodosRepository()
	todosService := tdservice.NewTodosService(todosRepository, DB, validate)
	todosController := tdcontroller.NewTodosController(todosService)

	router := httprouter.New()
	router.GET("/activity-groups", activityGroupsController.SelectAll)
	router.GET("/activity-groups/:activityGroupId", activityGroupsController.SelectById)
	router.POST("/activity-groups", activityGroupsController.Create)
	router.PATCH("/activity-groups/:activityGroupId", activityGroupsController.Update)
	router.DELETE("/activity-groups/:activityGroupId", activityGroupsController.Delete)

	//?activity_group_id=
	router.GET("/todo-items", todosController.SelectAll)
	//router.GET("/todo-items", todosController.SelectById)
	router.GET("/todo-items/:todoId", todosController.SelectById)
	router.POST("/todo-items", todosController.Create)
	router.PATCH("/todo-items/:todoId", todosController.Update)
	router.DELETE("/todo-items/:todoId", todosController.Delete)

	router.PanicHandler = exception.ErrorHandler

	err := godotenv.Load(".env")
	helper.PanicIfError(err)
	httpPort := os.Getenv("HTTP_PORT")

	log.Println("Server listen to port " + httpPort)

	server := http.Server{
		Addr:    ":" + httpPort,
		Handler: router,
	}

	// instanceID := os.Getenv("INASTANCE_ID")
	// log.Println(instanceID)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
