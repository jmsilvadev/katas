package slicerecursive

//Kata get the concrete object to manipualte the Kata slicing recursive
type Kata struct{}

func (k *Kata) chop(subject int, listValues []int) int {
	return k.chopSlice(subject, listValues, 0)
}

func (k *Kata) chopSlice(subject int, listValues []int, stPos int) int {
	s := len(listValues)
	if s == 0 {
		return -1
	}
	midPos := s / 2
	if listValues[midPos] == subject {
		return midPos + stPos
	}
	if subject > listValues[midPos] {
		shift := midPos + 1
		return k.chopSlice(subject, listValues[midPos+1:s], shift+stPos)
	}
	return k.chopSlice(subject, listValues[0:midPos], stPos)
}
