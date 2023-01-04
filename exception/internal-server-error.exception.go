package exception

const (
	InternalServerErrorCode    = 500
	InternalServerErrorMessage = "INTERNAL_SERVER_ERROR"
)

func NewInternalServerErrorException(err interface{}) *HttpException {
	return NewHttpException(InternalServerErrorCode, InternalServerErrorMessage, err)
}
