package domain

import "github.com/gobuffalo/nulls"

type Todos struct {
	ID              int
	ActivityGroupID nulls.String
	Title           string
	IsActive        string
	Priority        string
	CreatedAt       nulls.String
	UpdatedAt       nulls.String
	DeletedAt       nulls.String
}
