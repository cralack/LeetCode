package longestconsecutivesequence

import "testing"

func longestConsecutive(nums []int) int {
	set := make(map[int]bool, 0)
	for _, n := range nums {
		if _, has := set[n]; !has {
			set[n] = true
		}
	}
	max := 0
	for num := range set {
		if !set[num-1] {
			curNum := num
			cnt := 1
			for set[curNum+1] {
				curNum++
				cnt++
			}
			if max < cnt {
				max = cnt
			}
		}
	}
	return max
}
func Test_longest_consecutive_sequence(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	t.Log(longestConsecutive(nums))
	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	t.Log(longestConsecutive(nums))
}
