package bintreediameter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		} else {
			leftD := dfs(node.Left)
			rightD := dfs(node.Right)
			if leftD+rightD > maxDiameter {
				maxDiameter = leftD + rightD
			}
			return Max(leftD, rightD) + 1
		}
	}
	dfs(root)
	return maxDiameter
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
