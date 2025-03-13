package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode
	var curr *ListNode = head
	for curr != nil {
		next := curr.Next

		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
