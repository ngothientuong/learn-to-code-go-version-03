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
	// fmt.Printf("%v \t %v \t  %T \n", p1, p1.last, p1)
	// fmt.Printf("%#v \t %#v \t  %T \n", p1, p1.last, p1)
	// fmt.Printf("%#v \t %#v \t %T \n", p2, p2.last, p2)
	// for i, v := range p1.flavors {
	// 	fmt.Printf("%#v favoriate ice cream is: %#v \t %#v \n", p1.first, i, v)
	// }
	// for i, v := range p2.flavors {
	// 	fmt.Printf("%#v favoriate ice cream is:%#v \t %#v \n", p2.first, i, v)
	// }

	m := make(map[string]person)

	m[p1.last] = p1
	m[p2.last] = p2

	m1 := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}
	fmt.Printf("%v \t %T\n", m, m)
	fmt.Printf("%v \t %T\n", m1, m1)

	for k, v := range m {
		fmt.Printf("%v \t %v\n", k, v)
		for _, v1 := range v.flavors {
			fmt.Printf("%v \t %v \t %v \n", v.first, v.last, v1)
		}
	}
	for k, v := range m1 {
		fmt.Printf("%v \t %v\n", k, v)
		for _, v1 := range v.flavors {
			fmt.Printf("%v \t %v \t %v \n", v.first, v.last, v1)
		}
	}
}
