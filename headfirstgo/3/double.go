package main

import "fmt"

func main() {
	amount := 6
	double_b(&amount)
	fmt.Println(amount)
}

func double_b (number *int) {
	*number = 2
}