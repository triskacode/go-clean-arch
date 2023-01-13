package usecase

import (
	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/modules/article/adapter"
	"github.com/triskacode/go-clean-arch/modules/article/dto"
	"github.com/triskacode/go-clean-arch/modules/article/transformer"
)

type articleUsecase struct {
	articleTransformer transformer.ArticleTransformer
	articleRepository  adapter.ArticleRepository
}

func NewArticleUsecase(articleRepository adapter.ArticleRepository) (u *articleUsecase) {
	u = new(articleUsecase)
	u.articleTransformer = transformer.NewArticleTransformer()
	u.articleRepository = articleRepository

	return
}

func (u *articleUsecase) FindAll() (r *[]dto.ArticleResponseDto, e *exception.HttpException) {
	articles := make([]domain.Article, 0)

	if err := u.articleRepository.FindAll(&articles); err != nil {
		e = exception.NewInternalServerErrorException(err.Error())
		return
	}

	r = u.articleTransformer.ToSliceResponse(articles)
	return
}
