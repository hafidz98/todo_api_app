package web

type TodosCreateRequest struct {
	ActivityGroupId int    `json:"activity_group_id" validate:"required"`
	Title           string `json:"title" validate:"required"`
}
