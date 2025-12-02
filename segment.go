package beatles

import (
	"fmt"
	"strings"
)

type SegmentType uint32

const (
	SegmentStatic SegmentType = iota
	SegmentParam
	SegmentWildcard
)

func (segmentType SegmentType) Marshal() ([]byte, error) {
	switch segmentType {
	case SegmentStatic:
		return []byte("static"), nil
	case SegmentParam:
		return []byte("param"), nil
	case SegmentWildcard:
		return []byte("wildcard"), nil
	}

	return nil, fmt.Errorf("invalid log segment type: %s", segmentType.String())
}

func (segmentType *SegmentType) Unmarshal(text []byte) (SegmentType, error) {
	frmt := strings.ToLower(string(text))

	switch frmt {
	case "static":
		return SegmentStatic, nil
	case "param":
		return SegmentParam, nil
	case "wildcard":
		return SegmentWildcard, nil
	}

	var l SegmentType
	return l, fmt.Errorf("invalid log segment type: %s", frmt)
}

func (segmentType SegmentType) String() string {
	if b, err := segmentType.Marshal(); err == nil {
		return string(b)
	}

	return ""
}

func classifySegment(segment string) (SegmentType, string) {
	if strings.HasPrefix(segment, ":") {
		return SegmentParam, segment[1:]
	}

	if strings.HasPrefix(segment, "*") {
		name := ""

		if len(segment) > 1 {
			name = segment[1:]
		}

		return SegmentWildcard, name
	}

	return SegmentStatic, ""
}
