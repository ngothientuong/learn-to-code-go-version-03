package main

import "fmt"

type T struct {
	first string
}

func valueSemantic(val T) string {
	var returnV = val.first
	return returnV
}

func pointerSemantic(pointerVal *T) {
	var returnV = pointerVal.first
	fmt.Printf("Value of pointerVal is %v\n", returnV)
}

func main() {
	var tim = T{"Tuong"}
	var a = valueSemantic(tim)
	fmt.Printf("Value a is %v\n", a)
	pointerSemantic(&tim)
}
