package main

import (
	"fmt"
	"../../create-module/greetings"
)

func main()  {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}

// Can't work
// hello.go:5:2: local import "../../create-module/greetings" in non-local package