package transformer

import (
	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/domain/entity"
)

type ArticleTransformer interface {
	ToSingleResponse(article *entity.Article) *dto.ArticleResponseDto
	ToSliceResponse(articles []*entity.Article) *[]dto.ArticleResponseDto
}

type articleTransformer struct {
	transformAuthor *authorTransformer
}

var _articleTransformer *articleTransformer

func NewArticleTransformer() (t *articleTransformer) {
	if _articleTransformer == nil {
		_articleTransformer = new(articleTransformer)
		_articleTransformer.transformAuthor = NewAuthorTransformer()
	}

	t = _articleTransformer
	return
}

func (t *articleTransformer) ToSingleResponse(article *entity.Article) *dto.ArticleResponseDto {
	if article == nil {
		return nil
	}

	return &dto.ArticleResponseDto{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		AuthorID:  article.AuthorID,
		Author:    t.transformAuthor.ToSingleResponse(article.Author),
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}

func (t *articleTransformer) ToSliceResponse(articles []*entity.Article) *[]dto.ArticleResponseDto {
	resp := make([]dto.ArticleResponseDto, 0)

	for _, article := range articles {
		resp = append(resp, *t.ToSingleResponse(article))
	}

	return &resp
}
