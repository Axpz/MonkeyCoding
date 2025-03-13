package leetcode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	result := [][]int{}

	for len(q) > 0 {
		levelQ := q
		q = []*TreeNode{}
		levelResult := []int{}

		for _, node := range levelQ {
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
			levelResult = append(levelResult, node.Val)
		}
		result = append(result, levelResult)
	}

	// reverse
	l, r := 0, len(result)-1
	for l < r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}

	return result
}
