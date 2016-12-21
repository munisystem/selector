package main

import (
	"fmt"
	"log"
	"os"

	"github.com/munisystem/selector"
)

func main() {
	characters := []string{
		"Ruko Kominato",
		"Yuzuki Kurebayashi",
		"Hitoe Uemura",
		"Akira Aoi",
		"Iona Urazoe",
	}

	index, err := selector.List(characters, "Whom do you like?")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if index == -1 {
		fmt.Println("Not selected")
		os.Exit(0)
	}

	fmt.Println("I like " + characters[index] + "!!!")
}
