package main

import "fmt"

// This code will fail because you can't have the main process pass in the value in the channel at one stage
// then pull a value off of a channel in the later stage
// Instead, you MUST either:
// 1. Create a buffer channel `c := make(chan int, n)` so you can then pull the value from the channel later
// 2. Pass the value in the channel INSIDE the goroutine so it can run by itself, then have the MAIN process pull the value from that go routine. Think of baton passing in a team race

func main() {
	// Create a channel
	c := make(chan int)
	// Create a goroutine, send the value
	go func() {
		// Pass a value into channel
		c <- 42
	}()
	// Other goroutine SEND value 42  <---- Channel ---> Main Goroutine RECEIVE value 42
	fmt.Println(<-c)
}
