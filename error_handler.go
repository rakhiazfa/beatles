package beatles

import (
	"errors"
	"net/http"
)

type ErrorHandler = func(c Context, err error) error

func DefaultErrorHandler(c Context, err error) error {
	statusCode := http.StatusInternalServerError
	message := err.Error()

	var httpError *HTTPError

	if errors.As(err, &httpError) {
		statusCode = httpError.StatusCode
		message = httpError.Message
	}

	if c.Request().GetHeader(HeaderAccept) == ContentTypeJSON {
		return c.Response().SetStatus(statusCode).JSON(JSON{
			"message": message,
		})
	}

	return c.Response().SetStatus(statusCode).SetBody([]byte(message))
}
