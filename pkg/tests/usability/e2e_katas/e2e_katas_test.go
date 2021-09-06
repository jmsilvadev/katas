package e2ekatas

import (
	"testing"

	"github.com/jmsilvadev/cycloid/katas02/iterable"
	"github.com/jmsilvadev/cycloid/katas02/recursive"
	recursivereference "github.com/jmsilvadev/cycloid/katas02/recursive_reference"
	slicerecursive "github.com/jmsilvadev/cycloid/katas02/slice_recursive"
)

func TestE2EKatasUsability(t *testing.T) {
	pos := iterable.Chop(5, []int{1, 3, 5})
	if pos != 2 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: 2", pos)
	}

	pos = recursive.Chop(5, []int{1, 3, 5})
	if pos != 2 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: 2", pos)
	}

	pos = recursivereference.Chop(5, []int{1, 3, 5})
	if pos != 2 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: 2", pos)
	}

	pos = slicerecursive.Chop(5, []int{1, 3, 5})
	if pos != 2 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: 2", pos)
	}

}
