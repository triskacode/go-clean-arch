package dto

import (
	"time"

	"github.com/triskacode/go-clean-arch/domain"
)

type CreateAuthorDto struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Title string `json:"title" form:"title" validate:"required"`
}

type ParamIdDto struct {
	ID uint `params:"id"`
}

type AuthorResponseDto struct {
	ID        uint             `json:"id"`
	Name      string           `json:"name"`
	Title     string           `json:"title"`
	Articles  []domain.Article `json:"articles,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
