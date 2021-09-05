package e2ekataswithinternals

import (
	"testing"

	kataswithinternals "github.com/jmsilvadev/cycloid/katas_with_encapsulation"
)

func TestE2EKastasWithInternalsUsability(t *testing.T) {
	k1 := kataswithinternals.NewIterable()
	pos := k1.Chop(5, []int{1, 3, 5})
	if pos != 2 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: 2", pos)
	}

	k2 := kataswithinternals.NewRecursive()
	pos = k2.Chop(5, []int{1, 3, 5})
	if pos != 2 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: 2", pos)
	}

	k3 := kataswithinternals.NewRecursiveReference()
	pos = k3.Chop(5, []int{1, 3, 5})
	if pos != 2 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: 2", pos)
	}

	k4 := kataswithinternals.NewRecursiveSlice()
	pos = k4.Chop(5, []int{1, 3, 5})
	if pos != 2 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: 2", pos)
	}

}
