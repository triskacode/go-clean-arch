package entity

import (
	"database/sql"
	"time"
)

type Article struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Content   string
	AuthorID  uint
	Author    *Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
