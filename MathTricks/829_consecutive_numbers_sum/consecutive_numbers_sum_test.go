package consecutivenumberssum

import "testing"

func consecutiveNumbersSum_slidewindow(n int) int {
	//TLE
	sum, cnt := 0, 0
	left, right := 1, 1
	for right <= n+1 { //sum=[left,right)
		if sum < n {
			sum += right
			right++
		} else if sum >= n {
			if sum == n {
				cnt++
			}
			sum -= left
			left++
		}
	}
	return cnt
}
func consecutiveNumbersSum(n int) (cnt int) {
	n *= 2
	for k := 1; k*k < n; k++ {
		if n%k != 0 {
			continue
		}
		if (n/k-(k-1))%2 == 0 {
			cnt++
		}
	}
	return
}
func Test_consecutive_numbers_sum(t *testing.T) {
	n := 5
	t.Log(consecutiveNumbersSum(n))
	n = 9
	t.Log(consecutiveNumbersSum(n))
	n = 15
	t.Log(consecutiveNumbersSum(n))
}
