package main

import "fmt"

// This code will fail because you can't have the main process pass in the value in the channel at one stage
// then pull a value off of a channel in the later stage
// Instead, you MUST either:
// 1. Create a buffer channel `c := make(chan int, n)` so you can then pull the value from the channel later
// 2. Pass the value in the channel INSIDE the goroutine so it can run by itself, then have the MAIN process pull the value from that go routine. Think of baton passing in a team race

func main() {
	// Create a BUFFER channel  make(chan int, n)
	// The buffer channel can hold n number of values in the buffer
	c := make(chan int, 1)
	// Pass a value into channel
	c <- 42
	fmt.Println(<-c)

	// Create another BUFFERED channel  make(chan int, n)
	// The buffer channel can hold 2 value in the buffer
	c2 := make(chan int, 2)
	c2 <- 42
	c2 <- 43
	// Receive value 42 from the buffered channel
	fmt.Println(<-c2)
	// Now the buffered channel can have another value to be pulled from, which is 43
	fmt.Println(<-c2)

}
