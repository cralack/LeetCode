package maximumbinarytreemid

import (
	. "LeetCode/BinaryTree/BinTree"
	"LeetCode/GetPrime"
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

func constructMaximumBinaryTree_recursion(nums []int) *TreeNode {
	var dfs func([]int) *TreeNode
	dfs = func(arr []int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		max, idx := 0, 0
		for i, v := range arr {
			if max < v {
				max = v
				idx = i
			}
		}
		return &TreeNode{
			Val:   max,
			Left:  dfs(arr[:idx]),
			Right: dfs(arr[idx+1:]),
		}
	}
	return dfs(nums)
}

func constructMaximumBinaryTree_monoStack(nums []int) *TreeNode {
	//单调递减栈
	stack := []*TreeNode{}
	for _, num := range nums {
		cur := &TreeNode{
			Val:   num,
			Left:  nil,
			Right: nil}

		//栈不为空 持续pop出所有 栈中元素 < cur.val
		for len(stack) > 0 && stack[len(stack)-1].Val < cur.Val {
			//赋为LeftChild
			cur.Left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} //栈不为空 且 cur.val < 栈顶元素
		if len(stack) > 0 {
			//赋为RichtChild
			stack[len(stack)-1].Right = cur
		}
		stack = append(stack, cur)
	}
	return stack[0]
}

func Test_maximum_binary_tree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree_recursion(nums)
	root.Show()
	root = constructMaximumBinaryTree_monoStack(nums)
	root.Show()

}

func Benchmark_maxBinTree(b *testing.B) {
	n := 10000
	arr := GetPrime.PrimeArr(n)
	n = len(arr)
	//shuffle array
	for i := 0; i < n; i++ {
		idx := rand.Intn(n)
		arr[i], arr[idx] = arr[idx], arr[i]
	}

	b.Run("recursion", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = constructMaximumBinaryTree_recursion(arr)
		}
		b.StopTimer()
	})

	b.Run("monotonic stack", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = constructMaximumBinaryTree_monoStack(arr)
		}
		b.StopTimer()
	})
}
