package dto

type CreateAuthorDto struct {
	Name  string `validate:"required" json:"name" form:"name"`
	Title string `validate:"required" json:"title" form:"title"`
}
