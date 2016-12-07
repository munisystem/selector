package selector

import (
	"fmt"

	tty "github.com/mattn/go-tty"
)

func List(items []string) (int, error) {
	tty, err := tty.Open()
	if err != nil {
		return -1, err
	}

	defer tty.Close()

	out := tty.Output()

	out.Write([]byte(hideCursor))

	cursor := 0

	for {
		for i, item := range items {
			str := colorize(Dim, "   "+item)
			if cursor == i {
				str = colorize(Blue, " ❯ "+item)
			}
			out.Write([]byte(str + "\n"))
		}
		out.Write([]byte(fmt.Sprintf(upCursor, len(items))))

		r, err := tty.ReadRune()
		if err != nil {
			return -1, err
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
			out.Write([]byte(eraseDisplay + showCursor))
			return cursor, nil
		case 27:
			out.Write([]byte(eraseDisplay + showCursor))
			return -1, nil
		}
	}
}
