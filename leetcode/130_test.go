package leetcode

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		board    [][]byte
		expected [][]byte
	}{
		{
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
			},
		},
		{
			board: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			expected: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			board: [][]byte{
				{'X'},
			},
			expected: [][]byte{
				{'X'},
			},
		},
	}

	for _, test := range tests {
		solve_union_find(test.board)
		for i := range test.board {
			for j := range test.board[i] {
				if test.board[i][j] != test.expected[i][j] {
					t.Errorf("solve() = %v, expected %v", test.board, test.expected)
				}
			}
		}
	}
}
