package leetcode

func maximumDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maximumDepthOfBinaryTree(root.Left), maximumDepthOfBinaryTree(root.Right))
}
