package word

import (
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	type tests struct {
		data   string
		answer map[string]int
	}
	table_test := []tests{
		tests{`one two three three three`, map[string]int{"one": 1, "two": 1, "three": 3}},
		tests{`I love apple`, map[string]int{"I": 1, "love": 1, "apple": 1}},
	}
	for _, v := range table_test {
		x := UseCount(v.data)
		for v1, _ := range x {
			if x[v1] != v.answer[v1] {
				t.Error("Error")
			}
		}
	}
}

func ExampleUseCount() {
	s := `one two three three three`
	fmt.Println(UseCount(s))
	// Output:
	// map[one:1 three:3 two:1]
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(`one two three three three`)
	}
}
