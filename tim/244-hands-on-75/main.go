package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("Value a is %v \t type %T \t dereferenced value %v\n", a, a, *a)
	fmt.Printf("Value b is %v \t type %T \t dereferenced value %v\n", b, b, *b)
	fmt.Printf("Value c is %v \t type %T \t dereferenced value %v\n", c, c, *c)
	fmt.Printf("Value d is %v \t type %T \t dereferenced value %v\n", d, d, *d)
}
