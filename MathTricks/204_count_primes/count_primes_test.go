package countprimes

import "testing"

func countPrimes(n int) (cnt int) {
	notPrime := make([]bool, n)
	for i := 2; i*i < n; i++ { // 只需要遍历 [2,sqrt(n)]
		if !notPrime[i] {
			for j := i * i; j < n; j += i { // from i*i
				notPrime[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			cnt++
		}
	}
	return cnt
}
func Test_count_primes(t *testing.T) {
	t.Log(countPrimes(10))
}
