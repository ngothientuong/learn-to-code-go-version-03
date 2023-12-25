package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	// Populate channel c1 with 1000 number in a separate Goroutine
	go populate(c1)

	// While populating channel c1, create 10 separate goroutines where EACH continously pull from c1 at the same time relatively compare to each other
	// and feed the values to c2!!!
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 1000; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)
	// similar to example 016 in the same folder
	// But limit to just 10 Goroutines
	for i := 0; i < goroutines; i++ {
		// process 10 Goroutines at a time, then the next 10 set
		// although you sent 100 Goroutine to the wg
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
					fmt.Println("num gortins:", runtime.NumGoroutine())
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
