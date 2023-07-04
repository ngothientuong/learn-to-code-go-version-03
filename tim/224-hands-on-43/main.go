package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Printf("iteration %v \t %v \t %T \t %#v\n", i, v, v, v)
	}
	fmt.Printf("%T \t %v \t %#v\n", x, x, x)
}
