package main

import "fmt"

func main()  {
	s := []int{2,3,5,7,11,12}

	s = s[1:4]
	fmt.Println(s)
	printSlice(s)

	s = s[:2]
	fmt.Println(s)
	printSlice(s)

	s = s[1:]
	fmt.Println(s)
	printSlice(s)

	s = s[1:3]
	fmt.Println(s)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d,%v\n",len(s),cap(s),s)  // 整数占位符：%d 十进制表示
}
/*
Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

var a [10]int
these slice expressions are equivalent:

a[0:10]
a[:10]
a[0:]
a[:]
*/