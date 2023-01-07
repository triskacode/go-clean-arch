package exception

import "fmt"

type HttpException struct {
	Code    int32
	Message string
	Detail  interface{}
}

func NewHttpException(code int32, message string, err interface{}) *HttpException {
	return &HttpException{
		Code:    code,
		Message: message,
		Detail:  err,
	}
}

func (he *HttpException) Error() string {
	return fmt.Sprintf("%s: %s", he.Message, he.Detail)
}
