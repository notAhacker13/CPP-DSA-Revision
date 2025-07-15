//first program:

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, World!")
// }

//calling the greetings module

package main

import (
	"fmt"

	"abc.com/greetings"
)

func main() {
	//get a personal message and print it

	message := greetings.Hello("Ani")
	fmt.Println(message)
}
