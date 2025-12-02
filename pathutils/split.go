package pathutils

import "strings"

func Split(path string) []string {
	path = strings.Trim(path, "/")

	if path == "" {
		return []string{}
	}

	return strings.Split(path, "/")
}
