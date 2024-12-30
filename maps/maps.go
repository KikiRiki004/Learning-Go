package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // map is a built-in data type that associates keys with values

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])

// 	offset := timeZone["EST"]

// 	fmt.Println(offset)
// }

//Exercise:

// WordCount counts the occurrences of each word in the input string.
func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)   // Fields function from strings library splits the string into words
	for _, word := range words { // Range through "words" since it is a slice
		wordMap[word]++
	}
	return wordMap
}

func main() {
	// Example input
	s := "hello world hello world hello hello dude"
	result := WordCount(s)
	fmt.Println(result) // Output: map[hello:2 world:1]
}
