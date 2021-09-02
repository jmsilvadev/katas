package katas

type IKata02 interface {
	chop(subject int, listValues []int) int
}

type IterableKata struct{}

func NewIterable() *IterableKata {
	return &IterableKata{}
}

func (k *IterableKata) chop(subject int, listValues []int) int {
	return -1
}
