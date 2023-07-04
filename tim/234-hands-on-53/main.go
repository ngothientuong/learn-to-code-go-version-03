package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first:   "James",
		last:    "Bond",
		flavors: []string{"chocolate", "blue berry", "coffee"},
	}
	p2 := person{
		first:   "Jenny",
		last:    "Walker",
		flavors: []string{"milk", "black mama", "vanila"},
	}
	fmt.Printf("%#v \t %#v \t  %T \n", p1, p1.last, p1)
	fmt.Printf("%#v \t %#v \t %T \n", p2, p2.last, p2)
	for i, v := range p1.flavors {
		fmt.Printf("%#v favoriate ice cream is: %#v \t %#v \n", p1.first, i, v)
	}
	for i, v := range p2.flavors {
		fmt.Printf("%#v favoriate ice cream is:%#v \t %#v \n", p2.first, i, v)
	}
}
