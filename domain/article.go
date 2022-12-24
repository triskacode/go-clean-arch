package domain

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string
	Content  string
	AuthorID uint
}
