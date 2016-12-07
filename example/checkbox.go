package main

import (
	"fmt"
	"log"
	"os"

	"github.com/munisystem/selector"
)

func main() {
	titles := []string{
		"selector infected WIXOSS",
		"selector spread WIXOSS",
		"selector destructed WIXOSS",
		"Lostorage incited WIXOSS",
	}

	fmt.Println("What do you watch?")
	indexes, err := selector.Checkbox(titles)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if len(indexes) == 0 {
		fmt.Println("Not selected")
		os.Exit(0)
	}

	for index := range indexes {
		fmt.Println(titles[index])
	}
}
