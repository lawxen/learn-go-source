package main

import "fmt"

func sum(s []int, c chan int)  {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main()  {
	s := []int{7,33,65,2,45,6}

	c := make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)

	x,y := <-c,<-c // receive from c

	fmt.Println(x,y,x+y)
}