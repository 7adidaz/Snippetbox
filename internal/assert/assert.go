package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper() // telling the test runner that this function is a test helper
	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}
