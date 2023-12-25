/*
write a program that
○ puts 100 numbers to a channel
○ pull the numbers off the channel and print them
*/
package main

import "fmt"

func main() {
	c := make(chan int)
	send(c)
	receive(c)
	fmt.Println("about to exit")
}

func send(c1 chan<- int) {
	// Must have a separate goroutine to send value to the channel
	go func() {
		for i := 0; i < 100; i++ {
			c1 <- i
		}
		close(c1)
	}()
}

// Pull value at the main goroutine
func receive(c1 <-chan int) {
	for v := range c1 {
		fmt.Println(v)
	}
}
