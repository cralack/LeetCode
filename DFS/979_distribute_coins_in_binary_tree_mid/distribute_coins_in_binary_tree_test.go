package distributecoinsinbinarytree

import (
	"testing"

	. "LeetCode/util/BinTree"
)

func distributeCoins_v1(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		coinsL, nodesL := dfs(node.Left)
		coinsR, nodesR := dfs(node.Right)
		coins := coinsL + coinsR + node.Val // 子树硬币个数
		nodes := nodesL + nodesR + 1        // 子树节点数
		ans += abs(coins - nodes)
		return coins, nodes
	}
	dfs(root)
	return
}

func distributeCoins_v2(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		coinsL, nodesL := dfs(node.Left)
		coinsR, nodesR := dfs(node.Right)
		coins := coinsL + coinsR + node.Val // 子树硬币个数
		nodes := nodesL + nodesR + 1        // 子树节点数
		ans += abs(coins - nodes)
		return coins, nodes
	}
	dfs(root)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Test_distribute_coins_in_binary_tree(t *testing.T) {
	tests := []struct {
		root  *TreeNode
		wanna int
	}{
		{Init([]int{3, 0, 0}), 2},
		{Init([]int{0, 3, 0}), 3},
		{Init([]int{1, 0, 2}), 2},
		{Init([]int{1, 0, 0, -1, 3}), 4},
	}
	for i, tt := range tests {
		t.Logf("\n\rtest %d:\n\r\tv1:%v\n\r\tv2:%v", i+1,
			distributeCoins_v1(tt.root) == tt.wanna,
			distributeCoins_v2(tt.root) == tt.wanna,
		)
	}
}
