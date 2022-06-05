package mycalendariii

import (
	"testing"
)

/* ******** diff arr ************
type MyCalendarThree struct {
	k    int
	memo map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{0, map[int]int{}}
}

func (p *MyCalendarThree) Book(start int, end int) int {
	p.memo[start]++
	p.memo[end]--
	cnt := 0
	arr := make([]int, 0, len(p.memo))
	for key := range p.memo {
		arr = append(arr, key)
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	for _, key := range arr {
		cnt += p.memo[key]
		if cnt > p.k {
			p.k = cnt
		}
	}
	return p.k
}
*/

type pair struct{ num, lazy int }

type MyCalendarThree map[int]pair

func Constructor() MyCalendarThree {
	return MyCalendarThree{}
}

func (t MyCalendarThree) update(start, end, l, r, idx int) {
	if r < start || end < l {
		return
	}
	if start <= l && r <= end {
		p := t[idx]
		p.num++
		p.lazy++
		t[idx] = p
	} else {
		mid := (l + r) / 2
		t.update(start, end, l, mid, idx*2)
		t.update(start, end, mid+1, r, idx*2+1)
		p := t[idx]
		p.num = p.lazy + max(t[idx*2].num, t[idx*2+1].num)
		t[idx] = p
	}
}

func (t MyCalendarThree) Book(start, end int) int {
	t.update(start, end-1, 0, 1e9, 1)
	return t[1].num
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
func Test_my_calendar_iii(t *testing.T) {
	c1 := []string{"MyCalendarThree", "book", "book", "book", "book", "book", "book"}
	c2 := [][]int{{}, {10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}}
	var sol MyCalendarThree
	for i, c := range c1 {
		switch c {
		case "MyCalendarThree":
			sol = Constructor()
		case "book":
			t.Log(sol.Book(c2[i][0], c2[i][1]))
		}
	}
}
