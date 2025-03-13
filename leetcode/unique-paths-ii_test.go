package leetcode

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		obstacleGrid [][]int
		expected     int
	}{
		{
			obstacleGrid: [][]int{{0, 0}, {0, 1}, {0, 0}},
			expected:     1,
		},
		{
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expected:     2,
		},
		{
			obstacleGrid: [][]int{{1}},
			expected:     0,
		},
		{
			obstacleGrid: [][]int{{0, 0}, {0, 0}},
			expected:     2,
		},
	}

	for _, test := range tests {
		result := uniquePathsWithObstacles(test.obstacleGrid)
		if result != test.expected {
			t.Errorf("For grid %v, expected %d but got %d", test.obstacleGrid, test.expected, result)
		}
	}
}
