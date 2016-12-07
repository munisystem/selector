package main

import (
	"fmt"
	"log"

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
	}
	fmt.Println(characters[index])
}
