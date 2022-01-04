package web

import (
	"github.com/gobuffalo/nulls"
)

type ActivityGroupsResponse struct {
	ID        int            `json:"id"`
	Email     string         `json:"email"`
	Title     string         `json:"title"`
	CreatedAt nulls.String `json:"created_at"`
	UpdatedAt nulls.String `json:"updated_at"`
	DeletedAt nulls.String `json:"deleted_at"`
}
