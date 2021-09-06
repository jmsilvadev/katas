package bfs

import (
	"github.com/jmsilvadev/cycloid/katas19/internal/dictionary"
)

func getNextVertex(subject string, words []string) []string {
	vertex := []string{}
	for _, w := range words {
		if isSimilar(w, subject) {
			vertex = append(vertex, w)
		}
	}
	return vertex
}

func isSimilar(w string, subject string) bool {
	similarity := 0
	for i := 0; i < len(subject); i++ {
		if subject[i] != w[i] {
			similarity++
		}
	}
	if similarity > 1 {
		return false
	}
	return true
}

func getShortestPath(subject string, target string, markedVertex *map[string]string) []string {
	if _, exists := (*markedVertex)[target]; !exists {
		return nil
	}

	tree := []string{}
	currentWord := target
	for currentWord != subject {
		tree = append(tree, currentWord)
		currentWord = (*markedVertex)[currentWord]
	}
	tree = append(tree, subject)
	return tree
}

//GetWordChain build the chains of the words using BFS and returns the shortest path
func GetWordChain(subject string, target string, path string) []string {
	words := dictionary.Load(len(subject), path)
	markedVertex := make(map[string]string)
	fifoPipe := []string{}
	fifoPipe = append(fifoPipe, subject)
	for len(fifoPipe) > 0 {
		if fifoPipe[0] == target {
			break
		}
		vertices := getNextVertex(fifoPipe[0], words)
		for _, vertex := range vertices {
			if _, exists := markedVertex[vertex]; !exists {
				markedVertex[vertex] = fifoPipe[0]
				fifoPipe = append(fifoPipe, vertex)
			}
		}
		fifoPipe = fifoPipe[1:]
	}
	return getShortestPath(subject, target, &markedVertex)
}
