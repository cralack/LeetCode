package implementstackusingqueues

import (
	"testing"
)

type MyStack struct {
	MainQue, SubQue []int
}

func Constructor() MyStack {
	return MyStack{
		MainQue: []int{},
		SubQue:  []int{},
	}
}

func (s *MyStack) Push(x int) {
	s.MainQue = append(s.MainQue, x)
}

func (s *MyStack) Pop() int {
	x := s.Top()
	s.MainQue = s.MainQue[1:]
	s.MainQue, s.SubQue = s.SubQue, s.MainQue
	return x
}

func (s *MyStack) Top() int {
	for len(s.MainQue) > 1 {
		//sub.Push(main.Pop)
		s.SubQue = append(s.SubQue, s.MainQue[0])
		s.MainQue = s.MainQue[1:]
	}
	x := s.MainQue[0]
	s.MainQue = s.MainQue[1:]
	s.MainQue = append(s.MainQue, x)
	return x
}

func (s *MyStack) Empty() bool {
	return len(s.MainQue) == 0 && len(s.SubQue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
func Test_implement_stack_using_queues(t *testing.T) {
	c1 := []string{"MyStack", "push", "push", "top", "pop", "empty"}
	c2 := []int{-1, 1, 2, -1, -1, -1}
	var que MyStack
	for i, c := range c1 {
		switch c {
		case "MyStack":
			que = Constructor()
		case "push":
			que.Push(c2[i])
		case "top":
			t.Log(que.Top())
		case "pop":
			t.Log(que.Pop())
		case "empty":
			t.Log(que.Empty())
		}
	}
}
