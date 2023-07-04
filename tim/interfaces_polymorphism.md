# Type Struct & implementing method

// func (r receiver) identifier(p parameter) (return (s)) { code }

Joke The interface says to a struct: "Hey baby! If you have these methods, then you're my type"

When a type T is a receiver of a function which certain identifier, then the type T is said to have a method of identifier

Example:
```
type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}
```

The type `person` above is said to have a method `speak()`. That is, you can call `p.speak()`

# Interface
An interface I is a set of methods. If a type T has the method in the interface, then the VALUE of that type T is said to be a type of interface I.

Example: 

- Below, both `type secretAgent` and `type person` has method `speak()`. Since `type human interface` is a collection of method `speak()`, you can use use value of both `type secretAgent` or `type person` in place of `type human interface`. 
- That is, `h human` is a passed as an argument in a function as a generalized form, but in reality, you can pass in either `type secretAgent` or `type person` as an argument!

```
package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	p2 := person{
		first: "Jenny",
	}

	// sa1.speak()
	// p2.speak()

	saySomething(sa1)
	saySomething(p2)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

```