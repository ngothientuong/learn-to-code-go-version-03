package main

import "fmt"

type person struct {
	first string
}

// Attach speak method to pointer of a person
func (p *person) speak() {
	fmt.Println("I am a person named ", p.first)
}

// Interface human with speak method
// All types that has speak method is said to implement to human interface
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	newPerson := person{
		first: "Tuong",
	}

	var newPerson2 person
	newPerson2 = person{
		first: "An",
	}
	// The below works
	saySomething(&newPerson)
	saySomething(&newPerson2)
	// The below doesn't work
	// saySomething(newPerson)

}
