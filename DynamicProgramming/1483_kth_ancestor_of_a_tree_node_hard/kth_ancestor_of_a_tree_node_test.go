package kthancestorofatreenode

import "testing"

// 树的层深?
const mx int = 18

type TreeAncestor struct {
	memo [][mx]int
}

func Constructor(n int, parent []int) TreeAncestor {
	dp := make([][mx]int, n)
	for i, fa := range parent {
		dp[i][0] = fa
		for j := 1; j < mx; j++ {
			dp[i][j] = -1
		}
	}
	for i := range dp {
		for j := 1; j < mx; j++ {
			if dp[i][j-1] == -1 {
				continue
			}
			dp[i][j] = dp[dp[i][j-1]][j-1]
		}
	}
	return TreeAncestor{dp}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	for i := mx - 1; i >= 0; i-- {
		// k 的二进制从低到高第 i 位是 1
		if k>>i&1 == 1 {
			node = this.memo[node][i]
			if node == -1 {
				break
			}
		}
	}
	return node
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
func Test_kth_ancestor_of_a_tree_node(t *testing.T) {
	tests := []struct {
		n      int
		parent []int
		query  [][2]int
	}{
		{n: 7, parent: []int{-1, 0, 0, 1, 1, 2, 2},
			query: [][2]int{{3, 1}, {5, 2}, {6, 3}}},
	}
	for _, tt := range tests {
		sol := Constructor(tt.n, tt.parent)
		for _, que := range tt.query {
			t.Log(sol.GetKthAncestor(que[0], que[1]))
		}
	}
}
