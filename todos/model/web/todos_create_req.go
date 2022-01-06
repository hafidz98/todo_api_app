package web

type TodosCreateRequest struct {
	ActivityGroupId string `json:"activity_group_id"`
	Title           string `validate:"required" json:"title"`
}
