package selector

const (
	hideCursor   = "\x1b[?25l"
	showCursor   = "\x1b[?25h"
	eraseDisplay = "\x1b[0J"
	upCursor     = "\x1b[%dA"
)
