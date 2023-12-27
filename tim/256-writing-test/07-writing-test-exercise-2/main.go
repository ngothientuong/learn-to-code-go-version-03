package main

import (
	"fmt"
	"mymodule/tim/256-writing-test/07-writing-test-exercise-2/quote"
	"mymodule/tim/256-writing-test/07-writing-test-exercise-2/word"
)

func main() {
	fmt.Println("Word count: ", word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Printf("v %v \t k %v\n", v, k)
	}
}
