package beatles

type routerNode struct {
	segment       string
	segmentType   SegmentType
	children      []*routerNode
	parameterName string
	routes        map[string]*route
}

func newRouterNode(
	segment string,
	segmentType SegmentType,
	parameterName string,
) *routerNode {
	return &routerNode{
		segment:       segment,
		segmentType:   segmentType,
		children:      []*routerNode{},
		parameterName: parameterName,
		routes:        make(map[string]*route),
	}
}

func (rn *routerNode) findChildNodeBySegmentAndSegmentType(segment string, segmentType SegmentType) *routerNode {
	for _, child := range rn.children {
		if child.segmentType == segmentType {
			if segmentType == SegmentParam || segmentType == SegmentWildcard {
				return child
			}

			if child.segment == segment {
				return child
			}
		}
	}

	return nil
}
