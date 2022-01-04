package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/hafidz98/todo_api_app/activity-groups/controller"
	"github.com/hafidz98/todo_api_app/activity-groups/repository"
	"github.com/hafidz98/todo_api_app/activity-groups/service"
	"github.com/hafidz98/todo_api_app/app"
	"github.com/hafidz98/todo_api_app/exception"
	"github.com/hafidz98/todo_api_app/helper"
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
	activityGroupsRepository := repository.NewActivityGroupsRepository()
	activityGroupsService := service.NewActivityGroupsService(activityGroupsRepository, DB, validate)
	activityGroupsController := controller.NewActivityGroupsController(activityGroupsService)

	router := httprouter.New()
	router.GET("/activity-groups", activityGroupsController.SelectAll)
	router.GET("/activity-groups/:activityGroupId", activityGroupsController.SelectById)
	router.POST("/activity-groups", activityGroupsController.Create)
	router.PATCH("/activity-groups/:activityGroupId", activityGroupsController.Update)
	router.DELETE("/activity-groups/:activityGroupId", activityGroupsController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    ":3030",
		Handler: router,
	}

	// instanceID := os.Getenv("INASTANCE_ID")
	// log.Println(instanceID)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
