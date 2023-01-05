package repository

import (
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

func (r *authorRepository) FindAll(authors *[]domain.Author) error {
	if result := r.conn.Find(authors); result.Error != nil {
		return result.Error
	}
	
	return nil
}

func (r *authorRepository) Create(author *domain.Author) error {
	if result := r.conn.Create(author); result.Error != nil {
		return result.Error
	}

	return nil
}
