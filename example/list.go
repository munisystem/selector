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

	fmt.Println("Whom do you like?")
	index, err := selector.List(characters)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if index == -1 {
		fmt.Println("Not selected")
		os.Exit(0)
	}

	fmt.Println(characters[index])
}
