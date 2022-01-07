package web

type TodosUpdateRequest struct {
	ID       int    `validate:"required" json:"id"`
	Title    string `json:"title"`
	IsActive bool   `json:"is_active"`
}
