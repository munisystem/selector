package main

import (
	"fmt"
	"log"

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
	}
	for index := range indexes {
		fmt.Println(titles[index])
	}
}
