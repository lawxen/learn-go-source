package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	// %q      双引号围绕的字符串，由Go语法安全地转义  Printf("%q", "Go语言")         "Go语言"
	fmt.Printf("%v %v %v %q\n", i, f, b, s)  // => 0 0 false ""
}
