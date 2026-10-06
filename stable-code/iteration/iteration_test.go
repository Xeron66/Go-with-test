package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected: %q, but Got: %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// here the Loop() method keep on running Repeat() up untill the benchmark is done
	for b.Loop() {
		Repeat("a")
	}
}
