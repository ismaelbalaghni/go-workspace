package main

import (
	"bufio"
	"example/greetings"
	"example/mascot"
	"fmt"
	"log"
	"os"

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

	// Exemple personnel
	var scannedNames []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the first name: ")
	scanner.Scan()
	scannedNames = append(scannedNames, scanner.Text())
	for i := 0; i < 5; i++ {
		fmt.Print("Please enter the next name: ")
		scanner.Scan()
		scannedNames = append(scannedNames, scanner.Text())
	}
	messages, errs = greetings.Hellos(scannedNames)
	if errs != nil {
		log.Fatal(errs)
	}
	fmt.Println(messages)

	// Exemple personnel 2
	mascot := mascot.BestMascot()
	fmt.Println(mascot)
}
