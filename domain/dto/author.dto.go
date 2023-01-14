package dto

import (
	"time"
)

type CreateAuthorDto struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Title string `json:"title" form:"title" validate:"required"`
}

type ParamIdDto struct {
	ID uint `params:"id"`
}

type UpdateAuthorDto struct {
	Name  *string `json:"name" form:"name"`
	Title *string `json:"title" form:"title"`
}

type AuthorResponseDto struct {
	ID        uint                  `json:"id,omitempty"`
	Name      string                `json:"name,omitempty"`
	Title     string                `json:"title,omitempty"`
	Articles  *[]ArticleResponseDto `json:"articles,omitempty"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
	UpdatedAt time.Time             `json:"updated_at,omitempty"`
}
