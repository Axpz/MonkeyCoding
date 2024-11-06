package leetcode

import (
	"testing"
)

func TestLongestFibSubseq(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{2, 4, 5, 6, 7, 8, 11, 13, 14, 15, 21, 22, 34}, 5},
		{[]int{1, 3, 7, 11, 12, 14, 18}, 3},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15}, 0},
		{[]int{1, 2, 3, 5, 8, 13, 21}, 7},
		{[]int{1, 4, 6, 7, 10, 13, 16, 19}, 4},
	}

	for _, test := range tests {
		if result := lenLongestFibSubseq(test.arr); result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.arr, test.expected, result)
		}
	}
}
