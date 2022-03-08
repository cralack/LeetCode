package findmedianfromdatastream

import (
	"container/heap"
	"fmt"
	"testing"
)

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	old := *h
	x := old[h.Len()-1]
	*h = old[:h.Len()-1]
	return x
}
func (h *Heap) Peek() int {
	old := *h
	x := old[0]
	return x
}

type MedianFinder struct {
	smallQue, largeQue *Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		//小顶堆存大数(正数)
		smallQue: &Heap{},
		//大顶堆存小数(负数)
		largeQue: &Heap{},
	}
}

func (p *MedianFinder) AddNum(num int) {
	if p.smallQue.Len() >= p.largeQue.Len() {
		heap.Push(p.smallQue, num)
		heap.Push(p.largeQue, -heap.Pop(p.smallQue).(int))
	} else { //p.smallQue.Len() < p.largeQue.Len()
		heap.Push(p.largeQue, -num)
		heap.Push(p.smallQue, -heap.Pop(p.largeQue).(int))
	}
}

func (p *MedianFinder) FindMedian() float64 {
	if p.smallQue.Len() < p.largeQue.Len() {
		return -float64(p.largeQue.Peek())
	} /*else if p.smallQue.Len() > p.largeQue.Len() {
		return float64(p.smallQue.Peek())
	}*/
	return (float64(p.smallQue.Peek()) + float64(-p.largeQue.Peek())) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
func Test_find_median_from_data_stream(t *testing.T) {
	finder := Constructor()
	arr := []int{73, 15, 21, 19, 10, 31, 42, 65, 48}
	for _, i := range arr {
		finder.AddNum(i)
		res := finder.FindMedian()
		fmt.Println(res)
	}
}
