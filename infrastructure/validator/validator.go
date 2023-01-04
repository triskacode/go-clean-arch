package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type customValidator struct {
	validation *validator.Validate
}

func NewCustomValidator() (v *customValidator) {
	v = new(customValidator)
	v.validation = validator.New()
	v.validation.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return
}

func (v *customValidator) ValidateStruct(s interface{}) error {
	return v.validation.Struct(s)
}

func (v *customValidator) ParseErrors(err error) (model ValidationErrorModel) {
	model = make(ValidationErrorModel)
	for _, err := range err.(validator.ValidationErrors) {
		if _, ok := model[err.Field()]; !ok {
			model[err.Field()] = getMessageForError(err)
		}
	}

	return
}

func getMessageForError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("field %s is required", fe.Field())
	case "email":
		return fmt.Sprintf("field %s must be an email", fe.Field())
	}

	return fmt.Sprintf("field %s must implement rule %s", fe.Field(), fe.Tag())
}
