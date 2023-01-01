package weekly_contest

import (
	"strconv"
	"testing"
)

/************ 1st test************/
func countDigits(num int) int {
	cnt := 0
	tmp := num
	for tmp > 0 {
		cur := tmp % 10
		tmp /= 10
		if num%cur == 0 {
			cnt++
		}
	}
	return cnt
}
func Test_1st(t *testing.T) {
	num := 7
	t.Log(countDigits(num))
	num = 121
	t.Log(countDigits(num))
	num = 1248
	t.Log(countDigits(num))
}

/************ 2nd test************/

func distinctPrimeFactors(nums []int) int {
	n := 1000
	notPrime := make([]bool, n+1)
	for i := 2; i*i <= n; i++ { //只需要遍历 [2,sqrt(n)]
		if !notPrime[i] {
			for j := i * i; j <= n; j += i { //from i*i
				notPrime[j] = true
			}
		}
	}
	deparPrimeFactors := func(num int) (ans []int) {
		for i := 2; i <= num; i++ {
			if !notPrime[i] && num%i == 0 {
				ans = append(ans, i)
			}
			for num%i == 0 {
				num /= i
			}
		}
		return
	}

	cnt := make(map[int]int, 0)
	for _, n := range nums {
		primeFactors := deparPrimeFactors(n)
		for _, f := range primeFactors {
			cnt[f]++
		}
	}
	return len(cnt)
}
func Test_2nd(t *testing.T) {
	nums := []int{2, 4, 3, 7, 10, 6}
	t.Log(distinctPrimeFactors(nums))
	nums = []int{2, 4, 8, 16}
	t.Log(distinctPrimeFactors(nums))
}

/************ 3rd test************/
func minimumPartition(s string, k int) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		//最大可能len
		dp[i] = n + 1
	}

	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			cur, _ := strconv.Atoi(s[j:i])
			if cur <= k {
				dp[i] = min(dp[i], dp[j]+1)
			} else {
				break
			}
		}

	}
	if dp[n] == n+1 {
		return -1
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_3rd(t *testing.T) {
	s := "165462"
	k := 60
	t.Log(minimumPartition(s, k))
	s = "238182"
	k = 5
	t.Log(minimumPartition(s, k))
}

/************ 4th test************/
func closestPrimes(left int, right int) (ans []int) {
	n := right
	notPrime := make([]bool, n+1)
	primeArr := make([]int, 0, right-left)
	for i := 2; i*i <= n; i++ { //只需要遍历 [2,sqrt(n)]
		if !notPrime[i] {
			for j := i * i; j <= n; j += i { //from i*i
				notPrime[j] = true
			}
		}
	}
	for i := max(2, left); i <= right; i++ {
		if !notPrime[i] {
			primeArr = append(primeArr, i)
		}
	}
	if len(primeArr) < 2 {
		return []int{-1, -1}
	}
	dist := right - left + 1
	for i := 1; i < len(primeArr); i++ {
		cur, pre := primeArr[i], primeArr[i-1]
		if cur-pre < dist {
			dist = cur - pre
			ans = []int{pre, cur}
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_4th(t *testing.T) {
	left := 10
	right := 19
	t.Log(closestPrimes(left, right))
	left = 4
	right = 6
	t.Log(closestPrimes(left, right))
	left = 1
	right = 1000000
	t.Log(closestPrimes(left, right))
	left = 710119
	right = 710189
	t.Log(closestPrimes(left, right))
}
