package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// Must have a separate goroutine to send value to the channel
	go func() {
		c <- 42
	}()
	// Pull value at the main goroutine
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)
	// Pull value at the main goroutine
	v, ok = <-c
	fmt.Println(v, ok)
}
