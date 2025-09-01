package maximumaveragepassratiomid

import (
	"container/heap"
	"testing"
)

type (
	class struct {
		pass  int
		total int
		ratio float64
	}
	hp []class
)

func (h hp) Len() int           { return len(h) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool { return h[i].ratio > h[j].ratio }
func (h *hp) Push(x any)        { *h = append(*h, x.(class)) }
func (h *hp) Pop() any          { old := *h; x := old[h.Len()-1]; *h = old[:h.Len()-1]; return x }

func maxAverageRatio(classes [][]int, extraStudents int) (ans float64) {
	chp := hp{}
	for _, c := range classes {
		p, t := c[0], c[1]
		heap.Push(&chp, class{
			p, t, float64(p+1)/float64(t+1) - float64(p)/float64(t),
		})
	}
	for i := 0; i < extraStudents; i++ {
		cur := heap.Pop(&chp).(class)
		cur.pass++
		cur.total++
		cur.ratio = float64(cur.pass+1)/float64(cur.total+1) - float64(cur.pass)/float64(cur.total)
		heap.Push(&chp, cur)
	}

	for chp.Len() > 0 {
		cur := heap.Pop(&chp).(class)
		ans += float64(cur.pass) / float64(cur.total)
	}

	return ans / float64(len(classes))
}

func Test_maximum_average_pass_ratio(t *testing.T) {
	tests := []struct {
		classes       [][]int
		extraStudents int
	}{
		{classes: [][]int{{1, 2}, {3, 5}, {2, 2}}, extraStudents: 2},
		{classes: [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, extraStudents: 4},
	}
	for _, tt := range tests {
		avgRatio := maxAverageRatio(tt.classes, tt.extraStudents)
		t.Logf("%f", avgRatio)
	}
}
