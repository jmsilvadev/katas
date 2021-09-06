package recursivereference

//Chop get the position of subject in a list of values using recursive by reference way
func Chop(subject int, listValues []int) int {
	return chopRecursiveReference(subject, &listValues, 0, len(listValues)-1)
}

func chopRecursiveReference(subject int, listValues *[]int, stPos int, endPos int) int {
	if stPos > endPos {
		return -1
	}
	midPos := (stPos + endPos) / 2

	if (*listValues)[midPos] == subject {
		return midPos
	}
	if subject > (*listValues)[midPos] {
		return chopRecursiveReference(subject, listValues, midPos+1, endPos)
	}
	return chopRecursiveReference(subject, listValues, stPos, midPos-1)
}
