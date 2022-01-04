package web

type ActivityGroupsCreateRequest struct {
	Email string `json:"email"`
	Title string `validate:"required" json:"title"`
}
