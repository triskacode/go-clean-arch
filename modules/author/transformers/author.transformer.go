package transformers

import (
	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type AuthorTransformer interface {
	ToSingleResponse(author domain.Author) *dto.AuthorResponseDto
}

type authorTransformer struct{}

func NewAuthorTransformer() (t *authorTransformer) {
	t = new(authorTransformer)

	return
}

func (t authorTransformer) ToSingleResponse(author domain.Author) *dto.AuthorResponseDto {
	return &dto.AuthorResponseDto{
		ID:        author.ID,
		Name:      author.Name,
		Title:     author.Title,
		Articles:  author.Articles,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}
