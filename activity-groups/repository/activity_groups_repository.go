package repository

import (
	"context"
	"database/sql"

	"github.com/hafidz98/todo_api_app/activity-groups/model/domain"
)

type ActivityGroupsRepository interface {
	SelectAll(context context.Context, tx *sql.Tx) []domain.ActivityGroups
	SelectById(context context.Context, tx *sql.Tx, activityGroupId int) (domain.ActivityGroups, error)
	Create(context context.Context, tx *sql.Tx, activityGroup domain.ActivityGroups) domain.ActivityGroups
	Update(context context.Context, tx *sql.Tx, activityGroup domain.ActivityGroups) domain.ActivityGroups
	Delete(context context.Context, tx *sql.Tx, activityGroup domain.ActivityGroups)
}
