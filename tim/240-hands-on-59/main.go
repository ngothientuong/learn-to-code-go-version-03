package main

import "fmt"

func main() {
	defer runLater(1)
	defer runLater(2)
	defer runLater(3)
	x := []int{42, 43, 44}
	result1 := foo(x...)
	fmt.Printf("%v \n", result1)
	fmt.Println("---")

}

func foo(numSlice ...int) int {
	sum := 0
	for _, v := range numSlice {
		sum += v
	}
	return sum
}

func runLater(s int) {
	fmt.Printf("This function %v should show last because of defer func\n", s)
}
