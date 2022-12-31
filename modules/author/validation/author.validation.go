package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type AuthorValidator interface {
	ValidateCreateAuthorDto(dto dto.CreateAuthorDto) (model []exception.ErrorValidationModel)
}

type authorValidator struct {
	validate *validator.Validate
}

func NewAuthorValidator() (v *authorValidator) {
	v = new(authorValidator)

	v.validate = validator.New()
	return
}

func parseValidationError(model *[]exception.ErrorValidationModel, err error) {
	for _, err := range err.(validator.ValidationErrors) {
		fi := new(exception.ErrorValidationModel)
		fi.Field = err.Field()
		fi.Tag = err.Tag()
		fi.Value = err.Param()
		*model = append(*model, *fi)
	}
}

func (v authorValidator) ValidateCreateAuthorDto(dto dto.CreateAuthorDto) (model []exception.ErrorValidationModel) {
	if err := v.validate.Struct(dto); err != nil {
		model = make([]exception.ErrorValidationModel, 0)
		parseValidationError(&model, err)
	}

	return
}
