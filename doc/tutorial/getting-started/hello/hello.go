package main

import (
	"fmt"
	"github.com/lawxen/learn-golang/tree/main/doc/tutorial/create-module/greetings"
)

func main()  {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}