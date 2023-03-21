package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	c <- 1
	c <- 2  // 在这里阻塞了
	result := <-c
	fmt.Println(result)
	result = <-c
	time.Sleep(1 * time.Second)
	fmt.Println(result)
}
