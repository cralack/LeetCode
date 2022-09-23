package designlinkedlistmid

import "testing"

type MyLinkedList struct {
	Head, Tail *Node
	Size       int
}
type Node struct {
	Prev, Next *Node
	Data       int
}

func Constructor() MyLinkedList {
	head, tail := &Node{nil, nil, -1}, &Node{nil, nil, -2}
	head.Next = tail
	tail.Prev = head
	return MyLinkedList{
		Head: head,
		Tail: tail,
		Size: 0,
	}
}

func (this *MyLinkedList) find(idx int) *Node {
	if this.Size <= idx || idx < 0 {
		return nil
	}
	var cur *Node
	if idx <= this.Size/2 {
		cur = this.Head
		for i := 0; i <= idx; i++ {
			cur = cur.Next
		}
	} else {
		cur = this.Tail
		for i := this.Size - 1; i >= idx; i-- {
			cur = cur.Prev
		}
	}
	return cur
}

func (this *MyLinkedList) Get(index int) int {
	node := this.find(index)
	if node == nil {
		return -1
	}
	return node.Data
}

func (this *MyLinkedList) AddAtHead(val int) {
	tmp := &Node{Data: val}
	tmp.Prev = this.Head
	tmp.Next = this.Head.Next
	this.Head.Next.Prev = tmp
	this.Head.Next = tmp
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	tmp := &Node{Data: val}
	tmp.Prev = this.Tail.Prev
	tmp.Next = this.Tail
	this.Tail.Prev.Next = tmp
	this.Tail.Prev = tmp
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	if index > this.Size {
		return
	}
	cur := this.find(index)
	tmp := &Node{Data: val}
	tmp.Next = cur
	tmp.Prev = cur.Prev
	cur.Prev.Next = tmp
	cur.Prev = tmp
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.find(index)
	if cur != nil {
		cur.Next.Prev = cur.Prev
		cur.Prev.Next = cur.Next
		this.Size--
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func Test_designlist(t *testing.T) {
	var sol MyLinkedList
	foo := func(c1 []string, c2 [][]int) {
		for i, c := range c1 {
			switch c {
			case "MyLinkedList":
				sol = Constructor()
			case "addAtHead":
				sol.AddAtHead(c2[i][0])
			case "addAtTail":
				sol.AddAtTail(c2[i][0])
			case "addAtIndex":
				sol.AddAtIndex(c2[i][0], c2[i][1])
			case "get":
				t.Log(sol.Get(c2[i][0]))
			case "deleteAtIndex":
				sol.DeleteAtIndex(c2[i][0])
			}
		}
	}
	com1 := []string{"MyLinkedList", "addAtHead", "addAtTail",
		"addAtIndex", "get", "deleteAtIndex", "get"}
	com2 := [][]int{{-1}, {1}, {3}, {1, 2}, {1}, {1}, {1}}
	foo(com1, com2)

	com1 = []string{"MyLinkedList", "addAtHead", "get", "addAtHead",
		"addAtHead", "deleteAtIndex", "addAtHead", "get", "get", "get",
		"addAtHead", "deleteAtIndex"}
	com2 = [][]int{{-1}, {4}, {1}, {1}, {5}, {3}, {7}, {3}, {3}, {3}, {1}, {4}}
	foo(com1, com2)

	com1 = []string{"MyLinkedList", "addAtHead", "addAtHead", "addAtHead",
		"addAtIndex", "deleteAtIndex", "addAtHead", "addAtTail", "get",
		"addAtHead", "addAtIndex", "addAtHead"}
	com2 = [][]int{{-1}, {7}, {2}, {1}, {3, 0}, {2}, {6}, {4}, {4}, {4}, {5, 0}, {6}}
	foo(com1, com2)
}
