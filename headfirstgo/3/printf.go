package main

import "fmt"

func main() {
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Println(area/10.0, "liters needed!")

	width = 5.2
	height = 3.5
	area = width * height
	fmt.Println(area/10.0, "liters needed")

	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
	resultString := fmt.Sprintf("About one-third: %0.2f\n", 1.0/3.0)
	fmt.Printf(resultString)

	fmt.Printf("%v \n", 1.0/3.0)
	fmt.Printf("%v \n", "\n")

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("---------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %3d\n", "Paper clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	fmt.Println("********************")
	fmt.Printf("%%7.3f: %20.3f\n", 12.3456)
}
