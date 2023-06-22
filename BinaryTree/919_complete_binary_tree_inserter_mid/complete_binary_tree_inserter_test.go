package completebinarytreeinsertermid

import (
	. "LeetCode/util/BinTree"
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

type CBTInserter struct {
	Root      *TreeNode
	Candidate []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	queue := []*TreeNode{root}
	candi := []*TreeNode{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		if cur.Left == nil || cur.Right == nil {
			candi = append(candi, cur)
		}
	}
	return CBTInserter{Root: root, Candidate: candi}
}

func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{Val: val}
	cur := this.Candidate[0]
	if cur.Left == nil {
		cur.Left = node
	} else {
		cur.Right = node
		this.Candidate = this.Candidate[1:]
	}
	this.Candidate = append(this.Candidate, node)
	return cur.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.Root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
func Test_complete_binary_tree_inserter(t *testing.T) {
	c1 := []string{"CBTInserter", "insert", "insert", "get_root"}
	c2 := [][]int{{1, 2}, {3}, {4}, {-1}}
	var sol CBTInserter
	for i, c := range c1 {
		switch c {
		case "CBTInserter":
			sol = Constructor(Init(c2[i]))
		case "insert":
			sol.Insert(c2[i][0])
		case "get_root":
			sol.Get_root().Show()
		}
	}
}
