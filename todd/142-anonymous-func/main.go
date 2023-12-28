package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}() // the () at the end is the execution of the function

	func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}("Todd") // the `("Todd")` here is argument that is referenced by `(s string)` above

}

func foo() {
	fmt.Println("Foo ran")
}

// a named function with an identifier
// func (r receive) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func(p parameter(s)) (r return(s)) { code }
