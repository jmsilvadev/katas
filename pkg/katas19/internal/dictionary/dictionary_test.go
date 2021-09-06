package dictionary

import (
	"path/filepath"
	"testing"
)

func TestDictionary(t *testing.T) {
	realativePath, _ := filepath.Abs("../../../storage/wordlist.txt")
	for _, got := range Load(4, realativePath) {
		if len(got) != 4 {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: = 4", got)
		}
	}
	realativePath2, _ := filepath.Abs("../storage/wordlist_inexistent.txt")
	got := Load(4, realativePath2)
	if got != nil {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: nil", got)
	}

}
