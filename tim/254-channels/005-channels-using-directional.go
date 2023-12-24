package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)

	receive(c)
	// The below will failed if we're to receive value
	// from the same channel via a separate Goroutine as receive(c) might happens first
	// go receive (c)

	fmt.Println("about to exit")
}

// send channel
func send(c chan<- int) {
	c <- 42
}

// receive channel
func receive(c <-chan int) {
	fmt.Println("the value received from the channel:", <-c)
}
