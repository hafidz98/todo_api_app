package web

type ActivityGroupsUpdateRequest struct {
	ID    int    `json:"id"`
	Title string `validate:"required" json:"title"`
}
