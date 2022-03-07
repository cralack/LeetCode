package houserobberiii

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Init(arr []int) *TreeNode {
	root := &TreeNode{Val: arr[0]}
	queque := make([]*TreeNode, 0)
	queque = append(queque, root)
	for i := 1; i < len(arr); i += 2 {
		cur := queque[0]
		queque = queque[1:]
		if arr[i] != -1 {
			cur.Left = &TreeNode{Val: arr[i]}
			queque = append(queque, cur.Left)
		}
		if arr[i+1] != -1 {
			cur.Right = &TreeNode{Val: arr[i+1]}
			queque = append(queque, cur.Right)
		}
	}
	return root
}

// func rob(root *TreeNode) int {
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	if root == nil {
// 		return 0
// 	}
// 	memo := make(map[TreeNode]int)
// 	if memo[*root] > 0 {
// 		return memo[*root]
// 	}
// 	do_it := root.Val
// 	if root.Left != nil {
// 		do_it += rob(root.Left.Left) + rob(root.Left.Right)
// 	}
// 	if root.Right != nil {
// 		do_it += rob(root.Right.Left) + rob(root.Right.Right)
// 	}
// 	not_do := rob(root.Left) + rob(root.Right)
// 	res := max(do_it, not_do)
// 	memo[*root] = res
// 	return res
// }
func rob(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])
}
func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left, right := dfs(root.Left), dfs(root.Right)
	selected := root.Val + left[0] + right[0]
	unselected := max(left[0], left[1]) + max(right[0], right[1])
	return []int{unselected, selected}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_house_robber_iii(t *testing.T) {
	arr1 := []int{3, 2, 3, -1, 3, -1, 1}
	arr2 := []int{3, 4, 5, 1, 3, -1, 1}
	house1 := Init(arr1)
	house2 := Init(arr2)
	res1 := rob(house1)
	res2 := rob(house2)
	t.Log(res1, res2)
}
