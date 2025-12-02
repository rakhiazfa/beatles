package logger

import (
	"fmt"
	"strings"
)

type Format uint32

const (
	FormatText Format = iota
	FormatJSON
)

func (format Format) Marshal() ([]byte, error) {
	switch format {
	case FormatText:
		return []byte("text"), nil
	case FormatJSON:
		return []byte("json"), nil
	}

	return nil, fmt.Errorf("invalid log format: %s", format.String())
}

func (format *Format) Unmarshal(text []byte) (Format, error) {
	frmt := strings.ToLower(string(text))

	switch frmt {
	case "text":
		return FormatText, nil
	case "json":
		return FormatJSON, nil
	}

	var l Format
	return l, fmt.Errorf("invalid log format: %s", frmt)
}

func (format Format) String() string {
	if b, err := format.Marshal(); err == nil {
		return string(b)
	}

	return ""
}
