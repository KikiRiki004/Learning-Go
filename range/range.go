package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { //goes through a slice or a map
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
