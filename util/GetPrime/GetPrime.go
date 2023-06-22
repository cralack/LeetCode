package GetPrime

//get a prime array and max element <=n
func PrimeArr(n int) []int {
	notPrime := make([]bool, n)
	for i := 2; i*i < n; i++ { //只需要遍历 [2,sqrt(n)]
		if !notPrime[i] {
			for j := i * i; j < n; j += i { //from i*i
				notPrime[j] = true
			}
		}
	}

	ans := make([]int, n/2)
	cnt := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			ans[cnt] = i
			cnt++
		}
	}
	ans = ans[:cnt]
	return ans
}

func PrimeNumber(n int) int {
	arr := PrimeArr(n)
	return arr[n]
}
