package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// func main() {

// 	res := plus(1, 2)
// 	fmt.Println("1+2 =", res)

// 	res = plusPlus(1, 2, 3)
// 	fmt.Println("1+2+3 =", res)
// }

func vals() (int, int) { // The (int, int) in this function signature shows that the function returns 2 ints.
	return 3, 7
}

// func main() {

// 	a, b := vals()
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	_, c := vals()  // If you only want a subset of the returned values, use the blank identifier _.
// 	fmt.Println(c)
// }

// Variadic functions can be called with any number of trailing arguments.
func sum(nums ...int) { //Here’s a function that will take an arbitrary (any) number of ints as arguments.
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// func main() {

// 	sum(1, 2)
// 	sum(1, 2, 3)

// 	nums := []int{1, 2, 3, 4}
// 	sum(nums...)
// }

func intSeq() func() int { // This function intSeq returns another function, which we define anonymously in the body of intSeq.
	i := 0
	return func() int { // The returned function closes over the variable i to form a closure.
		i++
		return i
	}
}

// func main() {

// 	nextInt := intSeq()

// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())

// 	newInts := intSeq()
// 	fmt.Println(newInts())
// }

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

// Example

func counter() func() int {
	count := 0 // Treasure chest
	return func() int {
		count++      // Update the chest
		return count // Tell you the current value
	}
}

// func main() {
//     c := counter()   // Get a key to the chest
//     fmt.Println(c()) // Output: 1
//     fmt.Println(c()) // Output: 2
//     fmt.Println(c()) // Output: 3
// }

// Exercise

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }
