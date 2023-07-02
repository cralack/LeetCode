package uniquebinarysearchtreesii

import (
	"testing"

	. "LeetCode/util/BinTree"
)

func generateTrees(n int) []*TreeNode {
	dp := [10][10][]*TreeNode{}
	var dfs func(lo, hi int) []*TreeNode
	dfs = func(lo, hi int) []*TreeNode {
		if dp[lo][hi] != nil {
			return dp[lo][hi]
		}
		ret := []*TreeNode{}
		if hi < lo {
			return append(ret, nil)
		}
		if lo == hi {
			dp[lo][hi] = append(ret, &TreeNode{Val: lo})
			return dp[lo][hi]
		}
		for k := lo; k <= hi; k++ {
			leftTree, rightTree := dfs(lo, k-1), dfs(k+1, hi)
			for _, left := range leftTree {
				for _, right := range rightTree {
					ret = append(ret, &TreeNode{Val: k, Left: left, Right: right})
				}
			}
		}
		dp[lo][hi] = ret
		return ret
	}
	return dfs(1, n)
}
func Test_unique_binary_search_trees_ii(t *testing.T) {
	roots := generateTrees(3)
	for _, r := range roots {
		r.Show()
		t.Log("\r\n")
	}
}
