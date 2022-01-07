package domain

import (
	"github.com/gobuffalo/nulls"
)

type Todos struct {
	ID              int
	ActivityGroupID int
	Title           string
	IsActive        bool
	Priority        string
	CreatedAt       nulls.String
	UpdatedAt       nulls.String
	DeletedAt       nulls.String
}
