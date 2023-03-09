package main

import (
	"fmt"

	"github.com/lawxen/learn-golang/tree/main/doc/tutorial/create-module/stage1/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Jim")
	fmt.Println(message)
}
