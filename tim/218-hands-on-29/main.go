package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47}
	for _, v := range x {
		fmt.Printf("v is %v\n", v)
	}

}
