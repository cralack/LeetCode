package dailytemperatures

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
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	//stack存放索引
	var s Stack
	for i := n - 1; i >= 0; i-- {
		for !s.Empty() && temperatures[s.Peek()] <= temperatures[i] {
			s.Pop()
		}
		if s.Empty() {
			res[i] = 0
		} else {
			res[i] = s.Peek() - i
		}
		s.Push(i)
	}
	return res
}
func Test_daily_temperatures(t *testing.T) {
	tem := []int{73, 74, 75, 71, 69, 72, 76, 73}
	t.Log(dailyTemperatures(tem))
	tem = []int{30, 40, 50, 60}
	t.Log(dailyTemperatures(tem))
	tem = []int{30, 60, 90}
	t.Log(dailyTemperatures(tem))
}
