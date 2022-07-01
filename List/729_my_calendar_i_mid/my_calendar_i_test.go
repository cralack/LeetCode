package mycalendarimid

import (
	"testing"
)

type MyCalendar struct {
	root *node
}

type node struct {
	start       int
	end         int
	left, right *node
}

func Constructor() MyCalendar {
	return MyCalendar{
		root: &node{-1, 0, nil, nil},
	}
}

func insert(start int, end int, root *node) bool {
	if end <= root.start {
		if root.left == nil {
			root.left = &node{start, end, nil, nil}
			return true
		}
		return insert(start, end, root.left)
	} else if start >= root.end {
		if root.right == nil {
			root.right = &node{start, end, nil, nil}
			return true
		}
		return insert(start, end, root.right)
	}
	return false
}

func (this *MyCalendar) Book(start int, end int) bool {
	return insert(start, end, this.root)
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

func Test_my_calendar_i(t *testing.T) {
	c1 := []string{"MyCalendar", "book", "book", "book"}
	c2 := [][]int{{-1}, {10, 20}, {15, 25}, {20, 30}}
	var sol MyCalendar
	for i, c := range c1 {
		switch c {
		case "MyCalendar":
			sol = Constructor()
		case "book":
			t.Log(sol.Book(c2[i][0], c2[i][1]))
		}
	}
}
