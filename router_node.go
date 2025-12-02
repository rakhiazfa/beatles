package beatles

type routerNode struct {
	segment       string
	parameterName string
	isWildcard    bool
	children      []*routerNode
	handlers      []Handler
}
