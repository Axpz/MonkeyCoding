package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		names    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"", "b", "c"}, ""},
		{[]string{"a"}, "a"},
	}

	for _, test := range tests {
		result := longestCommonPrefix(test.names)
		if result != test.expected {
			t.Errorf("longestcommonprefix(%v) = %v; expected %v", test.names, result, test.expected)
		}
	}
}
