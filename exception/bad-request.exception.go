package exception

const (
	BadRequestCode    = 400
	BadRequestMessage = "BAD_REQUEST"
)

func NewBadRequestException(err interface{}) *HttpException {
	return NewHttpException(BadRequestCode, BadRequestMessage, err)
}
