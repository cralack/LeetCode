package printbinarytreemid

import (
	"strconv"
	"testing"

	. "LeetCode/util/BinTree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	maxDep := -1
	// get maxDep
	var dfs_height func(*TreeNode, int)
	dfs_height = func(cur *TreeNode, dep int) {
		if cur == nil {
			return
		}
		if dep > maxDep {
			maxDep = dep
		}
		dfs_height(cur.Left, dep+1)
		dfs_height(cur.Right, dep+1)
	}
	dfs_height(root, 0)
	// init matrix
	m, n := maxDep+1, pow(2, maxDep+1)-1
	matrix := make([][]string, m)
	for i := range matrix {
		matrix[i] = make([]string, n)
		// for j := range matrix[i] {
		// 	matrix[i][j] = "*"
		// }
	}
	// get matrix
	var dfs_matrix func(*TreeNode, int, int)
	dfs_matrix = func(cur *TreeNode, r, c int) {
		if cur == nil {
			return
		}
		matrix[r][c] = strconv.Itoa(cur.Val)

		dfs_matrix(cur.Left, r+1, c-pow(2, maxDep-r-1))
		dfs_matrix(cur.Right, r+1, c+pow(2, maxDep-r-1))
	}
	dfs_matrix(root, 0, (n-1)/2)

	return matrix
}
func pow(x, n int) (ans int) {
	if x == 2 && n >= 0 {
		ans = 1 << n
		return
	}
	ans = 1
	for i := 0; i < n; i++ {
		ans *= x
	}
	return
}
func Test_print_binary_tree(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(pow(2, i))
	}
	nums := []int{1, 2}
	root := Init(nums)
	t.Log(printTree(root))
	nums = []int{1, 2, 3, -1, 4}
	root = Init(nums)
	t.Log(printTree(root))

}
