package snapshot_array_mid

import (
	"sort"
	"testing"
)

type pair struct{ snapID, val int }

type SnapshotArray struct {
	curIdx  int
	history map[int][]pair
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		history: make(map[int][]pair, length),
	}
}

func (a *SnapshotArray) Set(index int, val int) {
	a.history[index] = append(a.history[index], pair{snapID: a.curIdx, val: val})
}

func (a *SnapshotArray) Snap() int {
	a.curIdx++
	return a.curIdx - 1
}

func (a *SnapshotArray) Get(index int, snapId int) int {
	h := a.history[index]
	j := sort.Search(len(h), func(i int) bool {
		return h[i].snapID >= snapId+1
	}) - 1
	if j >= 0 {
		return h[j].val
	}
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
func Test_snapshot_array(t *testing.T) {
	tests := []struct {
		cmd1 []string
		cmd2 [][]int
	}{
		{
			[]string{"SnapshotArray", "set", "snap", "set", "get"},
			[][]int{{3}, {0, 5}, {}, {0, 6}, {0, 0}},
		}, {
			[]string{"SnapshotArray", "set", "snap", "snap", "snap", "get", "snap", "snap", "get"},
			[][]int{{1}, {0, 15}, {}, {}, {}, {0, 2}, {}, {}, {0, 0}},
		},
	}
	for _, tt := range tests {
		var sol SnapshotArray
		for i, c := range tt.cmd1 {
			switch c {
			case "SnapshotArray":
				sol = Constructor(tt.cmd2[i][0])
			case "set":
				sol.Set(tt.cmd2[i][0], tt.cmd2[i][1])
			case "snap":
				t.Log(sol.Snap())
			case "get":
				t.Log(sol.Get(tt.cmd2[i][0], tt.cmd2[i][1]))
			}
		}
	}
}
