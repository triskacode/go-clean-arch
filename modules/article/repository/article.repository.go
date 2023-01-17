package repository

import (
	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type articleRepository struct {
	conn *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) (r *articleRepository) {
	r = new(articleRepository)
	r.conn = conn

	return
}

func (r *articleRepository) FindAll(articles *[]*entity.Article) error {
	if result := r.conn.Find(articles); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *articleRepository) Create(article *entity.Article) error {
	q := r.conn.Model(&entity.Article{})
	q = q.Omit(clause.Associations)
	if result := q.Create(article); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *articleRepository) FindOne(article *entity.Article) error {
	if result := r.conn.First(article); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *articleRepository) Update(article *entity.Article, f dto.UpdateArticleDto) error {
	updateSet := entity.Article{
		Title:    f.Title,
		Content:  f.Content,
		AuthorID: uint(f.AuthorID),
	}

	q := r.conn.Model(article)
	q = q.Clauses(clause.Returning{})
	switch result := q.Updates(updateSet); {
	case result.Error != nil:
		return result.Error
	case result.RowsAffected == 0:
		return gorm.ErrRecordNotFound
	}

	return nil
}
