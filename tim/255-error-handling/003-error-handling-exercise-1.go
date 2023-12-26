package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	// first defer function used to recover as it will be ran last in case of panic!
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("This func will be run last as a way to recover")
		}
	}()
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		// fmt.Println("Error is: ", err)
		// log.Println("Error is: ", err)
		log.Panicln("Panic mode will be ran and will trigger defer func at the top of the code. The error is: ", err)
	}

	fmt.Println(string(bs))

}
