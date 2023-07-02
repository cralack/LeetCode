package examroom

import (
	"testing"
)

type ExamRoom struct {
	N     int
	Slice []int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		N:     n,
		Slice: make([]int, 0),
	}
}

func (p *ExamRoom) Seat() int {
	student := 0
	idx := 0
	if len(p.Slice) > 0 {
		dist := p.Slice[0] // 初始化个两座间最大距离
		pre := -1          // 初始化个前一座位号
		for i, val := range p.Slice {
			if pre != -1 {
				d := (val - pre) / 2 // 两座距离d
				if d > dist {
					dist = d
					student = pre + d // 对应座号
					idx = i           // 座号对应下标
				}
			}
			pre = val
		}
		// n-1和slice末尾间的距离>dist
		// 则将n-1插入slice末尾
		if p.N-1-p.Slice[len(p.Slice)-1] > dist {
			student = p.N - 1
			idx = len(p.Slice)
		}
	}
	p.Slice = append(p.Slice, 0) // 末尾前新增
	copy(p.Slice[idx+1:], p.Slice[idx:])
	p.Slice[idx] = student // 并更改相应座号
	return student
}

func (p *ExamRoom) Leave(x int) {
	idx := 0
	for i := 0; i < len(p.Slice); i++ {
		if p.Slice[i] == x {
			idx = i // 确定x座号对应数组下标
			break
		}
	} // 对idx前后切片拼接 -> remove idx对应元素
	p.Slice = append(p.Slice[:idx], p.Slice[idx+1:]...)
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
func Test_exam_room(t *testing.T) {
	c1 := []string{"ExamRoom", "seat", "seat", "seat", "seat", "leave", "seat"}
	c2 := []int{10, -1, -1, -1, -1, 4, -1}
	var sol ExamRoom
	for i, c := range c1 {
		switch c {
		case "ExamRoom":
			sol = Constructor(c2[i])
		case "seat":
			sol.Seat()
		case "leave":
			sol.Leave(c2[i])
		}
	}
}
