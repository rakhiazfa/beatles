package beatles

import (
	"net/http"
	"strings"

	"github.com/rakhiazfa/beatles/pathutils"
)

type routerTree struct {
	trees    map[string]*routerNode
	handlers []Handler
}

func NewRouterTree() RouterTree {
	return &routerTree{
		trees:    make(map[string]*routerNode),
		handlers: []Handler{},
	}
}

func (rt *routerTree) Register(method string, path string, handlers ...Handler) {
	segments := pathutils.Split(path)

	root, ok := rt.trees[method]
	if !ok {
		root = &routerNode{}
		rt.trees[method] = root
	}

	currentNode := root

	for _, segment := range segments {
		var child *routerNode
		var isParameter bool
		var isWildcard bool
		var parameterName string

		if strings.HasPrefix(segment, ":") {
			isParameter = true
			parameterName = segment[1:]
		} else if strings.HasPrefix(segment, "*") {
			isWildcard = true
			parameterName = segment[1:]
		}

		for _, node := range currentNode.children {
			if node.segment == segment || (isParameter && node.parameterName != "") || (isWildcard && node.isWildcard) {
				child = node
				break
			}
		}

		if child == nil {
			child = &routerNode{
				segment:       segment,
				parameterName: parameterName,
				isWildcard:    isWildcard,
			}
			currentNode.children = append(currentNode.children, child)
		}

		currentNode = child

		if isWildcard {
			break
		}
	}

	currentNode.handlers = handlers
}

func (rt *routerTree) Search(method string, path string) *Route {
	root := rt.trees[method]
	if root == nil {
		return nil
	}

	segments := pathutils.Split(path)
	parameters := make(map[string]string)

	currentNode := root

	for index := range len(segments) {
		segment := segments[index]

		var next *routerNode

		for _, node := range currentNode.children {
			if node.segment == segment && node.parameterName == "" && !node.isWildcard {
				next = node
				break
			}
		}

		if next == nil {
			for _, node := range currentNode.children {
				if node.parameterName != "" && !node.isWildcard {
					next = node
					parameters[node.parameterName] = segment
					break
				}
			}
		}

		if next == nil {
			for _, node := range currentNode.children {
				if node.isWildcard {
					next = node
					parameters[node.parameterName] = strings.Join(segments[index:], "/")
					return &Route{
						handlers:   next.handlers,
						parameters: parameters,
					}
				}
			}
		}

		if next == nil {
			return nil
		}

		currentNode = next
	}

	return &Route{
		handlers:   currentNode.handlers,
		parameters: parameters,
	}
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
