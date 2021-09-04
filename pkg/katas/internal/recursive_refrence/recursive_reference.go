package recursivereference

//Kata get the concrete object to manipualte the Kata recursive by reference
type Kata struct{}

func (k *Kata) chop(subject int, listValues []int) int {
	return k.chopRecursiveReference(subject, &listValues, 0, len(listValues)-1)
}

func (k *Kata) chopRecursiveReference(subject int, listValues *[]int, stPos int, endPos int) int {
	if stPos > endPos {
		return -1
	}
	midPos := (stPos + endPos) / 2

	if (*listValues)[midPos] == subject {
		return midPos
	}
	if subject > (*listValues)[midPos] {
		return k.chopRecursiveReference(subject, listValues, midPos+1, endPos)
	}
	return k.chopRecursiveReference(subject, listValues, stPos, midPos-1)
}
