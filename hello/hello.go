package main

import (
	"example/greetings"
	"fmt"
	"log"

	"golang.org/x/example/stringutil"
)

func main() {
	// Paramétrage des logs
	log.SetPrefix("Greetings module: ")
	log.SetFlags(0)

	// Exemple 1
	fmt.Println(stringutil.ToUpper("Hello world!"))

	// Exemple 2
	message, error := greetings.Hello("Mark")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)

	// Exemple 3
	names := []string{"Joshua", "Henry", "John", "Drew"}
	messages, errs := greetings.Hellos(names)
	if errs != nil {
		log.Fatal(errs)
	}
	fmt.Println(messages)
}
