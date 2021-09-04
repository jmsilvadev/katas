package katas

import (
	"github.com/jmsilvadev/cycloid/katas/internal/iterable"
	"github.com/jmsilvadev/cycloid/katas/internal/recursive"
)

type IKata02 interface {
	chop(subject int, listValues []int) int
}

func NewIterable() *iterable.IterableKata {
	return &iterable.IterableKata{}
}

func NewRecursive() *recursive.RecursiveKata {
	return &recursive.RecursiveKata{}
}
