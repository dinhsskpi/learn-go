package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"dinhpv", "cbong", "thuy"}

	messages, error := greetings.Hellos(names)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messages)
}
