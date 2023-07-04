package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		last      string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "James",
		last:      "Bond",
		friends:   map[string]int{"Amanda Kayle": 17, "Malzaha Tran": 28},
		favDrinks: []string{"coke", "oj"},
	}
	fmt.Println(p1.first, p1.last, p1.friends, p1.favDrinks)
	for _, v := range p1.favDrinks {
		fmt.Printf("%v \t %v \t %v\n", p1.first, p1.last, v)
	}
	for _, v := range p1.friends {
		fmt.Printf("%v \t %v \t %v\n", p1.first, p1.last, v)
	}
}
