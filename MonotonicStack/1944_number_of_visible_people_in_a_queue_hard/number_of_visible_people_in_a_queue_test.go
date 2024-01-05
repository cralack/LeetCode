package number_of_visible_people_in_a_queue_hard

import (
	"testing"
)

func canSeePersonsCount(heights []int) (ans []int) {
	n := len(heights)
	stk := []int{1e6}
	ans = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for stk[len(stk)-1] < heights[i] {
			stk = stk[:len(stk)-1]
			ans[i]++
		}
		if len(stk) > 1 {
			ans[i]++
		}
		stk = append(stk, heights[i])
	}
	return
}

func Test_number_of_visible_people_in_a_queue(t *testing.T) {
	tests := []struct{ heights []int }{
		{heights: []int{10, 6, 8, 5, 11, 9}},
		{heights: []int{5, 1, 2, 3, 10}},
	}
	for _, tt := range tests {
		t.Log(canSeePersonsCount(tt.heights))
	}
}
