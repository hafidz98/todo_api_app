package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gobuffalo/nulls"
	activitygroups "github.com/hafidz98/todo_api_app/activity-groups"
	"github.com/hafidz98/todo_api_app/activity-groups/model/domain"
	"github.com/hafidz98/todo_api_app/activity-groups/model/web"
	"github.com/hafidz98/todo_api_app/activity-groups/repository"
	"github.com/hafidz98/todo_api_app/exception"
	"github.com/hafidz98/todo_api_app/helper"
)

type ActivityGroupsServiceImpl struct {
	ActivityGroupsRepository repository.ActivityGroupsRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewActivityGroupsService(activityRepository repository.ActivityGroupsRepository, DB *sql.DB, validate *validator.Validate) ActivityGroupsService {
	return &ActivityGroupsServiceImpl{
		ActivityGroupsRepository: activityRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *ActivityGroupsServiceImpl) SelectAll(context context.Context) []web.ActivityGroupsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	activityGroups := service.ActivityGroupsRepository.SelectAll(context, tx)
	return activitygroups.ToResponses(activityGroups)
}

func (service *ActivityGroupsServiceImpl) SelectById(context context.Context, activityGroupId int) web.ActivityGroupsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	activityGroup, err := service.ActivityGroupsRepository.SelectById(context, tx, activityGroupId)

	if err != nil {
		panic(exception.NewNotFound(strconv.Itoa(activityGroupId)))
	}

	return activitygroups.ToResponse(activityGroup)
}

func (service *ActivityGroupsServiceImpl) Create(context context.Context, request web.ActivityGroupsCreateRequest) web.ActivityGroupsResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	activityGroup := domain.ActivityGroups{
		Email:     request.Email,
		Title:     request.Title,
		CreatedAt: nulls.String{},
	}

	activityGroup = service.ActivityGroupsRepository.Create(context, tx, activityGroup)
	return activitygroups.ToResponse(activityGroup)
}

func (service *ActivityGroupsServiceImpl) Update(context context.Context, request web.ActivityGroupsUpdateRequest) web.ActivityGroupsResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	activityGroup, err := service.ActivityGroupsRepository.SelectById(context, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFound(strconv.Itoa(request.ID)))
	}

	activityGroup.Title = request.Title

	activityGroup = service.ActivityGroupsRepository.Update(context, tx, activityGroup)
	return activitygroups.ToResponse(activityGroup)
}

func (service *ActivityGroupsServiceImpl) Delete(context context.Context, activityGroupId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	activityGroup, err := service.ActivityGroupsRepository.SelectById(context, tx, activityGroupId)
	if err != nil {
		panic(exception.NewNotFound(strconv.Itoa(activityGroupId)))
	}

	service.ActivityGroupsRepository.Delete(context, tx, activityGroup)
}
