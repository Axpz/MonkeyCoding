package leetcode

import (
	"testing"
)

func TestRepeatLimitedString(t *testing.T) {
	tests := []struct {
		s           string
		repeatLimit int
		expected    string
	}{
		{"cczazcc", 3, "zzcccac"},
		{"aababab", 2, "bbabaa"},
		{"a", 1, "a"},
		{"aaabbbccc", 2, "ccbcbbaa"},
		{"abcabcabc", 1, "cbcbcba"},
	}

	for _, test := range tests {
		result := repeatLimitedString(test.s, test.repeatLimit)
		if result != test.expected {
			t.Errorf("For input s=%s and repeatLimit=%d, expected %s but got %s", test.s, test.repeatLimit, test.expected, result)
		}
	}
}
