package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution.
// Concurrency refers to the ability of a system to deal with multiple tasks or operations at the same time

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {

	f("direct") // the usual way, running synchronously

	go f("goroutine") // goroutine will execute concurrently with the calling one

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
