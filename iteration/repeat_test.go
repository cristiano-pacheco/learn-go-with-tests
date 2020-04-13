package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}

	repeated = Repeat("a", 0)
	expected = ""

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}

	repeated = Repeat("aa", 3)
	expected = "aaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("ha", 3)
	fmt.Println(repeated)
	// Output: hahaha
}
