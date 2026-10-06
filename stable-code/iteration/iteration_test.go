package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Expected: %q, but Got: %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	// here the Loop() method keep on running Repeat() up untill the benchmark is done
	for b.Loop() {
		Repeat("a", 5)
	}
}
