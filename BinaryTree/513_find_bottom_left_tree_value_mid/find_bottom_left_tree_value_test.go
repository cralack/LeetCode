package findbottomlefttreevaluemid

import (
	. "LeetCode/BinaryTree/BinTree"
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
func findBottomLeftValue(root *TreeNode) (ans int) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0] //pop
		queue = queue[1:]
		//先把非空右子节点放入队列
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		//再把非空左子节点放入队列
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		//保证从右到左遍历
		ans = cur.Val
	}
	return
}
func findBottomLeftValue_DFS(root *TreeNode) (curVal int) {
	curHeight := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
		if height > curHeight {
			curHeight = height
			curVal = node.Val
		}
	}
	dfs(root, 0)
	return
}
func Test_find_bottom_left_tree_value(t *testing.T) {
	arr := []int{2, 1, 3}
	root := Init(arr)
	t.Log(findBottomLeftValue(root))

	arr = []int{1, 2, 3, 4, -1, 5, 6, -1, -1, 7}
	root = Init(arr)
	t.Log(findBottomLeftValue_DFS(root))
}

func Benchmark_find_bot_left(b *testing.B) {
	arr := []int{1, 2, 3, 4, -1, 5, 6, -1, -1, 7}
	root := Init(arr)
	b.Run("BFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findBottomLeftValue(root)
		}
		b.StopTimer()
	})
	b.Run("DFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findBottomLeftValue_DFS(root)
		}
		b.StopTimer()
	})
}
