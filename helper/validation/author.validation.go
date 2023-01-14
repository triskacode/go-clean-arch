package validation

import (
	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/infrastructure/validator"
)

type AuthorValidator interface {
	ValidateCreateAuthorDto(f dto.CreateAuthorDto) (errModel validator.ValidationErrorModel)
	ValidateUpdateAuthorDto(f dto.UpdateAuthorDto) (errModel validator.ValidationErrorModel)
}

type authorValidator struct {
	validator validator.ValidatorAdapter
}

var _authorValidator *authorValidator

func NewAuthorValidator() (v *authorValidator) {
	if _authorValidator == nil {
		_authorValidator = new(authorValidator)
		_authorValidator.validator = validator.NewCustomValidator()
	}

	v = _authorValidator
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
