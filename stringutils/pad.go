package stringutils

import (
	"regexp"
	"strings"
)

// Regex for matching ANSI escape sequences used to color terminal output.
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// Returns the length of a string excluding ANSI color codes.
func visibleLength(s string) int {
	return len(ansiRegex.ReplaceAllString(s, ""))
}

// Appends spaces to the right of the string until it reaches the
func PadRight(s string, width int) string {
	visibleLen := visibleLength(s)

	if visibleLen >= width {
		return s
	}

	return s + strings.Repeat(" ", width-visibleLen)
}
