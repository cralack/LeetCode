package design_memory_allocator_mid

import (
	"testing"

	"github.com/emirpasic/gods/trees/redblacktree"
)

type Allocator struct {
	rbt *redblacktree.Tree
	dic map[int][]int
}

func Constructor(n int) Allocator {
	rbt := redblacktree.NewWithIntComparator()
	rbt.Put(-1, -1)
	rbt.Put(n, n)
	return Allocator{
		rbt: rbt,
		dic: make(map[int][]int),
	}
}

func (a *Allocator) Allocate(size int, mID int) int {
	start := -1
	it := a.rbt.Iterator()
	for it.Next() {
		curIt := it.Key().(int)
		if start != -1 {
			end := curIt - 1
			if end-start+1 >= size {
				a.rbt.Put(start, start+size-1)
				a.dic[mID] = append(a.dic[mID], start)
				return start
			}
		}
		start = it.Value().(int) + 1
	}
	return -1
}

func (a *Allocator) FreeMemory(mID int) (ans int) {
	for _, start := range a.dic[mID] {
		if end, ok := a.rbt.Get(start); ok {
			a.rbt.Remove(start)
			ans += end.(int) - start + 1
		}
	}
	a.dic[mID] = []int{}
	return
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */

func Test_design_memory_allocator(t *testing.T) {
	tests := []struct {
		cmd1 []string
		cmd2 [][]int
	}{
		{
			cmd1: []string{"Allocator", "allocate", "allocate", "allocate", "freeMemory", "allocate",
				"allocate", "allocate", "freeMemory", "allocate", "freeMemory"},
			cmd2: [][]int{{10}, {1, 1}, {1, 2}, {1, 3}, {2}, {3, 4}, {1, 1}, {1, 1}, {1}, {10, 2}, {7}},
		},
	}
	for _, tt := range tests {
		var sol Allocator
		for i, c := range tt.cmd1 {
			switch c {
			case "Allocator":
				sol = Constructor(tt.cmd2[i][0])
			case "allocate":
				t.Log(sol.Allocate(tt.cmd2[i][0], tt.cmd2[i][1]))
			case "freeMemory":
				t.Log(sol.FreeMemory(tt.cmd2[i][0]))
			}
		}
	}
}
