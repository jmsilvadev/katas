package helpers

//Scenarios struct of scenarios
type Scenarios []struct {
	Subject    int
	ListValues []int
	Expected   int
}

//HelperScenario gets an array of scenarios to use with tests
func HelperScenario() Scenarios {
	return Scenarios{
		{
			Subject:    3,
			ListValues: []int{},
			Expected:   -1,
		},
		{
			Subject:    3,
			ListValues: []int{1},
			Expected:   -1,
		},
		{
			Subject:    1,
			ListValues: []int{1},
			Expected:   0,
		},
		{
			Subject:    1,
			ListValues: []int{1, 3, 5},
			Expected:   0,
		},
		{
			Subject:    3,
			ListValues: []int{1, 3, 5},
			Expected:   1,
		},
		{
			Subject:    5,
			ListValues: []int{1, 3, 5},
			Expected:   2,
		},
		{
			Subject:    0,
			ListValues: []int{1, 3, 5},
			Expected:   -1,
		},
		{
			Subject:    2,
			ListValues: []int{1, 3, 5},
			Expected:   -1,
		},
		{
			Subject:    4,
			ListValues: []int{1, 3, 5},
			Expected:   -1,
		},
		{
			Subject:    6,
			ListValues: []int{1, 3, 5},
			Expected:   -1,
		},
		{
			Subject:    1,
			ListValues: []int{1, 3, 5, 7},
			Expected:   0,
		},
		{
			Subject:    3,
			ListValues: []int{1, 3, 5, 7},
			Expected:   1,
		},
		{
			Subject:    5,
			ListValues: []int{1, 3, 5, 7},
			Expected:   2,
		},
		{
			Subject:    7,
			ListValues: []int{1, 3, 5, 7},
			Expected:   3,
		},
		{
			Subject:    0,
			ListValues: []int{1, 3, 5, 7},
			Expected:   -1,
		},
		{
			Subject:    2,
			ListValues: []int{1, 3, 5, 7},
			Expected:   -1,
		},
		{
			Subject:    4,
			ListValues: []int{1, 3, 5, 7},
			Expected:   -1,
		},
		{
			Subject:    6,
			ListValues: []int{1, 3, 5, 7},
			Expected:   -1,
		},

		{
			Subject:    8,
			ListValues: []int{1, 3, 5, 7},
			Expected:   -1,
		},
	}
}
