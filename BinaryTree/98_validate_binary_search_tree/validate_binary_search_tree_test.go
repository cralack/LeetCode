package validatebinarysearchtree

import (
	. "LeetCode/util/BinTree"
	"testing"
)

func isValidBST(root *TreeNode) bool {
	var isValid func(root, min, max *TreeNode) bool
	isValid = func(root, min, max *TreeNode) bool {
		if root == nil {
			return true
		}
		if min != nil && root.Val <= min.Val {
			return false
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		return isValid(root.Left, min, root) &&
			isValid(root.Right, root, max)
	}
	return isValid(root, nil, nil)
}
func Test_validate_binary_search_tree(t *testing.T) {
	arr1 := []int{2, 1, 3}
	arr2 := []int{5, 1, 4, -1, -1, 3, 6}
	roots := []*TreeNode{Init(arr1), Init(arr2)}
	for _, root := range roots {
		if isValidBST(root) {
			t.Log("true")
		} else {
			t.Log("false")
		}
	}
}
