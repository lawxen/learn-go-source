package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	var dates [3]time.Time
	dates[0] = time.Unix(1258894000, 0)
	dates[1] = time.Unix(1258894000, 0)
	dates[2] = time.Unix(1258894000, 0)
	fmt.Println(dates[0])

	// Nil
	fmt.Println(primes[3])

	notesb := [7]string{"do", "re", "mi", "fa", "so"}
	primesb := [5]int{2, 3, 4, 5, 7}

	text := [3]string{
		"This is a seris of long strings",
		"This is a seris of long strings",
		"This is a seris of long strings",
	}

	fmt.Println(notesb)
	fmt.Println(primesb)
	fmt.Println(text)

	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}

	for index, value := range notes {
		fmt.Println(index, value)
	}
}
