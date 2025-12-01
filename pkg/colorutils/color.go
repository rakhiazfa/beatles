package colorutils

import "fmt"

const (
	Reset = "\x1b[0m"
)

func RGB(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func Colorize(str, color string) string {
	return color + str + Reset
}
