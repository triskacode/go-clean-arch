package usecase

import (
	"errors"
	"fmt"

	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/domain/entity"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/helper/transformer"
	"github.com/triskacode/go-clean-arch/modules/article/adapter"
	authorAdapter "github.com/triskacode/go-clean-arch/modules/author/adapter"
	"gorm.io/gorm"
)

type articleUsecase struct {
	articleTransformer transformer.ArticleTransformer
	articleRepository  adapter.ArticleRepository
	authorRepository   authorAdapter.AuthorRepository
}

func NewArticleUsecase(articleRepository adapter.ArticleRepository, authorRepository authorAdapter.AuthorRepository) (u *articleUsecase) {
	u = new(articleUsecase)
	u.articleTransformer = transformer.NewArticleTransformer()
	u.articleRepository = articleRepository
	u.authorRepository = authorRepository

	return
}

func (u *articleUsecase) FindAll() (r *[]dto.ArticleResponseDto, e *exception.HttpException) {
	articles := make([]*entity.Article, 0)

	if err := u.articleRepository.FindAll(&articles); err != nil {
		e = exception.NewInternalServerErrorException(err.Error())
		return
	}

	r = u.articleTransformer.ToSliceResponse(articles)
	return
}

func (u *articleUsecase) Create(f dto.CreateArticleDto) (r *dto.ArticleResponseDto, e *exception.HttpException) {
	author := &entity.Author{
		ID: uint(f.AuthorID),
	}

	if err := u.authorRepository.FindById(author); err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			message := fmt.Sprintf("author with id: %d not found", f.AuthorID)
			e = exception.NewBadRequestException(message)
			return
		default:	
			e = exception.NewInternalServerErrorException(err.Error())
			return
		}
	}

	article := &entity.Article{
		Title:    f.Title,
		Content:  f.Content,
		AuthorID: uint(f.AuthorID),
		Author:   author,
	}

	if err := u.articleRepository.Create(article); err != nil {
		e = exception.NewInternalServerErrorException(err.Error())
		return
	}

	r = u.articleTransformer.ToSingleResponse(article)
	return
}
