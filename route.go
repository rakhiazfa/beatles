package beatles

type Route struct {
	handlers   []Handler
	parameters map[string]string
}
