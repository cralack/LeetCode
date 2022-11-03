package lowestcommonancestorofabinarytree

import (
	. "LeetCode/BinaryTree/BinTree"
	"testing"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//base case
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//case 1
	if left != nil && right != nil {
		return root
	}
	//case 2
	if left == nil && right == nil {
		return nil
	}
	//case 3
	if left == nil {
		return right
	}
	return left
}
func Test_lowest_common_ancestor_of_a_binary_tree(t *testing.T) {
	arr := []byte{3, 5, 1, 6, 2, 0, 8, '#', '#', 7, 4}
	root := ByteInit(arr)
	nP, nQ := 5, 4
	p := root.Find(nP)
	q := root.Find(nQ)
	anc := lowestCommonAncestor(root, p, q)
	t.Log(anc.Val)
	// t.Log(Equals(p, q))
}
