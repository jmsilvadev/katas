package bfs

import (
	"path/filepath"
	"testing"

	"github.com/jmsilvadev/cycloid/katas19/internal/dictionary"
)

var pathToDictionary string

func TestGetNextVertex(t *testing.T) {
	pathToDictionary, _ = filepath.Abs("../../../storage/wordlist.txt")
	words := dictionary.Load(3, pathToDictionary)
	subject := "cat"
	vertex := getNextVertex(subject, words)
	for _, got := range vertex {
		if len(got) != 3 {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: = 3", got)
		}
	}
}

func TestIsSimilar(t *testing.T) {
	subject := "cat"
	target := "cot"
	got := isSimilar(subject, target)
	if !got {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: true", got)
	}

	subject = "cat"
	target = "dog"
	got = isSimilar(subject, target)
	if got {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: false", got)
	}
}

func TestGetShortestPath(t *testing.T) {
	markedVertex := make(map[string]string)
	markedVertex["dog"] = "cog"
	markedVertex["cot"] = "cat"
	markedVertex["cog"] = "com"
	markedVertex["com"] = "cot"
	subject := "cat"
	target := "dog"
	path := getShortestPath(subject, target, &markedVertex)
	for _, got := range path {
		if len(got) != 3 {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: = 3", got)
		}
	}

	if path[0] != "dog" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = dog", path[0])
	}

	if path[2] != "com" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = com", path[2])
	}

	if path[4] != "cat" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = cat", path[4])
	}
}

func TestGetWordChain(t *testing.T) {
	subject := "cat"
	target := "dog"
	chain := GetWordChain(subject, target, pathToDictionary)
	if chain[0] != "dog" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = dog", chain[0])
	}

	if chain[1] != "cog" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = cog", chain[1])
	}

	if chain[2] != "cot" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = cot", chain[2])
	}

	if chain[3] != "cat" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: = cat", chain[3])
	}

	got := GetWordChain(subject, "cccc", pathToDictionary)
	if got != nil {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: nil", got)
	}
}
