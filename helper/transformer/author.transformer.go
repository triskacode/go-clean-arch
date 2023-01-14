package transformer

import (
	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/domain/entity"
)

type AuthorTransformer interface {
	ToSingleResponse(author *entity.Author) *dto.AuthorResponseDto
	ToSliceResponse(authors []*entity.Author) *[]dto.AuthorResponseDto
}

type authorTransformer struct {
	transformArticle *articleTransformer
}

var _authorTransformer *authorTransformer

func NewAuthorTransformer() (t *authorTransformer) {
	if _authorTransformer == nil {
		_authorTransformer = new(authorTransformer)
		_authorTransformer.transformArticle = NewArticleTransformer()
	}

	t = _authorTransformer
	return
}

func (t *authorTransformer) ToSingleResponse(author *entity.Author) *dto.AuthorResponseDto {
	if author == nil {
		return nil
	}

	return &dto.AuthorResponseDto{
		ID:        author.ID,
		Name:      author.Name,
		Title:     author.Title,
		Articles:  *t.transformArticle.ToSliceResponse(author.Articles),
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}

func (t *authorTransformer) ToSliceResponse(authors []*entity.Author) *[]dto.AuthorResponseDto {
	resp := make([]dto.AuthorResponseDto, 0)

	for _, author := range authors {
		resp = append(resp, *t.ToSingleResponse(author))
	}

	return &resp
}
