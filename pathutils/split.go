package pathutils

import "strings"

// Split cleans the given path by trimming leading and trailing slashes, then splits it into its individual components
func Split(path string) []string {
	path = strings.Trim(path, "/")

	if path == "" {
		return []string{}
	}

	return strings.Split(path, "/")
}
