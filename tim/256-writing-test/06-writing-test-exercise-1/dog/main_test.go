package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	type test struct {
		data   int
		answer int
	}
	table_test := []test{
		test{1, 7},
		test{2, 14},
		test{3, 21},
	}
	for _, v := range table_test {
		x := Years(v.data)
		if x != v.answer {
			t.Error("Expected ", v.answer, "Got", x)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	type test struct {
		data   int
		answer int
	}
	table_test := []test{
		test{1, 7},
		test{2, 14},
		test{3, 21},
	}
	for _, v := range table_test {
		x := YearsTwo(v.data)
		if x != v.answer {
			t.Error("Expected ", v.answer, "Got", x)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func ExampleYearsTwo() {
	fmt.Println(Years(8))
	// Output:
	// 56
}

func BenchmarkYears(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(9)
	}
}
