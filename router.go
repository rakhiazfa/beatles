package beatles

import (
	"net/http"

	"github.com/rakhiazfa/beatles/pathutils"
)

var allMethods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
	http.MethodPatch,
}

type router struct {
	routerTree RouterTree
	prefix     string
	handlers   []Handler
}

func (r *router) Use(handlers ...Handler) Router {
	r.handlers = append(r.handlers, handlers...)
	return r
}

func (r *router) Group(path string, handlers ...Handler) Router {
	return &router{
		routerTree: r.routerTree,
		prefix:     pathutils.Join(r.prefix, path),
		handlers:   mergeHandlers(r.handlers, handlers),
	}
}

func (r *router) Get(path string, handlers ...Handler) Router {
	r.register(http.MethodGet, path, handlers...)
	return r
}

func (r *router) Head(path string, handlers ...Handler) Router {
	r.register(http.MethodHead, path, handlers...)
	return r
}

func (r *router) Post(path string, handlers ...Handler) Router {
	r.register(http.MethodPost, path, handlers...)
	return r
}

func (r *router) Put(path string, handlers ...Handler) Router {
	r.register(http.MethodPut, path, handlers...)
	return r
}

func (r *router) Delete(path string, handlers ...Handler) Router {
	r.register(http.MethodDelete, path, handlers...)
	return r
}

func (r *router) Connect(path string, handlers ...Handler) Router {
	r.register(http.MethodConnect, path, handlers...)
	return r
}

func (r *router) Options(path string, handlers ...Handler) Router {
	r.register(http.MethodOptions, path, handlers...)
	return r
}

func (r *router) Trace(path string, handlers ...Handler) Router {
	r.register(http.MethodTrace, path, handlers...)
	return r
}

func (r *router) Patch(path string, handlers ...Handler) Router {
	r.register(http.MethodPatch, path, handlers...)
	return r
}

func (r *router) All(path string, handlers ...Handler) Router {
	for _, method := range allMethods {
		r.register(method, path, handlers...)
	}

	return r
}

func (r *router) register(method string, path string, handlers ...Handler) {
	r.routerTree.Register(method, pathutils.Join(r.prefix, path), mergeHandlers(r.handlers, handlers)...)
}
