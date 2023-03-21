package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// v, ok := <- ch
	// if ok {
	// 	fmt.Println(v)	
	// }
	
}
