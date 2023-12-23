package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// Define WaitGroup and add 2 goroutines
	wg.Add(2)

	fmt.Println("This is main goroutine")
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("At the end of main goroutine")
}

func foo() {
	time.Sleep(time.Second)
	fmt.Println("This is foo goroutine")
	wg.Done()
}
func bar() {
	time.Sleep(time.Second)
	fmt.Println("This is bar goroutine")
	wg.Done()
}
