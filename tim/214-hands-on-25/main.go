package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		switch {
		case x == 0:
			fmt.Printf("iteration is %v and x is %v\n", i, x)
		case x == 1:
			fmt.Printf("iteration is %v and x is %v\n", i, x)
		case x == 2:
			fmt.Printf("iteration is %v and x is %v\n", i, x)
		case x == 3:
			fmt.Printf("iteration is %v and x is %v\n", i, x)
		case x == 4:
			fmt.Printf("iteration is %v and x is %v\n", i, x)
		}
	}

}
