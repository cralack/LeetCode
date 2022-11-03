package maximumsumbstinbinarytree

import (
	. "LeetCode/BinaryTree/BinTree"
	"testing"
)

func maxSumBST(root *TreeNode) int {
	const MAX_value = int(^uint(0) >> 1)
	findMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const MIN_value = ^MAX_value
	findMin := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const (
		isBST = iota
		min
		max
		sum
	)
	maxSum := 0
	// 函数返回 []int{ isBST, min, max, sum}
	var traverse func(node *TreeNode) []int
	traverse = func(node *TreeNode) []int {
		//base case
		if node == nil {
			return []int{1, MAX_value, MIN_value, 0}
		}
		// 递归计算左右子树
		left := traverse(node.Left)
		right := traverse(node.Right)
		/******* 后序遍历位置 *******/
		res := [4]int{}
		//左右子树都是BST
		if left[isBST] == 1 && right[isBST] == 1 &&
			//左子树max < node.Val < 右子树min
			left[max] < node.Val && node.Val < right[min] {
			// 以 node 为根的二叉树是 BST
			res[isBST] = 1
			// 计算以 node 为根的这棵 BST 的最小值
			res[min] = findMin(left[min], node.Val)
			// 计算以 node 为根的这棵 BST 的最大值
			res[max] = findMax(right[max], node.Val)
			// 计算以 node 为根的这棵 BST 所有节点之和
			res[sum] = left[sum] + node.Val + right[sum]
			// 更新全局变量
			maxSum = findMax(maxSum, res[sum])
		} else {
			res[isBST] = 0
		}
		return res[:]
	}
	traverse(root)
	return maxSum
}
func Test_maximum_sum_bst_in_binary_tree(t *testing.T) {
	// arr := []byte{1, 4, 3, 2, 4, 2, 5, '#', '#', '#', '#', '#', '#', 4, 6}
	arr := []byte{4, 3, '#', 1, 2}
	root := ByteInit(arr)
	res := maxSumBST(root)
	t.Log(res)
}
