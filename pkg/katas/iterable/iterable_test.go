package iterable

import (
	"testing"

	"github.com/jmsilvadev/cycloid/katas/internal/helpers"
)

func TestChopIterable(t *testing.T) {

	for i, sc := range helpers.HelperScenario() {
		got := Chop(sc.Subject, sc.ListValues)
		if got != sc.Expected {
			t.Errorf("Got and Expected are not equals. Iteration: %v, Got: %v, expected: %v", i, got, sc.Expected)
		}

	}
}
