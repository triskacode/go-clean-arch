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

func (u *authorUsecase) FindAll() (r *[]dto.AuthorResponseDto, e *exception.HttpException) {
	authors := make([]domain.Author, 0)

	if err := u.authorRepository.FindAll(&authors); err != nil {
		e = exception.NewInternalServerErrorException(err.Error())
		return
	}

	r = u.authorTransformer.ToSliceResponse(authors)
	return
}

func (u *authorUsecase) Create(f dto.CreateAuthorDto) (r *dto.AuthorResponseDto, e *exception.HttpException) {
	author := domain.Author{
		Name:  f.Name,
		Title: f.Title,
	}

	if err := u.authorRepository.Create(&author); err != nil {
		e = exception.NewInternalServerErrorException(err.Error())
		return
	}

	r = u.authorTransformer.ToSingleResponse(author)
	return
}
