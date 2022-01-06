package repository

import (
	"context"
	"database/sql"

	"github.com/hafidz98/todo_api_app/todos/model/domain"
)

type TodosRepository interface {
	SelectAll(context context.Context, tx *sql.Tx) []domain.Todos
	SelectById(context context.Context, tx *sql.Tx, todosId int) (domain.Todos, error)
	Create(context context.Context, tx *sql.Tx, todos domain.Todos) domain.Todos
	Update(context context.Context, tx *sql.Tx, todos domain.Todos) domain.Todos
	Delete(context context.Context, tx *sql.Tx, todos domain.Todos)
}
