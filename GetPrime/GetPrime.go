package GetPrime

func PrimeArr(n int) []int {
	//get a prime array and max element <=n
	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}

	for i := 2; i <= n; i++ {
		if prime[i] {
			for j := 2 * i; j <= n; j += i {
				prime[j] = false
			}
		}
	}

	ans := []int{}
	for idx, val := range prime {
		if val {
			ans = append(ans, idx)
		}
	}
	return ans
}

func PrimeNumber(n int) int {
	arr := PrimeArr(n)
	return arr[n]
}
