package main

import (
	"fmt"
	"time"
)

func main() {
	channels()
	buffering()
	synchronization()
	directions()
}

// Channels are the pipes that connect concurrent goroutines.

func channels() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

// Buffered channels accept a limited number of values without a corresponding receiver for those values.
func buffering() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// Channels can be used to synchronize execution across goroutines
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func synchronization() {

	done := make(chan bool, 1)
	go worker(done)

	<-done
}

// When using channels as function parameters, you can specify if a channel is meant to only send or receive values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func directions() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
