package word

import (
	"fmt"
	"testing"

	"github.com/Marisame254/golang_course/ninja13/exercise2/quote"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func TestUseCount(t *testing.T) {
	n := UseCount("one two three three three")
	for k, v := range n {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "expect", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "expect", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "expect", 1)
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "expect", 3)
	}
}
