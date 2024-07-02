package maximum_prime_difference_mid

import (
	"testing"
)

var notPrime []bool

func init() {
	var n = 101
	notPrime = make([]bool, n)
	for i := 2; i*i < n; i++ {
		if !notPrime[i] {
			for j := i * i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}
	notPrime[1] = true
}

func maximumPrimeDifference(nums []int) int {
	mn, mx := len(nums)-1, 0
	for i, n := range nums {
		if !notPrime[n] {
			mn = min(mn, i)
			mx = max(mx, i)
		}
	}
	return mx - mn
}

func Test_maximum_prime_difference(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{nums: []int{4, 2, 9, 5, 3}},
		{nums: []int{4, 8, 2, 8}},
	}
	for _, tt := range tests {
		t.Log(maximumPrimeDifference(tt.nums))
	}
}
