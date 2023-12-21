package beautiful_towers_ii_mid

import (
	"testing"
)

func maximumSumOfHeights(maxHeights []int) (ans int64) {
	n := len(maxHeights)
	left := make([]int, n)
	right := make([]int, n)
	stk := make(stack, 0, n)

	for i := range left {
		left[i] = -1
		right[i] = n
	}

	// 维护以i为顶左侧第一个小坐标
	for i, x := range maxHeights {
		// ↗栈
		for !stk.isEmpty() && maxHeights[stk.top()] > x {
			stk.pop()
		}
		if !stk.isEmpty() {
			left[i] = stk.top()
		}
		stk.push(i)
	}
	stk = make(stack, 0, n)
	for i := n - 1; i >= 0; i-- {
		x := maxHeights[i]
		for !stk.isEmpty() && maxHeights[stk.top()] >= x {
			stk.pop()
		}
		if !stk.isEmpty() {
			right[i] = stk.top()
		}
		stk.push(i)
	}

	f := make([]int64, n)
	g := make([]int64, n)
	// 以i为顶的左侧高度和
	for i, x := range maxHeights {
		if i > 0 && x >= maxHeights[i-1] {
			// 升高直接+
			f[i] = f[i-1] + int64(x)
		} else {
			// 降高分块+
			j := left[i]
			f[i] = int64(x) * int64(i-j)
			if j != -1 {
				f[i] += f[j]
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		x := maxHeights[i]
		if i < n-1 && x >= maxHeights[i+1] {
			g[i] = g[i+1] + int64(x)
		} else {
			j := right[i]
			g[i] = int64(x) * int64(j-i)
			if j != n {
				g[i] += g[j]
			}
		}
	}
	for i, x := range maxHeights {
		ans = max(ans, f[i]+g[i]-int64(x))
	}

	return
}

type stack []int

func (s *stack) isEmpty() bool { return len(*s) == 0 }
func (s *stack) top() int      { x := *s; return x[len(x)-1] }
func (s *stack) pop() int      { x := *s; tmp := s.top(); *s = x[:len(x)-1]; return tmp }
func (s *stack) push(x int)    { *s = append(*s, x) }

func Test_beautiful_towers_ii(t *testing.T) {
	tests := []struct{ maxHeights []int }{
		{[]int{5, 3, 4, 1, 1}},
		{[]int{6, 5, 3, 9, 2, 7}},
		{[]int{3, 2, 5, 5, 2, 3}},
	}
	for _, tt := range tests {
		t.Log(maximumSumOfHeights(tt.maxHeights))
	}
}
