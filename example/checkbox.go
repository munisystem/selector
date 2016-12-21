package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/munisystem/selector"
)

func main() {
	titles := []string{
		"selector infected WIXOSS",
		"selector spread WIXOSS",
		"selector destructed WIXOSS",
		"Lostorage incited WIXOSS",
	}

	indexes, err := selector.Checkbox(titles, "What do you watch?")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if len(indexes) == 0 {
		os.Exit(0)
	}

	watched := []string{}
	for _, index := range indexes {
		watched = append(watched, titles[index])
	}
	fmt.Println("I watched " + strings.Join(watched, " and ") + "!!!")
}
