package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type authorValidator struct {
	validate *validator.Validate
}

func NewAuthorValidator() (v *authorValidator) {
	v = new(authorValidator)

	v.validate = validator.New()
	return
}

func parseValidationError(obj []*exception.ErrorValidation, err error) {
	for _, err := range err.(validator.ValidationErrors) {
		fi := new(exception.ErrorValidation)
		fi.Field = err.Field()
		fi.Tag = err.Tag()
		fi.Value = err.Param()
		obj = append(obj, fi)
	}
}

func (v authorValidator) ValidateCreateAuthorDto(dto dto.CreateAuthorDto) []*exception.ErrorValidation {
	errors := make([]*exception.ErrorValidation, 0)
	if err := v.validate.Struct(dto); err != nil {
		parseValidationError(errors, err)
	}

	return errors
}
