package dto

type CreateAuthorDto struct {
	Name  string `validate:"required" json:"name"`
	Title string `validate:"required" json:"title"`
}
