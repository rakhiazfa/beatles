package beatles

type Request interface {
	// Returns the value of a header by its key
	GetHeader(key string) string

	// Returns the HTTP method of the request
	Method() string

	// Returns the request path
	Path() string

	// Returns the value of a path parameter by key
	PathVariable(key string) string
}
