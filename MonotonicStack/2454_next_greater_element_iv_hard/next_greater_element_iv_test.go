package next_greater_element_iv_hard

import (
	"testing"
)

func secondGreaterElement(nums []int) (ans []int) {
	ans = make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	s, t := make([]int, 0), make([]int, 0) // 存放nums中i、j下标的单调栈

	for i, x := range nums {
		for len(t) > 0 && nums[t[len(t)-1]] < x {
			ans[t[len(t)-1]] = x // t 栈顶的下下个更大元素是 x
			t = t[:len(t)-1]
		}
		j := len(s) - 1 // s 栈顶标
		for j >= 0 && nums[s[j]] < x {
			j-- // s 栈顶的下一个更大元素是 x
		}
		t = append(t, s[j+1:]...) // 把从 s 弹出的这一整段元素加到 t
		s = append(s[:j+1], i)    // 当前元素（的下标）加到 s 栈顶
	}

	return
}

func Test_next_greater_element_iv(t *testing.T) {
	tests := []struct{ nums []int }{
		{[]int{2, 4, 0, 9, 6}},
		{[]int{3, 3}},
	}
	for _, tt := range tests {
		t.Log(secondGreaterElement(tt.nums))
	}
}
