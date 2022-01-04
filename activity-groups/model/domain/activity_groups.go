package domain

import (
	"github.com/gobuffalo/nulls"
)

type ActivityGroups struct {
	ID        int
	Email     string
	Title     string
	CreatedAt nulls.String
	UpdatedAt nulls.String
	DeletedAt nulls.String
}
