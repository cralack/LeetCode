package weekly_contest

import (
	"sort"
	"testing"
)

/************ 1st test************/
func countDigits(num int) (ans int) {
	for cur := num; cur > 0; cur /= 10 {
		if num%(cur%10) == 0 {
			ans++
		}
	}
	return
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
	set := make(map[int]struct{}, 0)
	for _, n := range nums {
		// 枚举n的质因数
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				set[i] = struct{}{}
				for n /= i; n%i == 0; n /= i {
				}
			}
		}
		if n > 1 {
			set[n] = struct{}{}
		}
	}
	return len(set)
}
func Test_2nd(t *testing.T) {
	nums := []int{2, 4, 3, 7, 10, 6}
	t.Log(distinctPrimeFactors(nums))
	nums = []int{2, 4, 8, 16}
	t.Log(distinctPrimeFactors(nums))
}

/************ 3rd test************/
func minimumPartition(s string, k int) (ans int) {
	curNum := 0
	for _, char := range s {
		v := int(char - '0')
		if v > k {
			return -1
		}
		curNum = curNum*10 + v
		if curNum > k {
			ans++
			curNum = v
		}
	}
	return ans + 1
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
const mx = 1000033

var primes = make([]int, 0, mx)

func init() {
	notPrime := make([]bool, mx+1)
	for i := 2; i <= mx; i++ {
		if !notPrime[i] {
			// 预处理出primes
			primes = append(primes, i)
			for j := i * i; j <= mx; j += i {
				notPrime[j] = true
			}
		}
	}
}

func closestPrimes(left int, right int) (ans []int) {
	ans = []int{-1, -1}
	maxGap := right
	for i := sort.SearchInts(primes, left); primes[i+1] <= right; i++ {
		cur, pre := primes[i+1], primes[i]
		gap := cur - pre
		if gap < maxGap {
			maxGap = gap
			ans[0], ans[1] = pre, cur
			// 除了[2,3]其他都是[x,x+2]
			if gap <= 2 {
				break
			}
		}
	}
	return
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
