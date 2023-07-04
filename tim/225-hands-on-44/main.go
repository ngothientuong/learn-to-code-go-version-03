package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 := x[0:5]
	x2 := x[5:]
	x3 := x[2:7]
	x4 := x[1:6]
	fmt.Printf("%T \t %v \t %#v\n", x1, x1, x1)
	fmt.Printf("%T \t %v \t %#v\n", x2, x2, x2)
	fmt.Printf("%T \t %v \t %#v\n", x3, x3, x3)
	fmt.Printf("%T \t %v \t %#v\n", x4, x4, x4)
}
