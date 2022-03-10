package implementqueueusingstacks

import "testing"

type MyQueue struct {
	outStack, inStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		outStack: []int{},
		inStack:  []int{},
	}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	x := q.Peek()
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
			q.inStack = q.inStack[:len(q.inStack)-1]
		}
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.outStack) == 0 && len(q.inStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
func Test_implement_queue_using_stacks(t *testing.T) {
	c1 := []string{"MyQueue", "push", "push", "peek", "pop", "empty"}
	c2 := []int{-1, 1, 2, -1, -1, -1}
	var que MyQueue
	for i, c := range c1 {
		switch c {
		case "MyQueue":
			que = Constructor()
		case "push":
			que.Push(c2[i])
		case "peek":
			t.Log(que.Peek())
		case "pop":
			t.Log(que.Pop())
		case "empty":
			t.Log(que.Empty())
		}
	}
}
