package leetcode

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{[]int{}, [][]int{{}}},
	}

	for _, test := range tests {
		result := permutations(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("permutations(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
