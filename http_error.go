package beatles

import "fmt"

type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewHTTPError(statusCode int, message string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("statusCode: %d message: %s", e.StatusCode, e.Message)
}
