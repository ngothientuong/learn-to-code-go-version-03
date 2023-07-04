package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m[`bond_names`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	for k, v := range m {
		fmt.Printf("%v\n", k)
		for i, v2 := range v {
			fmt.Printf("%v \t %v\n", i, v2)
		}
	}

	delete(m, `no_dr`)
	fmt.Println("--------")
	for k, v := range m {
		fmt.Printf("%v\n", k)
		for i, v2 := range v {
			fmt.Printf("%v \t %v\n", i, v2)
		}
	}
}
