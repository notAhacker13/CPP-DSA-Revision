package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"))

	fmt.Println("\n below is new output")

	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
