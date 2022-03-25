package twosumivinputisabst

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
func findTarget_BST_DFS(root *TreeNode, k int) bool {
	vis := make(map[int]bool, 0)
	flag := false
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil || flag == true {
			return
		}
		if vis[k-root.Val] {
			flag = true
			return
		}
		vis[root.Val] = true
		if k-root.Val > root.Val {
			dfs(root.Right)
			dfs(root.Left)
		} else {
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(root)
	return flag
}
func findTarget_BFS(root *TreeNode, k int) bool {
	vis := make(map[int]bool, 0)
	que := []*TreeNode{root}
	for len(que) > 0 {
		//pop
		cur := que[0]
		que = que[1:]
		if _, ok := vis[k-cur.Val]; ok {
			return true
		}
		vis[cur.Val] = true
		if cur.Left != nil {
			que = append(que, cur.Left)
		}
		if cur.Right != nil {
			que = append(que, cur.Right)
		}
	}
	return false
}
func findTarget_DFS(root *TreeNode, k int) bool {
	vis := make(map[int]bool, 0)
	var dfs func(cur *TreeNode) bool
	dfs = func(cur *TreeNode) bool {
		if cur == nil {
			return false
		}
		if vis[k-cur.Val] {
			return true
		}
		vis[cur.Val] = true
		return dfs(cur.Left) || dfs(cur.Right)
	}
	return dfs(root)
}
func Test_two_sum_iv_input_is_a_bst(t *testing.T) {
	arr := []int{5, 3, 6, 2, 4, -1, 7}
	root := Init(arr)
	t.Log(findTarget_DFS(root, 9))
	t.Log(findTarget_DFS(root, 28))
	arr = []int{2, 1, 3}
	root = Init(arr)
	t.Log(findTarget_DFS(root, 4))
}

func Benchmark_two_sum_iv_input_is_a_bst(b *testing.B) {
	arr1 := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	root1 := Init(arr1)
	arr2 := []int{5, 3, 6, 2, 4, -1, 7}
	root2 := Init(arr2)
	arr3 := []int{2, 1, 3}
	root3 := Init(arr3)
	b.Run("BST_DFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findTarget_BST_DFS(root1, 13)
			findTarget_BST_DFS(root2, 9)
			findTarget_BST_DFS(root3, 4)
		}
		b.StopTimer()
	})
	b.Run("BFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findTarget_BFS(root1, 13)
			findTarget_BFS(root2, 9)
			findTarget_BFS(root3, 4)
		}
		b.StopTimer()
	})
	b.Run("DFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findTarget_DFS(root1, 13)
			findTarget_DFS(root2, 9)
			findTarget_DFS(root3, 4)
		}
		b.StopTimer()
	})
}
