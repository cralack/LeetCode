package dinnerplatestacks

import (
	"container/heap"
	"sort"
	"testing"
)

type DinnerPlates struct {
	cap    int     // 栈的容量
	stacks [][]int // 所有栈
	idx    hp      // 最小堆，保存未满栈的下标
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{cap: capacity}
}

func (d *DinnerPlates) Push(val int) {
	if d.idx.Len() > 0 && d.idx.IntSlice[0] >= len(d.stacks) {
		d.idx = hp{} // 堆中都是越界下标，直接清空
	}
	if d.idx.Len() == 0 { // 所有栈都是满的
		d.stacks = append(d.stacks, []int{val}) // 添加一个新的栈
		if d.cap > 1 {                          // 新的栈没有满
			heap.Push(&d.idx, len(d.stacks)-1) // 入堆
		}
	} else { // 还有未满栈
		i := d.idx.IntSlice[0]
		d.stacks[i] = append(d.stacks[i], val) // 入栈
		if len(d.stacks[i]) == d.cap {         // 栈满了
			heap.Pop(&d.idx) // 从堆中去掉
		}
	}
}

func (d *DinnerPlates) Pop() int {
	// 等价为 popAtStack 最后一个非空栈
	return d.PopAtStack(len(d.stacks) - 1)
}

func (d *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(d.stacks) || len(d.stacks[index]) == 0 {
		return -1 // 非法操作
	}
	if len(d.stacks[index]) == d.cap { // 满栈
		heap.Push(&d.idx, index) // 元素出栈后，栈就不满了，把下标入堆
	}
	bk := len(d.stacks[index]) - 1
	val := d.stacks[index][bk]
	d.stacks[index] = d.stacks[index][:bk]
	for len(d.stacks) > 0 && len(d.stacks[len(d.stacks)-1]) == 0 {
		d.stacks = d.stacks[:len(d.stacks)-1] // 去掉末尾的空栈（懒删除，堆中下标在 push 时处理）
	}
	return val
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */
func Test_dinner_plate_stacks(t *testing.T) {
	tests := []struct {
		strcmd []string
		intcmd []int
	}{
		{[]string{"DinnerPlates", "push", "push", "push", "push", "push", "popAtStack",
			"push", "push", "popAtStack", "popAtStack", "pop", "pop", "pop", "pop", "pop"},
			[]int{2, 1, 2, 3, 4, 5, 0, 20, 21, 0, 2, -1, -1, -1, -1, -1}},
	}
	var sol DinnerPlates
	for _, tt := range tests {
		for i, c := range tt.strcmd {
			switch c {
			case "DinnerPlates":
				sol = Constructor(tt.intcmd[i])
			case "push":
				sol.Push(tt.intcmd[i])
			case "pop":
				t.Log(sol.Pop())
			case "popAtStack":
				t.Log(sol.PopAtStack(tt.intcmd[i]))
			}
		}
	}
}
