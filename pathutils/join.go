package pathutils

import (
	"regexp"
	"strings"
)

var slashCollapse = regexp.MustCompile(`/+`)

func Join(prefix, path string) string {
	clearedPrefix := slashCollapse.ReplaceAllString(prefix, "/")
	clearedPrefix = strings.Trim(clearedPrefix, "/")

	clearedPath := slashCollapse.ReplaceAllString(path, "/")
	clearedPath = strings.Trim(clearedPath, "/")

	joined := "/" + clearedPrefix + "/" + clearedPath

	return slashCollapse.ReplaceAllString(joined, "/")
}
