package wordschain

import (
	"path/filepath"
	"testing"
)

func TestWordsChain(t *testing.T) {
	pathToDictionary, _ := filepath.Abs("../../storage/wordlist.txt")
	chain, _ := Discover("cat", "dog", pathToDictionary)
	for _, got := range chain {
		if len(got) != 3 {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: != 3", got)
		}
	}

	if chain[3] != "dog" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = dog", chain[3])
	}

	if chain[2] != "cog" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = cog", chain[2])
	}

	if chain[1] != "cot" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = cot", chain[1])
	}

	if chain[0] != "cat" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = cat", chain[0])
	}

	chain, _ = Discover("ruby", "code", pathToDictionary)
	for _, got := range chain {
		if len(got) != 4 {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: != 4", got)
		}
	}

	if chain[0] != "ruby" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: cat", chain[0])
	}

	if chain[len(chain)-1] != "code" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: cat", chain[len(chain)-1])
	}

	_, err := Discover("gold", "silver", pathToDictionary)
	if err.Error() != "subject and target must have same length" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: subject and target must have same length", err.Error())
	}

	_, err = Discover("ccc", "ddd", pathToDictionary)
	if err.Error() != "there are no chains to this words in dictionary" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: there are no chains to this words in dictionary", err.Error())
	}

}
