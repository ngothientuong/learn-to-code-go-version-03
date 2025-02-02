package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	// User bytes.Buffer to as space to write out the json string
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(&users)
	if err != nil {
		fmt.Println("We did something wrong and here is the error: ", err)
	}
	fmt.Println("\n")
	fmt.Printf("Here is byte buffer in json with type\n: %+v \t %T", b.String(), b)

	// You can also write to os.Stdout instead of bytes.Buffer
	fmt.Println("\n")
	err = json.NewEncoder(os.Stdout).Encode(&users)
	if err != nil {
		fmt.Println("We did something wrong and here is the error: ", err)
	}

	// Example of initialize a type vs producing a pointer to its type
	// Both cases, you can use the method belong to the type for the variables produced
	var c bytes.Buffer
	fmt.Printf("\nThis initializes a new bytes.Buffer: var c bytes.Buffer \n%v\n", c)
	d := bytes.NewBufferString(`my string`)
	fmt.Printf("This produces a new bytes.Buffer: d := bytes.NewBufferString(`my string`)\n %s", d)
}
