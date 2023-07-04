package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Printf("%v \t %v\n", len(xs1), cap(xs1))
	xs2 := []string{"Miss", "Moneypenny", "I'm 008."}
	xxs := [][]string{xs1, xs2}
	fmt.Println("-------")
	fmt.Printf("%#v \t %v\n", xxs, xxs)
	for i := 0; i < cap(xxs); i++ {
		for j := 0; j < cap(xs2); j++ {
			fmt.Printf("xxs[%v][%v] is %v\n", i, j, xxs[i][j])
		}
	}
}
