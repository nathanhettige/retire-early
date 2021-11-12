package main

import (
	"example/greetings"
	"fmt"
	"log"
)


func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samanatha", "Darring"}

	// Request a greeting message
	messages, err := greetings.Hellos(names)
	// If an error was returns, print it to the conosle and 
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console
	fmt.Println(messages)
}
