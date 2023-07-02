package main

import "fmt"

func main() {
	x := 0
	for {
		if x >= 20 {
			break
		} else {
			r := x % 2
			if r != 0 {
				fmt.Printf("x is %v and is an odd number\n", x)
				x++
			} else {
				fmt.Println("x is even")
				x++
				continue
			}

		}
	}

}
