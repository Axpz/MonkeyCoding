package leetcode

import (
	"testing"
)

func TestCalShortestJobFirst(t *testing.T) {
	reqTimes := []int{0, 1, 2, 3, 4}
	durTimes := []int{3, 9, 6, 1, 2}
	expected := 3.0

	result := calShortestJobFirst(reqTimes, durTimes)
	if result != expected {
		t.Errorf("expected %f, got %f", expected, result)
	}
}
