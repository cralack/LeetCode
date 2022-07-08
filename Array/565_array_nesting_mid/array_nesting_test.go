package arraynestingmid

import "testing"

func arrayNesting(nums []int) (ans int) {
	vis := make([]bool, len(nums))
	for i := range vis {
		cnt := 0
		for !vis[i] {
			vis[i] = true
			i = nums[i]
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return
}
func Test_array_nesting(t *testing.T) {
	arr := []int{5, 4, 0, 3, 1, 6, 2}
	t.Log(arrayNesting(arr))
}
