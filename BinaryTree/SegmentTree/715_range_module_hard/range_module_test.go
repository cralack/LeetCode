package rangemodulehard

import (
	"testing"
)

const MAXRANGE = 1000000009

type RangeModule struct {
	Root *SegmentNode
}

func Constructor() RangeModule {
	return RangeModule{&SegmentNode{Ls: nil, Rs: nil, Val: false, Add: false}}
}

func (p *RangeModule) AddRange(left int, right int) {
	p.Root.update(1, MAXRANGE, left, right-1, true)
}

func (p *RangeModule) QueryRange(left int, right int) bool {
	return p.Root.query(1, MAXRANGE, left, right-1)
}

func (p *RangeModule) RemoveRange(left int, right int) {
	p.Root.update(1, MAXRANGE, left, right-1, false)
}

type SegmentNode struct {
	Ls, Rs *SegmentNode //左右区间子节点
	Val    bool         //当前区间有多少个整数被追踪
	Add    bool         //懒标记
}

func (node *SegmentNode) update(lc, rc int, l, r int, v bool) {
	if l <= lc && rc <= r {
		node.Val, node.Add = v, true //注意产生变化懒标记就为True，因为更新有删除
		return
	}
	node.pushdown()
	mid := lc + (rc-lc)>>1
	if l <= mid {
		node.Ls.update(lc, mid, l, r, v)
	}
	if mid < r {
		node.Rs.update(mid+1, rc, l, r, v)
	}
	node.pushup()
}
func (node *SegmentNode) query(lc, rc int, l, r int) bool {
	if l <= lc && rc <= r {
		return node.Val
	}
	//先确保所有关联的懒标记下沉下去
	node.pushdown()
	mid, ans := lc+(rc-lc)>>1, true
	if l <= mid {
		ans = ans && node.Ls.query(lc, mid, l, r)
	}
	if mid < r {
		//同样为不同题目中的更新方式
		ans = ans && node.Rs.query(mid+1, rc, l, r)
	}
	return ans
}
func (node *SegmentNode) pushup() {
	//动态更新方式：此处为两者都true
	node.Val = node.Ls.Val && node.Rs.Val
}
func (node *SegmentNode) pushdown() {
	//懒标记, 在需要的时候才开拓节点和赋值
	if node.Ls == nil {
		node.Ls = &SegmentNode{Ls: nil, Rs: nil, Val: false, Add: false}
	}
	if node.Rs == nil {
		node.Rs = &SegmentNode{Ls: nil, Rs: nil, Val: false, Add: false}
	}
	if !node.Add {
		return
	}
	//注意产生变化懒标记就为True，因为更新有删除
	node.Ls.Val, node.Ls.Add = node.Val, true
	node.Rs.Val, node.Rs.Add = node.Val, true
	node.Add = false
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
func Test_range_module(t *testing.T) {
	c1 := []string{"RangeModule", "addRange", "removeRange",
		"queryRange", "queryRange", "queryRange"}
	c2 := [][]int{{-1}, {10, 20}, {14, 16}, {10, 14}, {13, 15}, {16, 17}}
	var sol RangeModule
	for i, c := range c1 {
		switch c {
		case "RangeModule":
			sol = Constructor()
		case "addRange":
			sol.AddRange(c2[i][0], c2[i][1])
		case "removeRange":
			sol.RemoveRange(c2[i][0], c2[i][1])
		case "QueryRange":
			t.Log(sol.QueryRange(c2[i][0], c2[i][1]))
		}
	}
}
