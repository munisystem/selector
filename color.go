package selector

import "fmt"

const Dim = 2

const (
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func colorize(color int, str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, str)
}
