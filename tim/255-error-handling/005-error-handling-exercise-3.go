package main

import "fmt"

type customErr struct {
	idx   int
	last  string
	first string
}

// Attach method Error() to the type customError
func (cs customErr) Error() string {
	return fmt.Sprintf("Error with index %v, last name %v and first name %v", cs.idx, cs.last, cs.first)
}

func main() {
	custom_error1 := customErr{1, "Ngo", "Tuong"}
	custom_error2 := customErr{idx: 2, last: "Tran", first: "An"}
	foo(custom_error1)
	// Or pointer of a struct can implicitly (automatically) have Error() method attached to it
	foo(&custom_error1)

	foo(custom_error2)
	// Or pointer of a struct can implicitly (automatically) have Error() method attached to it
	foo(&custom_error2)

}

func foo(e error) {
	fmt.Println("Error is: ", e)
}
