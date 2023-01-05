package usecase

import (
	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/exception"
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

func (u *authorUsecase) FindAll() (*[]dto.AuthorResponseDto, *exception.HttpException) {
	authors := make([]domain.Author, 0)

	if err := u.authorRepository.FindAll(&authors); err != nil {
		res := make([]dto.AuthorResponseDto, 0)
		return &res, exception.NewInternalServerErrorException(err.Error())
	}

	return u.authorTransformer.ToSliceResponse(authors), nil
}

func (u *authorUsecase) Create(f dto.CreateAuthorDto) (*dto.AuthorResponseDto, *exception.HttpException) {
	author := &domain.Author{
		Name:  f.Name,
		Title: f.Title,
	}

	if err := u.authorRepository.Create(author); err != nil {
		res := new(dto.AuthorResponseDto)
		return res, exception.NewInternalServerErrorException(err.Error())
	}

	return u.authorTransformer.ToSingleResponse(*author), nil
}
