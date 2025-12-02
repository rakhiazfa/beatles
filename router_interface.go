package beatles

type Router interface {
	Use(handler Handler, handlers ...Handler) Router
	Group(path string, handlers ...Handler) Router

	Get(path string, handler Handler, handlers ...Handler) Router
	Head(path string, handler Handler, handlers ...Handler) Router
	Post(path string, handler Handler, handlers ...Handler) Router
	Put(path string, handler Handler, handlers ...Handler) Router
	Delete(path string, handler Handler, handlers ...Handler) Router
	Connect(path string, handler Handler, handlers ...Handler) Router
	Options(path string, handler Handler, handlers ...Handler) Router
	Trace(path string, handler Handler, handlers ...Handler) Router
	Patch(path string, handler Handler, handlers ...Handler) Router

	All(path string, handler Handler, handlers ...Handler) Router
}
