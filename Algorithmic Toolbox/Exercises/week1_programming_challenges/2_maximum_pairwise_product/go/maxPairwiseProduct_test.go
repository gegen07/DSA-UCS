package main

import "testing"

func TestMaxPairwiseProd(t *testing.T) {
	for _, val := range testCases {
		result := maxPairwiseProductFast(val.input)
		if result != val.expected {
			t.Errorf("Max Pairwise Product test with: %v\n\texpected: %d\n\t  result: %d", val.input, val.expected, result)
		}
	}
}