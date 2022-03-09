package nextgreaterelementii

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
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	n := len(nums)
	var s Stack
	//双倍遍历数组模拟循环
	for i := 2*n - 1; i >= 0; i-- {
		//索引求模
		for !s.Empty() && s.Peek() <= nums[i%n] {
			s.Pop()
		}
		if s.Empty() {
			res[i%n] = -1
		} else {
			res[i%n] = s.Peek()
		}
		s.Push(nums[i%n])
	}
	return res
}
func Test_next_greater_element_ii(t *testing.T) {
	nums := []int{1, 2, 1}
	t.Log(nextGreaterElements(nums))
	nums = []int{1, 2, 3, 4, 3}
	t.Log(nextGreaterElements(nums))
}
