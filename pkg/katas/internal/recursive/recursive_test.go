package recursive

import (
	"testing"

	"github.com/jmsilvadev/cycloid/katas/internal/helpers"
)

func TestChopRecursive(t *testing.T) {

	kata := &RecursiveKata{}
	for i, sc := range helpers.HelperScenario() {
		got := kata.chop(sc.Subject, sc.ListValues)
		if got != sc.Expected {
			t.Errorf("Got and Expected are not equals. Iteration: %v, Got: %v, expected: %v", i, got, sc.Expected)
		}

	}
}
