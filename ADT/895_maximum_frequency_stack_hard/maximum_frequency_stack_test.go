package maximumfrequencystack

import (
	"testing"
)

// graceful code
type FreqStack struct {
	cnt    map[int]int
	stacks [][]int
}

func Constructor() FreqStack {
	return FreqStack{cnt: map[int]int{}}
}

func (f *FreqStack) Push(val int) {
	c := f.cnt[val]
	if c == len(f.stacks) { // 这个元素的频率已经是目前最多的，现在又出现了一次
		f.stacks = append(f.stacks, []int{val}) // 那么必须创建一个新栈
	} else {
		f.stacks[c] = append(f.stacks[c], val) // 否则就压入对应的栈
	}
	f.cnt[val]++ // 更新频率
}

func (f *FreqStack) Pop() int {
	back := len(f.stacks) - 1
	st := f.stacks[back]
	bk := len(st) - 1
	val := st[bk] // 弹出最右侧栈的栈顶
	if bk == 0 {  // 栈为空
		f.stacks = f.stacks[:back] // 删除
	} else {
		f.stacks[back] = st[:bk]
	}
	f.cnt[val]-- // 更新频率
	return val
}
func Test_maximum_frequency_stack(t *testing.T) {
	c1 := []string{"FreqStack", "push", "push",
		"push", "push", "push", "push",
		"pop", "pop", "pop", "pop"}
	c2 := []int{-1, 5, 7,
		5, 7, 4, 5,
		-1, -1, -1, -1}
	// c1 := []string{"FreqStack", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop"}
	// c2 := []int{-1, 30, 40, 4, 25, 20, 40, 48, 21, 12, 44, 1, 16, 20, 9, 34, 26, 12, 21, 35, 16, 3, 23, 9, 31, 10, 6, -1, 45, -1, 16, -1, 14, -1, 27, -1, 35, -1, 34, -1, 40, -1, 13, -1, 21, -1, 18, -1, 26, -1, 29, -1, 32, -1, 3, -1, 18, -1, 36, -1, 1, -1, 38, -1, 34, -1, 20, -1, 22, -1, 13, -1, 37, -1, 24, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	var stack FreqStack
	for idx, str := range c1 {
		// sb := &strings.Builder{}
		// sb.WriteString(strconv.Itoa(idx))
		// sb.WriteString(": c1:")
		// sb.WriteString(str)
		// sb.WriteString(" ")
		// sb.WriteString(strconv.Itoa(c2[idx]))
		// WriteLog.WriteFile("max_freq_stack", sb.String())
		switch str {
		case "FreqStack":
			stack = Constructor()
		case "push":
			stack.Push(c2[idx])
		case "pop":
			stack.Pop()
		}
	}
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
