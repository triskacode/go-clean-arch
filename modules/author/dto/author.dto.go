package dto

type CreateAuthorDto struct {
	Name  string `validate:"required"`
	Title string `validate:"required"`
}
