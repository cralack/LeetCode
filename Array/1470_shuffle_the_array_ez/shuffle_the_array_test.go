package shufflethearrayez

import "testing"

func shuffle(nums []int, n int) []int {
	ans := make([]int, n*2)
	for i := 0; i < n; i++ {
		ans[2*i] = nums[i]
		ans[2*i+1] = nums[n+i]
	}
	return ans
}
func Test_shuffle_the_array(t *testing.T) {
	arr := []int{2, 5, 1, 3, 4, 7}
	n := 3
	t.Log(shuffle(arr, n))
}
