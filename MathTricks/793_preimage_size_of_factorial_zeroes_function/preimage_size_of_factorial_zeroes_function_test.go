package preimagesizeoffactorialzeroesfunction

import (
	"math"
	"testing"
)

func preimageSizeFZF(k int) (res int) {
	res = right_bound(k) - left_bound(k) + 1
	return res
}
func left_bound(target int) int {
	lo, hi := 0, math.MaxInt
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if trailingZeroes(mi) >= target {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo
}
func right_bound(target int) int {
	lo, hi := 0, math.MaxInt
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if trailingZeroes(mi) <= target {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return lo - 1
}

func trailingZeroes(n int) (res int) {
	divisor := 5
	for divisor <= n {
		res += n / divisor
		divisor *= 5
	}
	return res
}
func Test_preimage_size_of_factorial_zeroes_function(t *testing.T) {
	t.Log(preimageSizeFZF(0))
	t.Log(preimageSizeFZF(5))
	t.Log(preimageSizeFZF(3))
}
