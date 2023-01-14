package validation

import (
	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/infrastructure/validator"
)

type ArticleValidator interface {
	ValidateCreateArticleDto(f dto.CreateArticleDto) (errModel validator.ValidationErrorModel)
}

type articleValidator struct {
	validator validator.ValidatorAdapter
}

var _articleValidator *articleValidator

func NewArticleValidator() (v *articleValidator) {
	if _articleValidator == nil {
		_articleValidator = new(articleValidator)
		_articleValidator.validator = validator.NewCustomValidator()
	}

	v = _articleValidator
	return
}

func (v *articleValidator) ValidateCreateArticleDto(f dto.CreateArticleDto) (errModel validator.ValidationErrorModel) {
	if err := v.validator.ValidateStruct(f); err != nil {
		errModel = v.validator.ParseErrors(err)
	}

	return
}
