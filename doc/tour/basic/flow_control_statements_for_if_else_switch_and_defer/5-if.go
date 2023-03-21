package main

import (
	"fmt"
	"math"
)

/*
If
Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
*/

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x)) // println printf 直接输出到终端，Sprint 不输出到终端，而是返回该字符串 
}

func main()  {
	fmt.Println(sqrt(2), sqrt(-4))
}