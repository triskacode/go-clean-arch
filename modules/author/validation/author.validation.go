package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateCreateAuthorDto(dto dto.CreateAuthorDto) []*ErrorResponse {
	errors := make([]*ErrorResponse, 0)
	if err := validate.Struct(dto); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fi := new(ErrorResponse)
			fi.FailedField = err.Field()
			fi.Tag = err.Tag()
			fi.Value = err.Param()
			errors = append(errors, fi)
		}
	}

	return errors
}
