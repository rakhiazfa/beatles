package beatles

type RouterTree interface {
	Router

	// Adds a route to the router with the given HTTP method, path, and one or more handlers
	Register(method string, path string, handler Handler, handlers ...Handler)

	// Search looks up a route based on the HTTP method and path
	Search(method string, path string) (*route, error)
}
