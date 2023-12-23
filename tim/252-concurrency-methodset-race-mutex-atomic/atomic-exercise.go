package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	// Define number of goroutines
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)
	// atomic works with type int64, mostly used for incrementer
	var counter int64 // counter := 0
	for i := 0; i < gs; i++ {
		// counter is a shared variable that is not inside the scope of the function below
		// This function have a bunch of go routine
		go func() {
			atomic.AddInt64(&counter, 1)
			// sleep for a bit and allow other goroutines to run
			// by yielding the processor
			// runtime.Gosched()
			fmt.Println("Counter inside goroutine's atomic is: ", atomic.LoadInt64(&counter))
			wg.Done()
			fmt.Println("Goroutines:", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter is: ", counter)
}
