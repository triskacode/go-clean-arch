package usecase

import (
	"errors"

	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/modules/author/adapter"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
	"github.com/triskacode/go-clean-arch/modules/author/transformers"
	"gorm.io/gorm"
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

func (u *authorUsecase) FindById(p dto.ParamIdDto) (r *dto.AuthorResponseDto, e *exception.HttpException) {
	author := domain.Author{
		ID: p.ID,
	}

	if err := u.authorRepository.FindById(&author); err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			e = exception.NewNotFoundException(err.Error())
			return
		default:
			e = exception.NewInternalServerErrorException(err.Error())
			return
		}
	}

	r = u.authorTransformer.ToSingleResponse(author)
	return
}

func (u *authorUsecase) Update(p dto.ParamIdDto, f dto.UpdateAuthorDto) (r *dto.AuthorResponseDto, e *exception.HttpException) {
	author := domain.Author{
		ID: p.ID,
	}

	if err := u.authorRepository.FindById(&author); err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			e = exception.NewNotFoundException(err.Error())
			return
		default:
			e = exception.NewInternalServerErrorException(err.Error())
			return
		}
	}

	if err := u.authorRepository.Update(&author, f); err != nil {
		e = exception.NewInternalServerErrorException(err.Error())
		return
	}

	r = u.authorTransformer.ToSingleResponse(author)
	return
}

func (u *authorUsecase) Delete(p dto.ParamIdDto) (r *dto.AuthorResponseDto, e *exception.HttpException) {
	author := domain.Author{
		ID: p.ID,
	}

	if err := u.authorRepository.FindById(&author); err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			e = exception.NewNotFoundException(err.Error())
			return
		default:
			e = exception.NewInternalServerErrorException(err.Error())
			return
		}
	}

	if err := u.authorRepository.Delete(&author); err != nil {
		e = exception.NewInternalServerErrorException(err.Error())
		return
	}

	r = u.authorTransformer.ToSingleResponse(author)
	return
}
