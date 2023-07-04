package main

import "fmt"

func main() {
	x := [5]int{}
	for i := 0; i < 5; i++ {
		x[i] = i*i + 1
	}
	for i, v := range x {
		fmt.Printf("iteration %v \t %v \t %T\n", i, v, v)
	}
}
