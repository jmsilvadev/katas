package e2ekatas19

import (
	"path/filepath"
	"testing"

	"github.com/jmsilvadev/cycloid/katas19/wordschain"
)

func TestE2EKatas19Usability(t *testing.T) {
	pathToDictionary, _ := filepath.Abs("../../../storage/wordlist.txt")
	chain, _ := wordschain.Discover("the", "end", pathToDictionary)

	if chain[0] != "the" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: cat", chain[0])
	}

	if chain[len(chain)-1] != "end" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: cat", chain[len(chain)-1])
	}
}
