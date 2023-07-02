package maximumaveragepassratio

import (
	"container/heap"
	"testing"
)

func maxAverageRatio(classes [][]int, extraStudents int) (ans float64) {
	pq := make(cheap, 0)
	n := len(classes)
	// initi pq
	for _, cla := range classes {
		p, t := cla[0], cla[1]
		heap.Push(&pq, class{
			pass:  p,
			total: t,
			rate:  float64(p+1)/float64(t+1) - float64(p)/float64(t),
		})
	}
	// handle extraStu
	for i := 0; i < extraStudents; i++ {
		cur := heap.Pop(&pq).(class)
		cur.pass++
		cur.total++
		cur.rate = float64(cur.pass+1)/float64(cur.total+1) -
			float64(cur.pass)/float64(cur.total)
		heap.Push(&pq, cur)
	}
	// get ans
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(class)
		ans += float64(cur.pass) / float64(cur.total)
	}
	return ans / float64(n)
}

type class struct {
	pass, total int
	rate        float64
}
type cheap []class

func (h cheap) Len() int           { return len(h) }
func (h cheap) Less(i, j int) bool { return h[i].rate > h[j].rate }
func (h cheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *cheap) Push(x any)        { *h = append(*h, x.(class)) }
func (h *cheap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

var _ heap.Interface = &cheap{}

func Test_maximum_average_pass_ratio(t *testing.T) {
	classes := [][]int{{1, 2}, {3, 5}, {2, 2}}
	extraStudents := 2
	t.Log(maxAverageRatio(classes, extraStudents))
	classes = [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}
	extraStudents = 4
	t.Log(maxAverageRatio(classes, extraStudents))
}
