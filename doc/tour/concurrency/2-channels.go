package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("sum before")
	c <- sum
	fmt.Println("20 before")
	c <- 20
	fmt.Println("30 after")
	c <- 30
	fmt.Println("40 before")
	c <- 40
	fmt.Println("test")
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	time.Sleep(1*time.Second)
	fmt.Println("1 second")
	time.Sleep(1*time.Second)
	fmt.Println("2 second")
	time.Sleep(1*time.Second)
	fmt.Println("3 second")
	time.Sleep(1*time.Second)
	fmt.Println("4 second")
	go sum(s[len(s)/2:], c)
	time.Sleep(1*time.Second)
	fmt.Println("1 second")
	time.Sleep(1*time.Second)
	fmt.Println("2 second")
	time.Sleep(1*time.Second)
	fmt.Println("3 second")
	time.Sleep(1*time.Second)
	fmt.Println("4 second")
	time.Sleep(10*time.Second)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
