package domain

import (
	"database/sql"
	"time"
)

type Article struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Content   string
	AuthorID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
