package leetcode

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// Remove the 2nd node from the end (should remove 4)
	result := removeNthFromEnd(head, 2)

	// Expected linked list: 1 -> 2 -> 3 -> 5
	expected := []int{1, 2, 3, 5}
	for _, val := range expected {
		if result == nil || result.Val != val {
			t.Errorf("Expected %v, got %v", val, result)
		}
		result = result.Next
	}
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
