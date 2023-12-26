package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	// Placing defer func as FIRST IN , it will be executed LAST when panic happened
	// and it allows the ability to recover!
	defer func() {
		// if the goroutine is in panic then the function recover() be panic => r = panic
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	// g will call itself => g(0), g(1), ...g(4)
	// when it hits g(4), panic() will be call and defer functions will be executed in the order of LAST IN, FIRST OUT ()
	g(0)
	// The next line fmt.Println("Returned normally from g.") will not run as only defer function are executed in panic()!
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	// calling itself => g(0), g(1), ...g(4)
	g(i + 1)
}
