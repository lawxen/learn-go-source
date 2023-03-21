package main

import (
	"fmt"
	"time"
	// "time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			// time.Sleep(5*time.Second)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			// fmt.Println("123")
		}
		time.Sleep(1*time.Second)
		quit <- 0
	}()
	fibonacci(c, quit)
}
