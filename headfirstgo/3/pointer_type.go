package main

import (
	"fmt"
)

func main() {
	//var myInt int
	//fmt.Println(reflect.TypeOf(myInt))
	//var myFloat float64
	//fmt.Println(reflect.TypeOf(myFloat))
	//var myBool bool
	//fmt.Println(reflect.TypeOf(myBool))

	fmt.Println("*****************")


	//var myInt int
	//var myIntPointer *int
	//myIntPointer = &myInt
	//fmt.Println(myIntPointer)
	//
	//
	//
	//var myBool bool
	//myBoolPointer := &myBool
	//fmt.Println(myBoolPointer)


	//myInt := 4
	//myIntPointer := &myInt
	//fmt.Println(myIntPointer)
	//fmt.Println(*myIntPointer)

	myInt := 4
	fmt.Println(myInt)
	myIntPointer := &myInt
	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}
