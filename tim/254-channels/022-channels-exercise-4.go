package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	// Must have a separate goroutine to send value to the channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		q <- 1
	}()

	return c
}

// Pull value at the main goroutine
func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
