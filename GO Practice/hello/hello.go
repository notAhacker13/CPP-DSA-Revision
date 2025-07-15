package main

import (
	"fmt"

	"log"

	"abc.com/greetings"
)

func main() {

	//setting properties of the predefined logger
	//the log entry prefix and a flag to disable printing
	//the time, source file, and line number

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Aniket", "Chandramani Tripathi", "Winter"}

	//get a personal message and print it

	messages, err := greetings.Hellos(names)

	//if an error was returned, print to console andexit the program
	if err != nil {
		log.Fatal(err)
	}

	//in case no error
	fmt.Println(messages)
}
