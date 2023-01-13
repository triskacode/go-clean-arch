package repository

import (
	"github.com/triskacode/go-clean-arch/domain"
	"gorm.io/gorm"
)

type articleRepository struct {
	conn *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) (r *articleRepository) {
	r = new(articleRepository)
	r.conn = conn

	return
}

func (r *articleRepository) FindAll(articles *[]domain.Article) error {
	if result := r.conn.Find(articles); result.Error != nil {
		return result.Error
	}

	return nil
}
