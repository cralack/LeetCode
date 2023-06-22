package maximumwidthofbinarytreemid

import (
	. "LeetCode/util/BinTree"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) (ans int) {
	arr := make([]int, 0)
	var dfs func(*TreeNode, int, int)
	dfs = func(cur *TreeNode, dep, idx int) {
		if cur == nil {
			return
		}
		if len(arr) == dep {
			arr = append(arr, idx)
		}
		ans = max(ans, idx-arr[dep]+1)

		dfs(cur.Left, dep+1, idx*2)
		dfs(cur.Right, dep+1, idx*2+1)
	}

	dfs(root, 0, 1)

	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_maximum_width_of_binary_tree(t *testing.T) {
	root := Init([]int{1, 3, 2, 5, 3, -1, 9})
	t.Log(widthOfBinaryTree(root))
	root = Init([]int{1, 3, 2, 5, -1, -1, 9, 6, -1, 7})
	t.Log(widthOfBinaryTree(root))
	root = Init([]int{1, 3, 2, 5})
	t.Log(widthOfBinaryTree(root))
}
