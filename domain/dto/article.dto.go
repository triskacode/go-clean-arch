package dto

import (
	"time"
)

type CreateArticleDto struct {
	Title    string `json:"title" form:"title" validate:"required"`
	Content  string `json:"content" form:"content" validate:"required"`
	AuthorID int    `json:"author_id,string" form:"author_id" validate:"required,gt=0"`
}

type UpdateArticleDto struct {
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	AuthorID int    `json:"author_id,string" form:"author_id" validate:"omitempty,gt=0"`
}

type ArticleResponseDto struct {
	ID        uint               `json:"id,omitempty"`
	Title     string             `json:"title,omitempty"`
	Content   string             `json:"content,omitempty"`
	AuthorID  uint               `json:"author_id,omitempty"`
	Author    *AuthorResponseDto `json:"author,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}
