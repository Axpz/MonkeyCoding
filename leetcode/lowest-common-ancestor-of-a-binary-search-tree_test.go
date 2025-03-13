package leetcode

import "testing"

func TestLowestCommonAncestor2(t *testing.T) {
	// Create a sample binary search tree
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Right.Left = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}

	tests := []struct {
		p, q   int
		expect int
	}{
		{2, 8, 6},
		{2, 4, 2},
		{4, 5, 4},
		{0, 9, 6},
	}

	for _, test := range tests {
		p := &TreeNode{Val: test.p}
		q := &TreeNode{Val: test.q}
		result := lowestCommonAncestor(root, p, q)
		if result != nil && result.Val != test.expect {
			t.Errorf("lowestCommonAncestor(%d, %d) = %d; want %d", test.p, test.q, result.Val, test.expect)
		}
	}
}
