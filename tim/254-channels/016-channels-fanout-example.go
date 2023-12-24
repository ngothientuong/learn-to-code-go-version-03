package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	// send 0 to 99 to c1
	go populate(c1)

	// receive values from c1
	// use each value to create a goroutines and send result of each goroutine to c2
	go fanOutIn(c1, c2)

	// receive values from c2
	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	// close for ranging over c1 in fanOutInt()
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	// receive values from c1, create a goroutine and add to waitgroup
	for v := range c1 {
		wg.Add(1)
		// notice (v) is argument being send to v2 in parameter v2 in the annonymous function
		// v2 in essence takes value from v, meaning from 0 to 99
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

// Take a value, sleep and return a random int
func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
