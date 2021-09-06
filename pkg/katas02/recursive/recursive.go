package recursive

//Chop get the position of subject in a list of values using recursive way
func Chop(subject int, listValues []int) int {
	return chopRecursive(subject, listValues, 0, len(listValues)-1)
}

func chopRecursive(subject int, listValues []int, stPos int, endPos int) int {
	if stPos > endPos {
		return -1
	}
	midPos := (stPos + endPos) / 2
	if listValues[midPos] == subject {
		return midPos
	}
	if subject > listValues[midPos] {
		return chopRecursive(subject, listValues, midPos+1, endPos)
	}
	return chopRecursive(subject, listValues, stPos, midPos-1)
}
