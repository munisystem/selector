package selector

import (
	"fmt"
	"strings"

	tty "github.com/mattn/go-tty"
)

func Checkbox(items []string, caption string) ([]int, error) {
	tty, err := tty.Open()
	if err != nil {
		return nil, err
	}

	defer tty.Close()

	out := tty.Output()

	out.Write([]byte(hideCursor))
	out.Write([]byte(colorize(Green, "? ") + caption + "\n"))

	cursor := 0
	checkbox := make([]bool, len(items))
	for i := 0; i < len(checkbox); i++ {
		checkbox[i] = false
	}

	for {
		for i, item := range items {
			pointer := "  "
			if cursor == i {
				pointer = colorize(Cyan, "❯ ")
			}

			checkmark := colorize(Dim, "✔ ")
			if checkbox[i] == true {
				checkmark = colorize(Green, "✔ ")
			}

			out.Write([]byte(pointer + checkmark + colorize(Dim, item) + "\n"))
		}
		out.Write([]byte(fmt.Sprintf(upCursor, len(items))))

		r, err := tty.ReadRune()
		if err != nil {
			return nil, err
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
		case 0x20:
			if checkbox[cursor] == true {
				checkbox[cursor] = false
			} else {
				checkbox[cursor] = true
			}

		case 13:
			var result []int
			for i, v := range checkbox {
				if v == true {
					result = append(result, i)
				}
			}
			out.Write([]byte(fmt.Sprintf(upCursor, 1) + colorize(Green, "? ") + caption + " " + colorize(Blue, strChain(items, result)) + "\n"))
			out.Write([]byte(eraseDisplay + showCursor))
			return result, nil
		case 27:
			out.Write([]byte(fmt.Sprintf(upCursor, 1) + colorize(Green, "? ") + caption + " " + "\n"))
			out.Write([]byte(eraseDisplay + showCursor))
			return nil, nil
		}
	}
}

func strChain(items []string, indexes []int) string {
	tmp := []string{}
	for index := range indexes {
		tmp = append(tmp, items[index])
	}
	return strings.Join(tmp, ", ")
}
