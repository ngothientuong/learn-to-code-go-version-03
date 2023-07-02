package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// CONDITIONAL
	// if statements
	// switch statements

	// concurrency
	// select a channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	// get an int64, convert it to type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type duration's underlying type is int64
	// time.Sleep takes type duration
	// time.Millisecond is a constant in the time package
	// https://pkg.go.dev/time#pkg-constants
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	fmt.Printf("d1 is %v \t %T\n", d1, d1)
	time.Sleep(d1 * time.Millisecond)
	fmt.Println("Done")

	fmt.Printf("d2 is %v \t %T\n", d2, d2)
	time.Sleep(d2 * time.Millisecond)
	fmt.Println("Done")

	// Create a 1st goroutine on a channel, make the 1st goroutine sleep for amount of time
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()
	// Create a 2nd goroutine on a channel that runs concurrently relatively to above 1st goroutine, make the 2nd goroutine sleep for amount of time
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#Select_statements
	// Whichever channel comes back witha value first, select will pick that case!!!
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}
}
