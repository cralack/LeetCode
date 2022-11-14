package splitarraywithsameaverage

import "testing"

func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	for i, val := range nums {
		nums[i] = val*n - sum
	}

	m := n >> 1

	vis := map[int]bool{}
	for i := 1; i < 1<<m; i++ {
		t := 0
		for j, v := range nums[:m] {
			if (i >> j & 1) == 1 {
				t += v
			}
		}
		if t == 0 {
			return true
		}
		vis[t] = true
	}
	for i := 1; i < 1<<(n-m); i++ {
		t := 0
		for j, v := range nums[m:] {
			if (i >> j & 1) == 1 {
				t += v
			}
		}
		if t == 0 || (i != (1<<(n-m))-1 && vis[-t]) {
			return true
		}
	}
	return false
}

func Test_split_array(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	t.Log(splitArraySameAverage(nums))
}
