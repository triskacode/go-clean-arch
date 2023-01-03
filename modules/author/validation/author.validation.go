package validation

import (
	"github.com/triskacode/go-clean-arch/infrastructure/validator"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type AuthorValidator interface {
	ValidateCreateAuthorDto(dto dto.CreateAuthorDto) (model validator.ValidationErrorModel)
}

type authorValidator struct {
	validator validator.ValidatorAdapter
}

func NewAuthorValidator() (v *authorValidator) {
	v = new(authorValidator)

	v.validator = validator.NewCustomValidator()
	return
}

func (v authorValidator) ValidateCreateAuthorDto(dto dto.CreateAuthorDto) (errModel validator.ValidationErrorModel) {
	if err := v.validator.ValidateStruct(dto); err != nil {
		errModel = v.validator.ParseErrors(err)
	}

	return
}
