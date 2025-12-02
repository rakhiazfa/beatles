package beatles

type Response interface {
	SetHeader(key string, value string) Response

	SetStatus(statusCode int) Response

	SetBody(body []byte) error

	JSON(v JSON) error
}
