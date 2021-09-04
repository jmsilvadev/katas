package katas

import (
	"github.com/jmsilvadev/cycloid/katas/internal/iterable"
	"github.com/jmsilvadev/cycloid/katas/internal/recursive"
	recursivereference "github.com/jmsilvadev/cycloid/katas/internal/recursive_reference"
	slicerecursive "github.com/jmsilvadev/cycloid/katas/internal/slice_recursive"
)

//IKata02 Kata interface to implements methods to manipulate kata structures
type IKata02 interface {
	chop(subject int, listValues []int) int
}

//NewIterable gets a Kata iterable object
func NewIterable() *iterable.Kata {
	return &iterable.Kata{}
}

//NewRecursive gets a Kata recursive object
func NewRecursive() *recursive.Kata {
	return &recursive.Kata{}
}

//NewRecursiveReference gets a Kata recursive object
func NewRecursiveReference() *recursivereference.Kata {
	return &recursivereference.Kata{}
}

//NewRecursiveSlice gets a Kata recursive object
func NewRecursiveSlice() *slicerecursive.Kata {
	return &slicerecursive.Kata{}
}
