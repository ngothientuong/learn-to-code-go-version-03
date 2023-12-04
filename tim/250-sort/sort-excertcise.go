package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	// Simple Sort
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)

	// Custom Sort
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
	fmt.Println("\n")

	// your code goes here
	// Sort by Age - pass in VALUE of type ByAge
	sort.Sort(ByAge(users))
	fmt.Println("SORT BY AGE")
	for i, v := range users {
		fmt.Printf("Person %d: \t %v\n", i, v)
	}

	// Sort by Last Name - pass in VALUE of type ByLast
	sort.Sort(ByLast{u1, u2, u3})
	// or:  sort.Sort(ByLast(users))
	fmt.Println("\nSORT BY LAST NAME")
	for i, v := range users {
		fmt.Printf("Person %d: \t %v\n", i, v)
	}

}

// Sort custom
// Sort by Age
// Implement the 3 required methods for type ByAge
type ByAge []user

func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

// Sort by Last
// Implement the 3 required method for type ByLast
type ByLast []user

func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }
