package katas

type IKata02 interface {
	chop(subject int, listValues []int) int
}

type IterableKata struct{}

type RecursiveKata struct{}

func NewIterable() *IterableKata {
	return &IterableKata{}
}

func NewRecursive() *RecursiveKata {
	return &RecursiveKata{}
}

func (k *IterableKata) chop(subject int, listValues []int) int {
	stPos := 0
	endPos := len(listValues) - 1
	for stPos <= endPos {
		midPos := (stPos + endPos) / 2
		if listValues[midPos] == subject {
			return midPos
		}
		if subject > listValues[midPos] {
			stPos = midPos + 1
			continue
		}
		endPos = midPos - 1
	}
	return -1
}

func (k *RecursiveKata) chop(subject int, listValues []int) int {
	return k.chopRecursive(subject, listValues, 0, len(listValues)-1)
}

func (k *RecursiveKata) chopRecursive(subject int, listValues []int, stPos int, endPos int) int {
	if stPos > endPos {
		return -1
	}
	midPos := (stPos + endPos) / 2
	if listValues[midPos] == subject {
		return midPos
	}
	if subject > listValues[midPos] {
		return k.chopRecursive(subject, listValues, midPos+1, endPos)
	}
	return k.chopRecursive(subject, listValues, stPos, midPos-1)
}
