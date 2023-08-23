package defusethebomb

import "testing"

func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k == 0 {
		return ans
	} else {
		preSum := make([]int, n*2)
		preSum[0] = code[0]
		for i := 1; i < 2*n; i++ {
			preSum[i] = preSum[i-1] + code[i%n]
		}
		if k > 0 {
			for i := 0; i < n; i++ {
				ans[i] = preSum[i+k] - preSum[i]
			}
		} else {
			for i := 0; i < n; i++ {
				ans[i] = preSum[i+n-1] - preSum[i+n+k-1]
			}
		}
	}
	return ans
}
func Test_defuse_the_bomb(t *testing.T) {
	code := []int{5, 7, 1, 4}
	k := 3
	t.Log(decrypt(code, k))
	code = []int{1, 2, 3, 4}
	k = 0
	t.Log(decrypt(code, k))
	code = []int{2, 4, 9, 3}
	k = -2
	t.Log(decrypt(code, k))
}
