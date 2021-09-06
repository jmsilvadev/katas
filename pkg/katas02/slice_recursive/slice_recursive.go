package slicerecursive

//Chop get the position of subject in a list of values using slicing way
func Chop(subject int, listValues []int) int {
	return chopSlice(subject, listValues, 0)
}

func chopSlice(subject int, listValues []int, stPos int) int {
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
		return chopSlice(subject, listValues[midPos+1:s], shift+stPos)
	}
	return chopSlice(subject, listValues[0:midPos], stPos)
}
