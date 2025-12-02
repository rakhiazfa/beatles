package beatles

type Request interface {
	GetHeader(key string) string

	Method() string

	Path() string

	PathVariable(key string) string
}
