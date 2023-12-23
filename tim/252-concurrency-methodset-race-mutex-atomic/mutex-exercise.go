package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	// Wait for 100 goroutines
	wg.Add(100)

	// Use mutex to lock the code containing the shared variable
	var mu sync.Mutex
	var counter int // counter := 0
	for i := 0; i < 100; i++ {
		// counter is a shared variable that is not inside the scope of the function below
		// This function have a bunch of go routine
		go func() {
			mu.Lock()
			v := counter
			// sleep for a bit and allow other goroutines to run
			// by yielding the processor
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			fmt.Println("Counter in goroutine is: ", counter)
			fmt.Println("=============")
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Final Counter is: ", counter)
}
