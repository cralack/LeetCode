package deepestleavessummid

import (
	. "Learning/LeetCode/BinaryTree/BinTree"
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
func deepestLeavesSum(root *TreeNode) int {
	sum, maxDep := 0, 0
	var dfs func(*TreeNode, int)
	dfs = func(cur *TreeNode, dep int) {
		if cur == nil { //base case
			return
		}
		if maxDep < dep { //抵达更深一层
			maxDep = dep
			sum = cur.Val
		} else if maxDep == dep {
			sum += cur.Val
		}
		dfs(cur.Left, dep+1)
		dfs(cur.Right, dep+1)
	}

	dfs(root, 0)
	return sum
}
func Test_deepest_leaves_sum(t *testing.T) {
	root := []int{1, 2, 3, 4, 5, -1, 6, 7, -1, -1, -1, -1, 8}
	t.Log(deepestLeavesSum(Init(root)))
	root = []int{6, 7, 8, 2, 7, 1, 3, 9, -1, 1, 4, -1, -1, -1, 5}
	t.Log(deepestLeavesSum(Init(root)))
}
