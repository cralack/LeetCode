package primearrangementseasy

import "testing"

const mod int = 1e9 + 7

func numPrimeArrangements(n int) int {
	arr := primeArr(n)
	primeCnt := 0
	for i := 2; i <= n; i++ {
		if !arr[i] {
			primeCnt++
		}
	}
	return factorial(n-primeCnt) * factorial(primeCnt) % mod
}
func primeArr(n int) []bool {
	notPrime := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !notPrime[i] {
			for j := i * i; j <= n; j += i {
				notPrime[j] = true
			}
		}
	}
	return notPrime
}
func factorial(n int) (ans int) {
	ans = 1
	for i := 1; i <= n; i++ {
		ans = ans * i % mod
	}
	return ans
}
func Test_prime_arrangements(t *testing.T) {
	n := 5
	t.Log(numPrimeArrangements(n))
	n = 100
	t.Log(numPrimeArrangements(n))
}
