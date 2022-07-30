package designcircularqueuemid

import "testing"

type MyCircularQueue struct {
	Data []int
	Ft   int
	Rr   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k+1), 0, 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Data[this.Rr] = value
	this.Rr = this.Nxt(this.Rr)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Ft = this.Nxt(this.Ft)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Data[this.Ft]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Data[(this.Rr-1+len(this.Data))%len(this.Data)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Rr == this.Ft
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Nxt(this.Rr) == this.Ft
}

func (this *MyCircularQueue) Nxt(cur int) int {
	return (cur + 1) % len(this.Data)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
func Test_design_circular_queue(t *testing.T) {
	var sol MyCircularQueue
	sol = Constructor(3)
	t.Log(sol.EnQueue(1))
	t.Log(sol.EnQueue(2))
	t.Log(sol.EnQueue(3))
	t.Log(sol.EnQueue(4))
	t.Log(sol.Rear())
	t.Log(sol.IsFull())
	t.Log(sol.DeQueue())
	t.Log(sol.EnQueue(4))
	t.Log(sol.Rear())
}
