package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	// Test cases
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{12321, true},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%d) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
