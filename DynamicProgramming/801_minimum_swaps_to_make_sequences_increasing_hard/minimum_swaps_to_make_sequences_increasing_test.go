package minimumswapstomakesequencesincreasing

import "testing"

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
		dp[i][0] = n + 10
		dp[i][1] = n + 10
	}
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		}
		if nums2[i-1] < nums1[i] && nums1[i-1] < nums2[i] {
			dp[i][0] = min(dp[i][0], dp[i-1][1])
			dp[i][1] = min(dp[i][1], dp[i-1][0]+1)
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}
func minSwap_spaceShrink(nums1 []int, nums2 []int) int {
	n := len(nums1)
	a, b := 0, 1
	for i := 1; i < n; i++ {
		at, bt := a, b
		a, b = n, n
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			a = at
			b = bt + 1
		}
		if nums2[i-1] < nums1[i] && nums1[i-1] < nums2[i] {
			a = min(a, bt)
			b = min(b, at+1)
		}
	}
	return min(a, b)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test_minSwap(t *testing.T) {
	nums1 := []int{1, 3, 5, 4}
	nums2 := []int{1, 2, 3, 7}
	t.Log(minSwap(nums1, nums2))
	nums1 = []int{0, 3, 5, 8, 9}
	nums2 = []int{2, 1, 4, 6, 9}
	t.Log(minSwap(nums1, nums2))
}
