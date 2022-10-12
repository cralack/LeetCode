package findmodeinbinarysearchtree

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
func findMode(root *TreeNode) (res []int) {
	maxCnt, cnt := 0, 0
	base := 0
	var dfs func(*TreeNode)
	dfs = func(cur *TreeNode) {
		//base case
		if cur == nil {
			return
		}

		dfs(cur.Left)
		//inorder
		if cur.Val == base {
			cnt++
		} else {
			base = cur.Val
			cnt = 1
		}
		if maxCnt < cnt {
			maxCnt = cnt
			res = []int{cur.Val}
		} else if maxCnt == cnt {
			res = append(res, cur.Val)
		}

		dfs(cur.Right)
	}

	dfs(root)
	return //O(n)
}

func findMode_v2(root *TreeNode) (res []int) {
	freq := make(map[int]int, 0)
	max := 0
	var dfs func(*TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		freq[cur.Val]++
		if max < freq[cur.Val] {
			max = freq[cur.Val]
		}
		dfs(cur.Left)
		dfs(cur.Right)
	}
	dfs(root)
	for k, v := range freq {
		if v == max {
			res = append(res, k)
		}
	}
	return
}
func Test_find_mode_in_binary_search_tree(t *testing.T) {
	root := Init([]int{1, -1, 2, 2})
	t.Log(findMode(root))
	root = Init([]int{5, 2, 6, 2, -1, 6, -1})
	t.Log(findMode(root))
}
