package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m[`bond_names`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
	i := 0
	for k, v := range m {
		for j := 0; j < len(v); j++ {
			fmt.Printf("iteration %v \t key %v \t value %v \t single value %v\n", i, k, v, v[j])

		}
		i++
	}
}
