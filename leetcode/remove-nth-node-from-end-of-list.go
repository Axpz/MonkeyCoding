package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	h, k := dummy, dummy
	for i := 0; i <= n; i++ {
		h = h.Next
	}

	for h != nil {
		h = h.Next
		k = k.Next
	}

	k.Next = k.Next.Next

	return dummy.Next
}
