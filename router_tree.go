package beatles

import (
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

func (rt *routerTree) Register(method string, path string, handlers ...Handler) {
	method = strings.ToUpper(method)

	current := rt.root
	segments := pathutils.Split(path)

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
		handlers:   handlers,
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

	// Copy parameters
	for k, v := range parameters {
		route.parameters[k] = v
	}

	return route, nil
}

func (rt *routerTree) Use(handlers ...Handler) Router {
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

func (rt *routerTree) Get(path string, handlers ...Handler) Router {
	rt.Register(http.MethodGet, path, handlers...)
	return rt
}

func (rt *routerTree) Head(path string, handlers ...Handler) Router {
	rt.Register(http.MethodHead, path, handlers...)
	return rt
}

func (rt *routerTree) Post(path string, handlers ...Handler) Router {
	rt.Register(http.MethodPost, path, handlers...)
	return rt
}

func (rt *routerTree) Put(path string, handlers ...Handler) Router {
	rt.Register(http.MethodPut, path, handlers...)
	return rt
}

func (rt *routerTree) Delete(path string, handlers ...Handler) Router {
	rt.Register(http.MethodDelete, path, handlers...)
	return rt
}

func (rt *routerTree) Connect(path string, handlers ...Handler) Router {
	rt.Register(http.MethodConnect, path, handlers...)
	return rt
}

func (rt *routerTree) Options(path string, handlers ...Handler) Router {
	rt.Register(http.MethodOptions, path, handlers...)
	return rt
}

func (rt *routerTree) Trace(path string, handlers ...Handler) Router {
	rt.Register(http.MethodTrace, path, handlers...)
	return rt
}

func (rt *routerTree) Patch(path string, handlers ...Handler) Router {
	rt.Register(http.MethodPatch, path, handlers...)
	return rt
}

func (rt *routerTree) All(path string, handlers ...Handler) Router {
	for _, method := range allMethods {
		rt.Register(method, path, handlers...)
	}

	return rt
}
