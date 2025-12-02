package beatles

type route struct {
	handlers   []Handler
	parameters map[string]string
}
