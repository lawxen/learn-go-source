package main

import (
	"fmt"
	"time"
)

func add(c chan int, num int) {
	num++
	Sleep()
	c <- num
}

func Sleep() {
	fmt.Println("1 second")
	time.Sleep(1*time.Second)
	fmt.Println("2 second")
	time.Sleep(1*time.Second)
	fmt.Println("3 second")
	time.Sleep(1*time.Second)
	fmt.Println("4 second")
	time.Sleep(1*time.Second)
	fmt.Println("5 second")
	time.Sleep(1*time.Second)
}
func main() {
	c := make(chan int)
	var apple int
	go add(c, apple)

	fmt.Println("before")
	result := <-c
	fmt.Println("after")
	fmt.Println(result)

}
