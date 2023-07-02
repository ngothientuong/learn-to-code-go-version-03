package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		x := rand.Intn(5)
		if x == 3 {
			fmt.Printf("counter number %v: x \t %v\n", i, x)
		} else {
			fmt.Printf("Counter number %v:\n", i)
		}

	}
}
