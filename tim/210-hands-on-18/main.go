package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(250)
	fmt.Printf("Variable name is x with value %v\n", x)
	fmt.Println("Switch statement:")
	switch {
	case x >= 0 && x < 100:
		fmt.Println("Between 0-100")
	case x >= 101 && x < 200:
		fmt.Println("Between 101-200")
	default:
		fmt.Println("Between 201 and 250")
	}
}
