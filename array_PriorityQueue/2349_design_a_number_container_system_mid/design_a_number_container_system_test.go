package design_a_number_container_system_mid

import (
	"container/heap"
	"testing"
)

type stk []int

func (s *stk) Len() int           { return len(*s) }
func (s *stk) Swap(i, j int)      { (*s)[i], (*s)[j] = (*s)[j], (*s)[i] }
func (s *stk) Less(i, j int) bool { return (*s)[i] < (*s)[j] }
func (s *stk) Push(x interface{}) { *s = append(*s, x.(int)) }
func (s *stk) Pop() interface{}   { old := *s; x := old[len(old)-1]; *s = old[:len(old)-1]; return x }
func (s *stk) Peak() int          { return (*s)[0] }

type NumberContainers struct {
	idxPQ map[int]*stk // [number] stk{idx...}
	idxMP map[int]int  // [idx]number
}

func Constructor() NumberContainers {
	return NumberContainers{
		idxPQ: make(map[int]*stk),
		idxMP: make(map[int]int),
	}
}

func (p *NumberContainers) Change(index int, number int) {
	p.idxMP[index] = number
	if _, has := p.idxPQ[number]; !has {
		p.idxPQ[number] = &stk{}
	}
	heap.Push(p.idxPQ[number], index)
}

func (p *NumberContainers) Find(number int) int {
	indices, has := p.idxPQ[number]
	if !has {
		return -1
	}
	for indices.Len() > 0 && p.idxMP[indices.Peak()] != number {
		heap.Pop(indices)
	}
	if indices.Len() == 0 {
		return -1
	}

	return indices.Peak()
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

func Test_design_a_number_container_system(t *testing.T) {
	tests := []struct {
		comd1 []string
		comd2 [][]int
	}{
		{comd1: []string{"NumberContainers", "find", "change", "change",
			"change", "change", "find", "change", "find"},
			comd2: [][]int{{}, {10}, {2, 10}, {1, 10}, {3, 10}, {5, 10},
				{10}, {1, 20}, {10}},
		},
		{comd1: []string{"NumberContainers", "change", "change", "find", "find", "find",
			"change", "find", "find", "change", "find", "change", "change", "change", "find",
			"find", "change", "find", "change", "change", "change"},
			comd2: [][]int{{}, {25, 50}, {56, 31}, {50}, {50}, {43}, {30, 50}, {31}, {43}, {25, 20},
				{50}, {56, 43}, {68, 31}, {56, 31}, {20}, {43}, {25, 43}, {43}, {56, 31}, {54, 43}, {63, 43}},
		},
	}
	for _, test := range tests {
		var sol = NumberContainers{}
		for i, c := range test.comd1 {
			switch c {
			case "NumberContainers":
				sol = Constructor()
			case "find":
				t.Log(sol.Find(test.comd2[i][0]))
			case "change":
				sol.Change(test.comd2[i][0], test.comd2[i][1])
			}
		}
	}
}
