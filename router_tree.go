package beatles

import (
	"maps"
	"net/http"
	"strings"

	"github.com/rakhiazfa/beatles/pathutils"
)

type routerTree struct {
	root     *routerNode
	handlers []Handler
}

func NewRouterTree() RouterTree {
	return &routerTree{
		root: newRouterNode("/", SegmentStatic, ""),
	}
}

func (rt *routerTree) Register(method string, path string, handler Handler, handlers ...Handler) {
	method = strings.ToUpper(method)

	current := rt.root
	segments := pathutils.Split(path)

	allHandlers := append([]Handler{handler}, handlers...)

	for _, segment := range segments {
		segmentType, parameterName := classifySegment(segment)

		child := current.findChildNodeBySegmentAndSegmentType(segment, segmentType)

		if child == nil {
			child = newRouterNode(segment, segmentType, parameterName)
			current.children = append(current.children, child)
		}

		current = child
	}

	current.routes[method] = &route{
		handlers:   allHandlers,
		parameters: make(map[string]string),
	}
}

func (rt *routerTree) Search(method string, path string) (*route, error) {
	method = strings.ToUpper(method)

	current := rt.root
	segments := pathutils.Split(path)

	parameters := make(map[string]string)

	for _, segment := range segments {
		var match *routerNode

		// Try to get router node with segment type 'static'
		for _, child := range current.children {
			if child.segmentType == SegmentStatic && child.segment == segment {
				match = child

				break
			}
		}

		// Try to get router node with segment type 'parameter'
		if match == nil {
			for _, child := range current.children {
				if child.segmentType == SegmentParam {
					match = child
					parameters[child.parameterName] = segment

					break
				}
			}
		}

		// Try to get router node with segment type 'wildcard'
		if match == nil {
			for _, child := range current.children {
				if child.segmentType == SegmentWildcard {
					match = child
					parameters[child.parameterName] = segment

					break
				}
			}
		}

		if match == nil {
			return nil, NewHTTPError(http.StatusNotFound, "Route not found")
		}

		current = match
	}

	route, ok := current.routes[method]
	if !ok {
		return nil, NewHTTPError(http.StatusMethodNotAllowed, "Method not allowed")
	}

	maps.Copy(route.parameters, parameters)

	return route, nil
}

func (rt *routerTree) Use(handler Handler, handlers ...Handler) Router {
	rt.handlers = append(rt.handlers, handler)
	rt.handlers = append(rt.handlers, handlers...)

	return rt
}

func (rt *routerTree) Group(path string, handlers ...Handler) Router {
	return &router{
		routerTree: rt,
		prefix:     pathutils.Join("/", path),
		handlers:   mergeHandlers(rt.handlers, handlers),
	}
}

func (rt *routerTree) Get(path string, handler Handler, handlers ...Handler) Router {
	rt.Register(http.MethodGet, path, handler, handlers...)
	return rt
}

func (rt *routerTree) Head(path string, handler Handler, handlers ...Handler) Router {
	rt.Register(http.MethodHead, path, handler, handlers...)
	return rt
}

func (rt *routerTree) Post(path string, handler Handler, handlers ...Handler) Router {
	rt.Register(http.MethodPost, path, handler, handlers...)
	return rt
}

func (rt *routerTree) Put(path string, handler Handler, handlers ...Handler) Router {
	rt.Register(http.MethodPut, path, handler, handlers...)
	return rt
}

func (rt *routerTree) Delete(path string, handler Handler, handlers ...Handler) Router {
	rt.Register(http.MethodDelete, path, handler, handlers...)
	return rt
}

func (rt *routerTree) Connect(path string, handler Handler, handlers ...Handler) Router {
	rt.Register(http.MethodConnect, path, handler, handlers...)
	return rt
}

func (rt *routerTree) Options(path string, handler Handler, handlers ...Handler) Router {
	rt.Register(http.MethodOptions, path, handler, handlers...)
	return rt
}

func (rt *routerTree) Trace(path string, handler Handler, handlers ...Handler) Router {
	rt.Register(http.MethodTrace, path, handler, handlers...)
	return rt
}

func (rt *routerTree) Patch(path string, handler Handler, handlers ...Handler) Router {
	rt.Register(http.MethodPatch, path, handler, handlers...)
	return rt
}

func (rt *routerTree) All(path string, handler Handler, handlers ...Handler) Router {
	for _, method := range allMethods {
		rt.Register(method, path, handler, handlers...)
	}

	return rt
}
