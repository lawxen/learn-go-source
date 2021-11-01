package main

import "fmt"

type StingError struct {
	A string
	B string
}

func (se *StingError) Error() string {
	return fmt.Sprintf("Error %v %v", se.A, se.B)
}

func (se *StingError) String() string {
	return fmt.Sprintf("String %v %v", se.A, se.B)
}

func main() {
	vvv := &StingError{
		"Hi",
		"Jim",
	}

	fmt.Println(vvv)
}

/*
Execute Error() then String()
*/