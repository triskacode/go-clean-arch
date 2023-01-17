package usecase

import (
	"errors"
	"fmt"

	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/domain/entity"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/helper/transformer"
	"github.com/triskacode/go-clean-arch/modules/article/adapter"
	author "github.com/triskacode/go-clean-arch/modules/author/adapter"
	"gorm.io/gorm"
)

type articleUsecase struct {
	transformArticle  transformer.ArticleTransformer
	articleRepository adapter.ArticleRepository
	authorRepository  author.AuthorRepository
}

func NewArticleUsecase(articleRepository adapter.ArticleRepository, authorRepository author.AuthorRepository) (u *articleUsecase) {
	u = new(articleUsecase)
	u.transformArticle = transformer.NewArticleTransformer()
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

	r = u.transformArticle.ToSliceResponse(articles)
	return
}

func (u *articleUsecase) Create(f dto.CreateArticleDto) (r *dto.ArticleResponseDto, e *exception.HttpException) {
	author := &entity.Author{
		ID: uint(f.AuthorID),
	}

	if err := u.authorRepository.FindOne(author); err != nil {
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

	r = u.transformArticle.ToSingleResponse(article)
	return
}

func (u *articleUsecase) FindById(p dto.ParamIdDto) (r *dto.ArticleResponseDto, e *exception.HttpException) {
	article := &entity.Article{
		ID: p.ID,
	}

	if err := u.articleRepository.FindOne(article); err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			e = exception.NewNotFoundException(err.Error())
			return
		default:
			e = exception.NewInternalServerErrorException(err.Error())
			return
		}
	}

	r = u.transformArticle.ToSingleResponse(article)
	return
}

func (u *articleUsecase) Update(p dto.ParamIdDto, f dto.UpdateArticleDto) (r *dto.ArticleResponseDto, e *exception.HttpException) {
	article := &entity.Article{
		ID: p.ID,
	}

	if f.AuthorID != 0 {
		author := &entity.Author{
			ID: uint(f.AuthorID),
		}

		if err := u.authorRepository.FindOne(author); err != nil {
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

		article.Author = author
	}

	if err := u.articleRepository.Update(article, f); err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			e = exception.NewNotFoundException(err.Error())
			return
		default:
			e = exception.NewInternalServerErrorException(err.Error())
			return
		}
	}

	r = u.transformArticle.ToSingleResponse(article)
	return
}
