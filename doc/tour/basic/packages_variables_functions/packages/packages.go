// By convention, the package name is the same as the last element of the import path
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is",rand.Intn(10))
}