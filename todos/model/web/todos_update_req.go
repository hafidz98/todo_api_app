package web

type TodosUpdateRequest struct {
	ID    int    `json:"id"`
	Title string `validate:"required" json:"title"`
}
