package iterable

//Kata get the concrete object to manipualte the Kata Iterable
type Kata struct{}

func (k *Kata) chop(subject int, listValues []int) int {
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
