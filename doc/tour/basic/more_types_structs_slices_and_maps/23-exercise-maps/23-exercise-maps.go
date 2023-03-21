package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// return map[string]int{"x": 1}

	result := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
	  result[word] += 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}

/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/