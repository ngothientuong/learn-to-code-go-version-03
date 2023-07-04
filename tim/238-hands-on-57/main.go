package main

import "fmt"

func main() {
	num, str := bar()
	fmt.Printf("%v\n", foo())
	fmt.Printf("%v \t %v \n", num, str)
}

func foo() int {
	return 3
}
func bar() (int, string) {
	return 3, "Hello World"
}
