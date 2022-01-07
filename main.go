package main

import (
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
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

func init() {
	app.NewMigrate()
}

func main() {
	log.Println("Todos API APP start") //TODO: aaaaaa

	runtime.GOMAXPROCS(2)

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
	go router.GET("/activity-groups", activityGroupsController.SelectAll)
	go router.GET("/activity-groups/:activityGroupId", activityGroupsController.SelectById)
	go router.POST("/activity-groups", activityGroupsController.Create)
	go router.PATCH("/activity-groups/:activityGroupId", activityGroupsController.Update)
	go router.DELETE("/activity-groups/:activityGroupId", activityGroupsController.Delete)

	//?activity_group_id=
	go router.GET("/todo-items", todosController.SelectAll)
	//router.GET("/todo-items", todosController.SelectById)
	go router.GET("/todo-items/:todoId", todosController.SelectById)
	go router.POST("/todo-items", todosController.Create)
	go router.PATCH("/todo-items/:todoId", todosController.Update)
	go router.DELETE("/todo-items/:todoId", todosController.Delete)

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
