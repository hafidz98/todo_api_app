package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafidz98/todo_api_app/helper"
	"github.com/hafidz98/todo_api_app/todos/model/domain"
)

type TodosRepositoryImpl struct{}

func NewTodosRepository() TodosRepository {
	return &TodosRepositoryImpl{}
}

func (repository *TodosRepositoryImpl) SelectAll(context context.Context, tx *sql.Tx) []domain.Todos {
	query := "SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at FROM todos"
	rows, err := tx.QueryContext(context, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var todos []domain.Todos
	for rows.Next() {
		todo := domain.Todos{}
		err := rows.Scan(
			&todo.ID,
			&todo.ActivityGroupID,
			&todo.Title,
			&todo.IsActive,
			&todo.Priority,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.DeletedAt,
		)
		helper.PanicIfError(err)
		todos = append(todos, todo)
	}

	return todos
}

func (repository *TodosRepositoryImpl) SelectById(context context.Context, tx *sql.Tx, todoId int) (domain.Todos, error) {
	query := "SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at FROM todos WHERE id=?"
	rows, err := tx.QueryContext(context, query, todoId)
	helper.PanicIfError(err)
	defer rows.Close()

	todo := domain.Todos{}
	if rows.Next() {
		err := rows.Scan(
			&todo.ID,
			&todo.ActivityGroupID,
			&todo.Title,
			&todo.IsActive,
			&todo.Priority,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.DeletedAt,
		)
		helper.PanicIfError(err)
		return todo, nil
	}

	return todo, errors.New("todos not found")
}

func (repository *TodosRepositoryImpl) SelectByAgId(context context.Context, tx *sql.Tx, agId string) []domain.Todos {
	query := "SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at FROM todos WHERE activity_group_id=?"
	rows, err := tx.QueryContext(context, query, agId)
	helper.PanicIfError(err)
	defer rows.Close()

	var todos []domain.Todos
	for rows.Next() {
		todo := domain.Todos{}
		err := rows.Scan(
			&todo.ID,
			&todo.ActivityGroupID,
			&todo.Title,
			&todo.IsActive,
			&todo.Priority,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.DeletedAt,
		)
		helper.PanicIfError(err)
		todos = append(todos, todo)
	}

	return todos
}

func (repository *TodosRepositoryImpl) Create(context context.Context, tx *sql.Tx, todo domain.Todos) domain.Todos {
	query := "INSERT INTO todos(activity_group_id,title) VALUES (?,?)"
	result, err := tx.ExecContext(context, query, todo.ActivityGroupID, todo.Title)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todo.ID = int(id)
	return todo
}

func (repository *TodosRepositoryImpl) Update(context context.Context, tx *sql.Tx, todo domain.Todos) domain.Todos {
	query := "UPDATE todos SET title = COALESCE(?, title), is_active = COALESCE(?, is_active) WHERE id = ?"
	_, err := tx.ExecContext(context, query, todo.Title, todo.IsActive, todo.ID)
	helper.PanicIfError(err)
	return todo
}

func (repository *TodosRepositoryImpl) Delete(context context.Context, tx *sql.Tx, todo domain.Todos) {
	query := "DELETE FROM todos WHERE id = ?"
	_, err := tx.ExecContext(context, query, todo.ID)
	helper.PanicIfError(err)
}
