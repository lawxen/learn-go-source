package main

import (
	"fmt"
	// "time"
)

func main() {
	// t := time.Now()
	a := 8
	switch {
	case a < 12:
		fmt.Println("12")
	case a < 11:			// Not be executed
		fmt.Println("11")
	default:
		fmt.Println("default")
	}
}

/*
Switch with no condition

Switch without a condition is the same as switch true.
This construct can be a clean way to write long if-then-else chains.
*/