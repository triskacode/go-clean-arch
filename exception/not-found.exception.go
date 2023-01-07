package exception

const (
	NotFoundCode    = 404
	NotFoundMessage = "NOT_FOUND"
)

func NewNotFoundException(err interface{}) *HttpException {
	return NewHttpException(NotFoundCode, NotFoundMessage, err)
}
