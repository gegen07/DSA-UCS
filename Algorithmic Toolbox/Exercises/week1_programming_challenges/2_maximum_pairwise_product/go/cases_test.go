package main

var testCases = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, 2, 3},
		expected: 6,
	},
	{
		input:    []int{20, 6, 120, 43, 12, 17},
		expected: 5160,
	},
	{
		input:    []int{90, 12738, 97, 12763, 654, 8171235, 2163},
		expected: 104289472305,
	},
}
