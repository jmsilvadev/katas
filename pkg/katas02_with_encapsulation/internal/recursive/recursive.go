package recursive

//Kata get the concrete object to manipualte the Kata recursive
type Kata struct{}

//Chop get the position of subject in a list of values using recursive way
func (k *Kata) Chop(subject int, listValues []int) int {
	return k.chopRecursive(subject, listValues, 0, len(listValues)-1)
}

func (k *Kata) chopRecursive(subject int, listValues []int, stPos int, endPos int) int {
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
