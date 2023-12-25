package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	// Must have a separate goroutine to send value to the channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

// Pull value at the main goroutine
func receive(c1 <-chan int) {
	for v := range c1 {
		fmt.Println("the value received from the channel:", v)
	}
}
