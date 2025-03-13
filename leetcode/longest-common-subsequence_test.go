package leetcode

import (
	"fmt"
	"testing"
)

// ... existing code ...

func TestLongestCommonSubstring(t *testing.T) {
	tests := []struct {
		text1 string
		text2 string
		want  int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"", "abc", 0},
		{"abc", "", 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%s", tt.text1, tt.text2), func(t *testing.T) {
			got := longestCommonSubsequence(tt.text1, tt.text2)
			if got != tt.want {
				t.Errorf("longestCommonSubstring(%q, %q) = %d; want %d", tt.text1, tt.text2, got, tt.want)
			}
		})
	}
}

// ... existing code ...
