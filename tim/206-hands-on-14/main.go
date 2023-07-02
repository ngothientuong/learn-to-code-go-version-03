package main

import "fmt"

var outer int = 1

const mystr string = "This is outer string constant"

func main() {
	innerstr := "This is inner string"

	fmt.Printf("This is outter var: %b ,\nand here is outter mystr: %s,\n then last but not least is innerstr: %s", outer, mystr, innerstr)
}
