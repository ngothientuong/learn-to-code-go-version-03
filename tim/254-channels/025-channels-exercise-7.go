package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	c := make(chan int)
	// Lanch 1 big goroutine. In side this big goroutines, run multiple sets where each set is with 10 sub-goroutines
	go Fanout(c)
	receive(c)
	fmt.Println("about to exit")
	fmt.Println("num gortins:", runtime.NumGoroutine())
}

func Fanout(c1 chan<- int) {
	// Define 10 goroutines
	const goroutines = 10
	// Define throttling waitgroup
	var wg sync.WaitGroup
	// to process X numbers of process at a time
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		fmt.Println("Round i: ", i)
		// each goroutine add 10 number to the channels
		go func(c2 chan<- int, i int) {
			// Do something - send values to channel c2
			for j := 0; j < 10; j++ {
				fmt.Printf("i\t %v \t j \t %v \t (10*i+j) \t %v \n", i, j, 10*i+j)
				c2 <- 10*i + j
				fmt.Println("num gortins:", runtime.NumGoroutine())
			}
			wg.Done()
		}(c1, i) // send incrementer from bigger loop to the sub-goroutine
	}
	wg.Wait()
	close(c1)
}

func receive(c2 <-chan int) {
	for v := range c2 {
		fmt.Println(v)
	}
}
