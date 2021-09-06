package recursive

import (
	"testing"

	"github.com/jmsilvadev/cycloid/katas02/internal/helpers"
)

func TestChopRecursive(t *testing.T) {
	for i, sc := range helpers.HelperScenario() {
		got := Chop(sc.Subject, sc.ListValues)
		if got != sc.Expected {
			t.Errorf("Got and Expected are not equals. Iteration: %v, Got: %v, expected: %v", i, got, sc.Expected)
		}

	}
}
