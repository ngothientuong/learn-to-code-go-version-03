package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	puppy.From12()

	puppy.From13()

	//also like this
	fmt.Println(puppy.Bark(), "\n", puppy.Barks(), "\n", s3, "\n", s4)
}
