package wordschain

import (
	"errors"

	"github.com/jmsilvadev/cycloid/katas19/internal/bfs"
)

//Discover finds the similar words and discover the shortest chain
func Discover(subject string, target string, pathToDictionary string) ([]string, error) {
	if len(subject) != len(target) {
		return nil, errors.New("subject and target must have same length")
	}

	wchain := bfs.GetWordChain(subject, target, pathToDictionary)
	if len(wchain) == 0 {
		return nil, errors.New("there are no chains to this words in dictionary")
	}

	reversedchain := []string{}
	for i := len(wchain); i > 0; i-- {
		reversedchain = append(reversedchain, wchain[i-1])
	}

	return reversedchain, nil
}
