package main

import "fmt"

func main() {
	x := []int{42, 43, 44}
	result1 := foo(x...)
	fmt.Printf("%v \n", result1)
	fmt.Println("---")

	y := []int{50, 51, 52}
	result2 := bar(y)
	fmt.Printf("%v \n", result2)
}

func foo(numSlice ...int) int {
	sum := 0
	for _, v := range numSlice {
		sum += v
	}
	return sum
}

func bar(numSlice []int) int {
	sum := 0
	for _, v := range numSlice {
		sum += v
	}
	return sum
}
