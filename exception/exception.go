package exception

import "fmt"

type HttpException struct {
	Code    int32
	Message string
	Errors  interface{}
}

func NewHttpException(code int32, message string, err interface{}) *HttpException {
	return &HttpException{
		Code:    code,
		Message: message,
		Errors:  err,
	}
}

func (he *HttpException) Error() string {
	return fmt.Sprintf("%s: %s", he.Message, he.Errors)
}
