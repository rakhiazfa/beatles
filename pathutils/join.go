package pathutils

import (
	"regexp"
	"strings"
)

var slashCollapse = regexp.MustCompile(`/+`)

// Normalizes and concatenates two URL path segments (prefix and path).
func Join(prefix, path string) string {
	prefix = strings.TrimSuffix(prefix, "/")
	path = strings.TrimPrefix(path, "/")

	joined := prefix + "/" + path

	return joined
}
