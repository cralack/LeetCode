package kthsmallestnumberinmultiplicationtable

import (
	"testing"
)

func findKthNumber(m int, n int, k int) int {
	getCnt := func(mid int) int {
		a, b := 0, 0
		for i := 1; i <= n; i++ {
			if i*m < mid { //若i*m<mid，整行都是小于mid的数，
				a += m //直接在a基础上累加m
			} else { //若i*m>=mid，根据mid是否存在于该行进行分情况讨论
				if mid%i == 0 && b+1 >= 0 { //mid能够被i整除，说明mid存在于该行
					a += mid/i - 1 //那么比mid小的数的个数为mid/i-1,累加到a
					b++            //同时对b进行加一
				} else { //mid不能够被i整除，说明mid不存在于该行
					a += mid / i //那么比mid小的数的个数为mid/i
				}
			}
		}
		return a + b
	}

	lo, hi := 1, m*n+1
	for lo < hi {
		mi := lo + (hi-lo)>>1
		cnt := getCnt(mi)
		if cnt >= k {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return hi
}
func Test_kth_smallest_number_in_multiplication_table(t *testing.T) {
	m, n, k := 3, 3, 5
	t.Log(findKthNumber(m, n, k))
	m, n, k = 2, 3, 6
	t.Log(findKthNumber(m, n, k))
}
