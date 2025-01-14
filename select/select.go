// Select lets you wait on multiple channel operations

package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string) // two channels
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ { // loop to run once
		select { // use select to await both of these values simultaneously, printing each one as it arrives
		case msg1 := <-c1:
			fmt.Println("received", msg1) // received one
		case msg2 := <-c2:
			fmt.Println("received", msg2) // received two
		}
	}
}

// Run time is ~2 sec since the 1 and 2 second Sleeps execute concurrently
