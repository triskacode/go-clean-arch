package domain

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name     string
	Title    string
	Articles []Article `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
