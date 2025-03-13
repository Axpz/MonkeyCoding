package leetcode

import (
	"reflect"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
	// Create a sample tree:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}

	expected := [][]int{{4, 5}, {2, 3}, {1}}
	result := levelOrderBottom(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
