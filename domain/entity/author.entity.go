package entity

import (
	"database/sql"
	"time"
)

type Author struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Title     string
	Articles  []*Article `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
