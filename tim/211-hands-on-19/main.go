package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("Variable name is x with value %v\n", x)
	fmt.Printf("Variable name is y with value %v\n", y)

	if x < 4 && y < 4 {
		fmt.Println("x & y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x & y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("None of the previous cases were met")
	}

}
