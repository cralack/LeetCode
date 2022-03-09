package nextgreaterelementi

import "testing"

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}
func (s *Stack) Empty() bool {
	return s.Len() <= 0
}
func (s *Stack) Push(x int) {
	*s = append(*s, x)
}
func (s *Stack) Pop() int {
	old := *s
	x := old[s.Len()-1]
	*s = old[0 : s.Len()-1]
	return x
}
func (s *Stack) Peek() int {
	old := *s
	x := old[s.Len()-1]
	return x
}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var s Stack
	res := make([]int, len(nums2))
	pos := make(map[int]int, len(nums2))
	//O(n2)
	//反向压入栈
	for i := len(nums2) - 1; i >= 0; i-- {
		// 判定个子高矮
		for !s.Empty() && s.Peek() <= nums2[i] {
			// 矮个起开，反正也被挡着了。
			s.Pop()
		}
		// nums[i] 身后的 next great number
		if s.Empty() {
			res[i] = -1
		} else {
			res[i] = s.Peek()
		}
		//针对此题对n2记录答案位置
		pos[nums2[i]] = i
		s.Push(nums2[i])
	}
	//O(n1)
	ans := make([]int, len(nums1))
	for idx, val := range nums1 {
		ans[idx] = res[pos[val]]
	}
	return ans
}
func Test_next_greater_element_i(t *testing.T) {
	n1 := []int{4, 1, 2}
	n2 := []int{1, 3, 4, 2}
	t.Log(nextGreaterElement(n1, n2))
	// n2 = []int{2, 1, 2, 4, 3}
	// t.Log(nextGreaterElement(n1, n2))
}
