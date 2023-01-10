package repository

import (
	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
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

func (r *authorRepository) FindById(author *domain.Author) error {
	if result := r.conn.First(author); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *authorRepository) Update(author *domain.Author, f dto.UpdateAuthorDto) error {
	updateSet := domain.Author{
		Name:  *f.Name,
		Title: *f.Title,
	}

	switch result := r.conn.Model(author).Updates(updateSet); {
	case result.Error != nil:
		return result.Error
	case result.RowsAffected == 0:
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *authorRepository) Delete(author *domain.Author) error {
	switch result := r.conn.Delete(author); {
	case result.Error != nil:
		return result.Error
	case result.RowsAffected == 0:
		return gorm.ErrRecordNotFound
	}

	return nil
}
