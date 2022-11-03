package deletenodeinabst

import (
	. "LeetCode/BinaryTree/BinTree"
	"testing"
)

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		//case 1 & 2(无分支或单分支):
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		//case 3(双分支):
		//找到root直接后继,即右子树的最左支
		succ := root.Right
		for succ.Left != nil {
			succ = succ.Left
		}
		//先从右子树删除succ
		root.Right = deleteNode(root.Right, succ.Val)
		//将root替换为succ
		succ.Left = root.Left
		succ.Right = root.Right
		root = succ
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func Test_delete_node_in_a_bst(t *testing.T) {
	// arr := []int{5, 3, 6, 2, 4, -1, 7}
	arr := []int{36, 27, 58, 6, -1, 53, 64, -1, -1, 40, -1, -1, -1, -1, 46}
	root := Init(arr)
	key := 36
	root = deleteNode(root, key)
	root.Show()
}
