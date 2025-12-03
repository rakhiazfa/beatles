package beatles

type Response interface {
	// Sets or updates a response header by key and value
	SetHeader(key string, value string) Response

	// Sets HTTP status code for the response
	SetStatus(statusCode int) Response

	// Sets the raw response body as a byte slice
	SetBody(body []byte) error

	// Serializes the given value into JSON and writes it to the response body
	JSON(v JSON) error
}
