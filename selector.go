package selector

import (
	"fmt"

	tty "github.com/mattn/go-tty"
)

func List(items []string) (int, error) {
	tty, err := tty.Open()
	if err != nil {
		return 0, err
	}

	defer tty.Close()

	out := tty.Output()

	// Hide the cursor
	out.Write([]byte("\x1b[?25l"))

	cursor := 0

	for {
		for i, item := range items {
			if cursor == i {
				out.Write([]byte("\x1b[34m" + " ❯ " + item + "\x1b[0m"))
			} else {
				out.Write([]byte("\x1b[2m" + "   " + item + "\x1b[0m"))
			}
			out.Write([]byte("\n"))
		}
		out.Write([]byte(fmt.Sprintf("\x1b[%dA", len(items))))

		r, err := tty.ReadRune()
		if err != nil {
			return 0, err
		}

		switch r {
		case 'j', 0x0E:
			if cursor < len(items)-1 {
				cursor++
			}
		case 'k', 0x10:
			if cursor > 0 {
				cursor--
			}
		case 13:
			// display the cursor
			out.Write([]byte("\x1b[0J\x1b[?25h"))
			return cursor, nil
		case 27:
			out.Write([]byte("\x1b[0J\x1b[?25h"))
			return 0, nil
		}
	}
}