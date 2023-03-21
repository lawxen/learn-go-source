package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
		fmt.Println("3333")
		c <- 2
		fmt.Println("4444")
	}()
	result := <-c
	fmt.Println(result)
	result = <-c
	time.Sleep(1 * time.Second)
	fmt.Println(result)
}
