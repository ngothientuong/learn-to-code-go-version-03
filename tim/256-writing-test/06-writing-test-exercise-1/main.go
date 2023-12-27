package starting_code

import (
	"fmt"
	"mymodule/tim/256-writing-test/06-writing-test-exercise-1/dog"
	// "github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/01/starting-code/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
