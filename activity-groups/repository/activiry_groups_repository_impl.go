package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafidz98/todo_api_app/activity-groups/model/domain"
	"github.com/hafidz98/todo_api_app/helper"
)

type ActivityGroupsRepositoryImpl struct{}

func NewActivityGroupsRepository() ActivityGroupsRepository {
	return &ActivityGroupsRepositoryImpl{}
}

func (repository *ActivityGroupsRepositoryImpl) SelectAll(context context.Context, tx *sql.Tx) []domain.ActivityGroups {
	query := "SELECT id, email, title, created_at, updated_at, deleted_at FROM activities"
	rows, err := tx.QueryContext(context, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var activitGroups []domain.ActivityGroups
	for rows.Next() {
		activityGroup := domain.ActivityGroups{}
		err := rows.Scan(
			&activityGroup.ID,
			&activityGroup.Email,
			&activityGroup.Title,
			&activityGroup.CreatedAt,
			&activityGroup.UpdatedAt,
			&activityGroup.DeletedAt,
		)
		helper.PanicIfError(err)
		activitGroups = append(activitGroups, activityGroup)
	}

	return activitGroups
}

func (repository *ActivityGroupsRepositoryImpl) SelectById(context context.Context, tx *sql.Tx, activityGroupId int) (domain.ActivityGroups, error) {
	query := "SELECT id, email, title, created_at, updated_at, deleted_at FROM activities WHERE id=?"
	rows, err := tx.QueryContext(context, query, activityGroupId)
	helper.PanicIfError(err)
	defer rows.Close()

	activitGroup := domain.ActivityGroups{}
	if rows.Next() {
		err := rows.Scan(
			&activitGroup.ID,
			&activitGroup.Email,
			&activitGroup.Title,
			&activitGroup.CreatedAt,
			&activitGroup.UpdatedAt,
			&activitGroup.DeletedAt,
		)
		helper.PanicIfError(err)
		return activitGroup, nil
	}

	return activitGroup, errors.New("activity group not found")
}

func (repository *ActivityGroupsRepositoryImpl) Create(context context.Context, tx *sql.Tx, activityGroup domain.ActivityGroups) domain.ActivityGroups {
	query := "INSERT INTO activities(email,title) VALUES (?,?)"
	result, err := tx.ExecContext(context, query, activityGroup.Email, activityGroup.Title)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	activityGroup.ID = int(id)
	return activityGroup
}

func (repository *ActivityGroupsRepositoryImpl) Update(context context.Context, tx *sql.Tx, activityGroup domain.ActivityGroups) domain.ActivityGroups {
	query := "UPDATE activities SET title = ? WHERE id = ?"
	_, err := tx.ExecContext(context, query, activityGroup.Title, activityGroup.ID)
	helper.PanicIfError(err)
	return activityGroup
}

func (repository *ActivityGroupsRepositoryImpl) Delete(context context.Context, tx *sql.Tx, activityGroup domain.ActivityGroups) {
	query := "DELETE FROM activities WHERE id = ?"
	_, err := tx.ExecContext(context, query, activityGroup.ID)
	helper.PanicIfError(err)
}
