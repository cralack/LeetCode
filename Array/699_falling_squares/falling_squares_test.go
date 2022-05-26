package fallingsquares

import "testing"

type Node struct {
	l, r, h, maxR int
	left, right   *Node
}

func fallingSquares(positions [][]int) []int {
	var res = make([]int, 0)
	var root *Node = new(Node)
	maxH := 0
	for _, position := range positions {
		l := position[0]
		r := l + position[1]
		e := position[1]
		curH := query(root, l, r)
		root = insert(root, l, r, curH+e)
		maxH = max(maxH, curH+e)
		res = append(res, maxH)
	}
	return res
}

func insert(root *Node, l int, r int, h int) *Node {
	if root == nil {
		return &Node{
			l:    l,
			r:    r,
			h:    h,
			maxR: r,
		}
	}
	if l <= root.l {
		root.left = insert(root.left, l, r, h)
	} else {
		root.right = insert(root.right, l, r, h)
	}
	root.maxR = max(r, root.maxR)
	return root
}

func query(root *Node, l int, r int) int {
	if root == nil || l >= root.maxR {
		return 0
	}
	curH := 0
	if !(r <= root.l || l >= root.r) {
		curH = root.h
	}
	curH = max(curH, query(root.left, l, r))
	if r >= root.l {
		curH = max(curH, query(root.right, l, r))
	}
	return curH
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_falling_squares(t *testing.T) {
	positions := [][]int{{1, 2}, {2, 3}, {6, 1}}
	t.Log(fallingSquares(positions))
	positions = [][]int{{100, 100}, {200, 100}}
	t.Log(fallingSquares(positions))
	positions = [][]int{{1, 2}, {1, 3}}
	t.Log(fallingSquares(positions))
}
