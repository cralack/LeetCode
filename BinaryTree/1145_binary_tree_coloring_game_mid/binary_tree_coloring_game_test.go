package binarytreecoloringgame

import (
	. "LeetCode/BinaryTree/BinTree"
	"math/rand"
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

// x越靠近叶子节点性能越优
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var xNode *TreeNode
	var dfs func(*TreeNode) int
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 0
		}
		if cur.Val == x {
			xNode = cur
		}
		return 1 + dfs(cur.Left) + dfs(cur.Right)
	}
	dfs(root)
	lCsize := dfs(xNode.Left)
	if lCsize >= (n+1)/2 {
		return true
	}
	rCsize := dfs(xNode.Right)
	if rCsize >= (n+1)/2 {
		return true
	}
	remain := n - lCsize - rCsize - 1
	return remain >= (n+1)/2
}

// x越靠近根节点性能越优
func btreeGameWinningMove_v2(root *TreeNode, n int, x int) bool {
	lSize, rSize := 0, 0
	var dfs func(*TreeNode) int
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 0
		}
		ls := dfs(cur.Left)
		rs := dfs(cur.Right)
		if cur.Val == x {
			lSize = ls
			rSize = rs
		}
		return ls + rs + 1
	}
	dfs(root)
	return max(max(lSize, rSize), n-1-lSize-rSize)*2 > n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_binary_tree_coloring_game(t *testing.T) {
	root := Init([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	n, x := 11, 3
	t.Log(btreeGameWinningMove(root, n, x))
	root = Init([]int{1, 2, 3})
	n, x = 3, 1
	t.Log(btreeGameWinningMove_v2(root, n, x))
	n = 1e3

	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	x = rand.Intn(n) + 1 //[1,n]
	root = Init(arr)
	t.Log(btreeGameWinningMove(root, n, x))
	t.Log(btreeGameWinningMove_v2(root, n, x))
}

func Benchmark_btree(b *testing.B) {
	var n int = 1e3
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	x := rand.Intn(n) + 1
	// x := 3
	root := Init(arr)

	b.Run("v1", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			btreeGameWinningMove(root, n, x)
		}
		b.StopTimer()
	})

	b.Run("v2", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			btreeGameWinningMove_v2(root, n, x)
		}
		b.StopTimer()
	})
}
