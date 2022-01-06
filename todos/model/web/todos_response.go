package web

import (
	"github.com/gobuffalo/nulls"
)

type TodosResponse struct {
	ID              int          `json:"id"`
	ActivityGroupId nulls.String `json:"activity_group_id"`
	Title           string       `json:"title"`
	IsActive        string       `json:"is_active"`
	Priority        string       `json:"priority"`
	CreatedAt       nulls.String `json:"created_at"`
	UpdatedAt       nulls.String `json:"updated_at"`
	DeletedAt       nulls.String `json:"deleted_at"`
}
