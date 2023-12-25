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
	fmt.Println(<-c)

	// buffer channel
	c1 := make(chan int, 5)
	c1 <- 42
	c1 <- 43
	fmt.Println(<-c1)
	fmt.Println(<-c1)

}
