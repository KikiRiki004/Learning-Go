package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // 1 low, 4 high; takes 1 through 3 elements
	fmt.Println(s)

	var k []int = primes[2:6] // 2 low, 6 high; takes 2 through 5 elements
	fmt.Println(k)
}
