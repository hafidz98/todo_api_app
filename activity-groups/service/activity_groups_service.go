package service

import (
	"context"

	"github.com/hafidz98/todo_api_app/activity-groups/model/web"
)

type ActivityGroupsService interface {
	SelectAll(context context.Context) []web.ActivityGroupsResponse
	SelectById(context context.Context, activityGroupId int) web.ActivityGroupsResponse
	Create(context context.Context, request web.ActivityGroupsCreateRequest) web.ActivityGroupsResponse
	Update(context context.Context, request web.ActivityGroupsUpdateRequest) web.ActivityGroupsResponse
	Delete(context context.Context, activityGroupId int)
}