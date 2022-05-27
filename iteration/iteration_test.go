package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 5)
	expected := "aaaaa"

	if expected != result {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestRepeat2(t *testing.T) {
	result := Repeat2("a", 5)
	expected := "aaaaa"

	if expected != result {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	out := Repeat("c", 3)
	fmt.Println(out)

	// Output: ccc
}
