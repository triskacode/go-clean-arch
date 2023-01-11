package exception

import "fmt"

type HttpException struct {
	Code    int
	Message string
	Detail  interface{}
}

func NewHttpException(code int, message string, err interface{}) *HttpException {
	return &HttpException{
		Code:    code,
		Message: message,
		Detail:  err,
	}
}

func (he HttpException) Error() string {
	return fmt.Sprintf("%s: %s", he.Message, he.Detail)
}
