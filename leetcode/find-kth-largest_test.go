package leetcode

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{1}, 1, 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums=%v,k=%d", tt.nums, tt.k), func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("findKthLargest(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
