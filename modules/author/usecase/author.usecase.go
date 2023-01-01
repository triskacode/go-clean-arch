package usecase

import (
	"fmt"

	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/modules/author/adapter"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
	"github.com/triskacode/go-clean-arch/modules/author/transformers"
)

type authorUsecase struct {
	authorTransformer transformers.AuthorTransformer
	authorRepository  adapter.AuthorRepository
}

func NewAuthorUsecase(authorRepository adapter.AuthorRepository) (u *authorUsecase) {
	u = new(authorUsecase)

	u.authorTransformer = transformers.NewAuthorTransformer()
	u.authorRepository = authorRepository
	return
}

func (u authorUsecase) Create(dto dto.CreateAuthorDto) dto.AuthorResponseDto {
	author := &domain.Author{
		Name:  dto.Name,
		Title: dto.Title,
	}

	if err := u.authorRepository.Create(author); err != nil {
		fmt.Println(err)
	}
	fmt.Println(author)

	return u.authorTransformer.ToSingleResponse(*author)
}
