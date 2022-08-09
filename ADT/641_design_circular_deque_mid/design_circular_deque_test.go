package designcirculardequemid

import "testing"

type MyCircularDeque struct {
	Front, Rear int
	Cache       []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		Cache: make([]int, k+1),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Front = (this.Front + len(this.Cache) - 1) % len(this.Cache)
	this.Cache[this.Front] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Cache[this.Rear] = value
	this.Rear = (this.Rear + 1) % len(this.Cache)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.Front = (this.Front + 1) % len(this.Cache)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.Rear = (this.Rear + len(this.Cache) - 1) % len(this.Cache)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Cache[this.Front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Cache[(this.Rear+len(this.Cache)-1)%len(this.Cache)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Front == this.Rear
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.Rear+1)%len(this.Cache) == this.Front
}

func Test_design_circular_deque(t *testing.T) {
	c1 := []string{"MyCircularDeque", "insertLast", "insertLast",
		"insertFront", "insertFront", "getRear", "isFull", "deleteLast",
		"insertFront", "getFront"}
	c2 := []int{3, 1, 2, 3, 4, -1, -1, -1, 4, -1}
	var sol MyCircularDeque
	for i, c := range c1 {
		switch c {
		case "MyCircularDeque":
			sol = Constructor(c2[i])
			t.Log("nil")
		case "insertFront":
			t.Log(sol.InsertFront(c2[i]))
		case "insertLast":
			t.Log(sol.InsertLast(c2[i]))
		case "deleteFront":
			t.Log(sol.DeleteFront())
		case "deleteLast":
			t.Log(sol.DeleteLast())
		case "getFront":
			t.Log(sol.GetFront())
		case "getRear":
			t.Log(sol.GetRear())
		case "isEmpty":
			t.Log(sol.IsEmpty())
		case "isFull":
			t.Log(sol.IsFull())
		}
	}
}
