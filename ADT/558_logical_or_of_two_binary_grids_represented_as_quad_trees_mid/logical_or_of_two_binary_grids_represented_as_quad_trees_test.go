package logicaloroftwobinarygridsrepresentedasquadtreesmid

import "testing"

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */
type Node struct {
	IsLeaf      bool
	Val         bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func Constructor(arr [][]int) *Node {
	root := &Node{}
	que := []*Node{root}
	for _, ele := range arr {
		if ele[0] == -1 {
			continue
		} //pop
		cur := que[0]
		que = que[1:]
		if cur == nil {
			break
		}
		cur.IsLeaf = ele[0] == 1
		cur.Val = ele[1] == 1
		if !cur.IsLeaf {
			cur.TopLeft = &Node{}
			cur.TopRight = &Node{}
			cur.BottomLeft = &Node{}
			cur.BottomRight = &Node{}
			que = append(que, cur.TopLeft, cur.TopRight, cur.BottomLeft, cur.BottomRight)
		}
	}
	return root
}
func intersect(quadTree1, quadTree2 *Node) *Node {
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return &Node{Val: true, IsLeaf: true}
		}
		return quadTree2
	}
	if quadTree2.IsLeaf {
		return intersect(quadTree2, quadTree1)
	}
	o1 := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	o2 := intersect(quadTree1.TopRight, quadTree2.TopRight)
	o3 := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	o4 := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	if o1.IsLeaf && o2.IsLeaf && o3.IsLeaf && o4.IsLeaf && o1.Val == o2.Val && o1.Val == o3.Val && o1.Val == o4.Val {
		return &Node{Val: o1.Val, IsLeaf: true}
	}
	return &Node{false, false, o1, o2, o3, o4}
}
func Test_logical_or_of_two_binary_grids_represented_as_quad_trees(t *testing.T) {
	quadTree1 := [][]int{{0, 1}, {1, 1}, {1, 1}, {1, 0}, {1, 0}}
	quadTree2 := [][]int{{0, 1}, {1, 1}, {0, 1}, {1, 1}, {1, 0}, {-1}, {-1}, {-1}, {-1}, {1, 0}, {1, 0}, {1, 1}, {1, 1}}
	roo1 := Constructor(quadTree1)
	roo2 := Constructor(quadTree2)
	intersect(roo1, roo2)
}
