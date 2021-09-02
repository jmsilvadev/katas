package katas

import "testing"

type scenarios []struct {
	subject    int
	listValues []int
	expected   int
}

func TestChop(t *testing.T) {

	kata := NewIterable()

	for _, sc := range helperScenario() {
		got := kata.chop(sc.subject, sc.listValues)
		if got != sc.expected {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: %v", got, sc.expected)
		}

	}
}

func helperScenario() scenarios {
	return scenarios{
		{
			subject:    3,
			listValues: []int{},
			expected:   -1,
		},
		{
			subject:    3,
			listValues: []int{1},
			expected:   -1,
		},
		{
			subject:    1,
			listValues: []int{1},
			expected:   0,
		},
		{
			subject:    1,
			listValues: []int{1, 3, 5},
			expected:   0,
		},
		{
			subject:    3,
			listValues: []int{1, 3, 5},
			expected:   1,
		},
		{
			subject:    5,
			listValues: []int{1, 3, 5},
			expected:   2,
		},
		{
			subject:    0,
			listValues: []int{1, 3, 5},
			expected:   -1,
		},
		{
			subject:    2,
			listValues: []int{1, 3, 5},
			expected:   -1,
		},
		{
			subject:    4,
			listValues: []int{1, 3, 5},
			expected:   -1,
		},
		{
			subject:    6,
			listValues: []int{1, 3, 5},
			expected:   -1,
		},
		{
			subject:    1,
			listValues: []int{1, 3, 5, 7},
			expected:   0,
		},
		{
			subject:    3,
			listValues: []int{1, 3, 5, 7},
			expected:   1,
		},
		{
			subject:    5,
			listValues: []int{1, 3, 5, 7},
			expected:   2,
		},
		{
			subject:    7,
			listValues: []int{1, 3, 5, 7},
			expected:   3,
		},
		{
			subject:    0,
			listValues: []int{1, 3, 5, 7},
			expected:   -1,
		},
		{
			subject:    2,
			listValues: []int{1, 3, 5, 7},
			expected:   -1,
		},
		{
			subject:    4,
			listValues: []int{1, 3, 5, 7},
			expected:   -1,
		},
		{
			subject:    6,
			listValues: []int{1, 3, 5, 7},
			expected:   -1,
		},

		{
			subject:    8,
			listValues: []int{1, 3, 5, 7},
			expected:   -1,
		},
	}
}
