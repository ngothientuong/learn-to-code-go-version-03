package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("Round %v\n", i)
		for x := 0; x < 100; x++ {
			fmt.Printf("%v\t", x)
		}
		fmt.Printf("End of Round %v\n", i)
	}

}
