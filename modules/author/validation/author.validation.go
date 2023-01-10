package validation

import (
	"github.com/triskacode/go-clean-arch/infrastructure/validator"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type AuthorValidator interface {
	ValidateCreateAuthorDto(f dto.CreateAuthorDto) (errModel validator.ValidationErrorModel)
	ValidateUpdateAuthorDto(f dto.UpdateAuthorDto) (errModel validator.ValidationErrorModel)
}

type authorValidator struct {
	validator validator.ValidatorAdapter
}

func NewAuthorValidator() (v *authorValidator) {
	v = new(authorValidator)
	v.validator = validator.NewCustomValidator()

	return
}

func (v *authorValidator) ValidateCreateAuthorDto(f dto.CreateAuthorDto) (errModel validator.ValidationErrorModel) {
	if err := v.validator.ValidateStruct(f); err != nil {
		errModel = v.validator.ParseErrors(err)
	}

	return
}

func (v *authorValidator) ValidateUpdateAuthorDto(f dto.UpdateAuthorDto) (errModel validator.ValidationErrorModel) {
	if err := v.validator.ValidateStruct(f); err != nil {
		errModel = v.validator.ParseErrors(err)
	}

	return
}
