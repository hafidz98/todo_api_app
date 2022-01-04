package activitygroups

import (
	"github.com/hafidz98/todo_api_app/activity-groups/model/domain"
	"github.com/hafidz98/todo_api_app/activity-groups/model/web"
)

func ToResponse(activityGroup domain.ActivityGroups) web.ActivityGroupsResponse {
	return web.ActivityGroupsResponse{
		ID:        activityGroup.ID,
		Email:     activityGroup.Email,
		Title:     activityGroup.Title,
		CreatedAt: activityGroup.CreatedAt,
		UpdatedAt: activityGroup.UpdatedAt,
		DeletedAt: activityGroup.DeletedAt,
	}
}

func ToResponses(activityGroups []domain.ActivityGroups) []web.ActivityGroupsResponse {
	var ActivityGroupsResponses []web.ActivityGroupsResponse
	for _, activityGroup := range activityGroups {
		ActivityGroupsResponses = append(ActivityGroupsResponses, ToResponse(activityGroup))
	}
	return ActivityGroupsResponses
}
