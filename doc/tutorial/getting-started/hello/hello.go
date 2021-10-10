// Work
// go mod edit -replace github.com/lawxen/learn-golang/tree/main/doc/tutorial/create-module/greetings=../../greetings
package main

import (
	"fmt"
    "log"
	"github.com/lawxen/learn-golang/tree/main/doc/tutorial/create-module/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    message, err := greetings.Hellos(names)
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}

/*
go build: compiles the packages, along with their dependencies, but it doesn't install the results.

go
*/