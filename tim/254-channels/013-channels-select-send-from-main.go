package main

import "fmt"

func foo(c, quit chan int) {
	x := 3
	for {
		select {
		case c <- x:
			x += x
		case <-quit:
			fmt.Println("about to exit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	foo(c, quit)
}
