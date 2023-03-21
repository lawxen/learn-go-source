package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	fmt.Println("When's saturday?")
	today := time.Now().Weekday()
	fmt.Println(strconv.Atoi(today.String()))
	fmt.Println(time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
