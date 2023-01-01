package repository

import (
	"fmt"

	"github.com/triskacode/go-clean-arch/domain"
	"gorm.io/gorm"
)

type authorRepository struct {
	conn *gorm.DB
}

func NewAuthorRepository(conn *gorm.DB) (r *authorRepository) {
	r = new(authorRepository)

	r.conn = conn
	return
}

func (r authorRepository) Create(author *domain.Author) error {
	result := r.conn.Create(author)
	fmt.Println(result)
	return nil
}
