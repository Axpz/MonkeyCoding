package leetcode

import "testing"

func TestFindNumberOfLIS(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 1},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 7},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := findNumberOfLIS(test.nums)
		if result != test.expect {
			t.Errorf("For input %v, expected %d but got %d", test.nums, test.expect, result)
		}
	}
}
