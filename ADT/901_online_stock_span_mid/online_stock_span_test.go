package onlinestockspan

import (
	"testing"
)

type StockSpanner struct {
	stack []*stock
}
type stock struct {
	val, day int
}

func Constructor() StockSpanner {
	return StockSpanner{stack: []*stock{}}
}

func (s *StockSpanner) Next(price int) int {
	cur := &stock{val: price, day: 1}
	for !s.IsEmpty() && s.Peak().val <= cur.val {
		tail := s.Pop()
		cur.day += tail.day
	}
	s.Push(cur)
	return s.Peak().day
}
func (s *StockSpanner) IsEmpty() bool {
	return len(s.stack) == 0
}
func (s *StockSpanner) Peak() *stock {
	return s.stack[len(s.stack)-1]
}
func (s *StockSpanner) Pop() *stock {
	tail := s.Peak()
	s.stack = s.stack[:len(s.stack)-1]
	return tail
}

func (s *StockSpanner) Push(x *stock) {
	s.stack = append(s.stack, x)
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

func Test_online_stock_span(t *testing.T) {
	cmd1 := []string{"StockSpanner", "next", "next", "next", "next", "next", "next", "next"}
	cmd2 := []int{-1, 100, 80, 60, 70, 60, 75, 85}
	var sol StockSpanner
	for i, c := range cmd1 {
		switch c {
		case "StockSpanner":
			sol = Constructor()
		case "next":
			t.Log(sol.Next(cmd2[i]))
		}
	}
}
