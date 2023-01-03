package validator

type ValidatorAdapter interface {
	ValidateStruct(s interface{}) error
	ParseErrors(err error) (model ValidationErrorModel)
}

type ValidationErrorModel map[string]string
