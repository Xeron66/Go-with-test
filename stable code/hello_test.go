package main

import (
	"testing"
)

/*
	Here, we are introducing another tool in our testing arsenal: subtests.
	Sometimes, it is useful to group tests around a "thing" and then have subtests describing different scenarios.
*/

func TestHello(t *testing.T) {

	// test run for Hello function with an argument
	t.Run("saying, Hello to people", func(t *testing.T) {
		actual := Hello("Chris")
		expected := "Hello, Chris"

		if actual != expected {
			t.Errorf("Actual : %q but Expected: %q", actual, expected)
		}
	})

	// test run for Hello function with no argument
	t.Run("say, 'Hello, World' when empty string is supplied", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, World"

		if actual != expected {
			t.Errorf("Actual : %q but Expected: %q", actual, expected)
		}
	})
}
