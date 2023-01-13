package transformer

import (
	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/modules/article/dto"
)

type ArticleTransformer interface {
	ToSingleResponse(article domain.Article) *dto.ArticleResponseDto
	ToSliceResponse(articles []domain.Article) *[]dto.ArticleResponseDto
}

type articleTransformer struct{}

func NewArticleTransformer() (t *articleTransformer) {
	t = new(articleTransformer)

	return
}

func (t *articleTransformer) ToSingleResponse(article domain.Article) *dto.ArticleResponseDto {
	return &dto.ArticleResponseDto{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		AuthorID:  article.AuthorID,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}

func (t * articleTransformer) ToSliceResponse(articles []domain.Article) *[]dto.ArticleResponseDto {
	resp := make([]dto.ArticleResponseDto, 0)

	for _, article := range articles {
		resp = append(resp, *t.ToSingleResponse(article))
	}

	return &resp
}