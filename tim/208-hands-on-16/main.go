package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("Variable name is x with value %v\n", x)
	if x >= 0 && x < 100 {
		fmt.Println("Between 0-100")
	} else if x >= 101 && x < 200 {
		fmt.Println("Between 101-200")
	} else {
		fmt.Println("Between 201 and 250")
	}
}
