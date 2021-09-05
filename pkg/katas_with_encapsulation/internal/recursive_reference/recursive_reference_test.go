package recursivereference

import (
	"testing"

	"github.com/jmsilvadev/cycloid/katas_with_encapsulation/internal/helpers"
)

func TestChopRecursive(t *testing.T) {

	kata := &Kata{}
	for i, sc := range helpers.HelperScenario() {
		got := kata.Chop(sc.Subject, sc.ListValues)
		if got != sc.Expected {
			t.Errorf("Got and Expected are not equals. Iteration: %v, Got: %v, expected: %v", i, got, sc.Expected)
		}

	}
}
