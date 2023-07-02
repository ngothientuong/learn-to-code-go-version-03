package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for name, age := range m {
		fmt.Printf("name is %v and age is %v\n", name, age)
	}
	testName := "James"
	if val, ok := m[testName]; ok {
		fmt.Printf("Test name %v is with value %v\n", testName, val)
	}

}
