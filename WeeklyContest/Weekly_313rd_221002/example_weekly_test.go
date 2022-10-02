package weekly_contest

import (
	"strconv"
	"testing"
)

/************ 1st test************/
func commonFactors(a int, b int) (ans int) {
	if a < b { //keep a>b
		return commonFactors(b, a)
	}
	for i := 1; i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			ans++
		}
	}
	return
}
func Test_1st(t *testing.T) {
	a, b := 12, 6
	t.Log(commonFactors(a, b))
	a, b = 25, 30
	t.Log(commonFactors(a, b))
}

/************ 2nd test************/
func maxSum(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			sum := grid[i][j]                                     //core
			sum += grid[i-1][j-1] + grid[i-1][j] + grid[i-1][j+1] //top line
			sum += grid[i+1][j-1] + grid[i+1][j] + grid[i+1][j+1] //botom line
			if ans < sum {
				ans = sum
			}
		}
	}
	return
}
func Test_2nd(t *testing.T) {
	grid := [][]int{
		{6, 2, 1, 3},
		{4, 2, 1, 5},
		{9, 2, 8, 7},
		{4, 1, 2, 9}}
	t.Log(maxSum(grid))
	grid = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	t.Log(maxSum(grid))
}

/************ 3rd test************/
func minimizeXor(num1 int, num2 int) int {
	//cnt num2 bin.1
	x := num2
	cnt := 0
	for x > 0 {
		if x&1 == 1 {
			cnt++
		}
		x >>= 1
	}
	//replace num1  big bin.1 → small bin.1
	bn1 := convertToBin(num1)
	ans := make([]int, 0)
	for i := 0; i < len(bn1); i++ {
		if bn1[i] == '1' && cnt > 0 {
			ans = append(ans, 1)
			cnt--
		} else {
			ans = append(ans, 0)
		}
	}
	//if cnt still > 0
	for i := len(ans) - 1; i >= 0 && cnt > 0; i-- {
		if ans[i] == 0 {
			ans[i] = 1
			cnt--
		}
	}
	for i := 0; i < cnt; i++ {
		ans = append(ans, 1)
	}
	//turn bin to dec
	x = 0
	for i := 0; i < len(ans); i++ {
		if ans[i] == 1 {
			a := 1 << (len(ans) - i - 1)
			x |= a
		}
	}
	return x
}

// dec to bin
func convertToBin(num int) string {
	s := ""
	if num == 0 {
		return "0"
	}
	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return s
}
func Test_3rd(t *testing.T) {
	num1, num2 := 3, 5
	t.Log(minimizeXor(num1, num2))
	num1, num2 = 1, 12
	t.Log(minimizeXor(num1, num2))
	num1, num2 = 25, 72
	t.Log(minimizeXor(num1, num2))
	num1, num2 = 65, 84
	t.Log(minimizeXor(num1, num2))
}

/************ 4th test************/
func deleteString(s string) int {
	n := len(s)
	lcp := make([][]int, n+1) // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
	lcp[n] = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		lcp[i] = make([]int, n+1)
		for j := n - 1; j >= 0; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}
	f := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := 1; i+j*2 <= n; j++ {
			if lcp[i][i+j] >= j { // 说明 s[i:i+j] == s[i+j:i+j*2]
				f[i] = max(f[i], f[i+j])
			}
		}
		f[i]++
	}
	return f[0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func Test_4th(t *testing.T) {
	s := "abcabcdabc"
	t.Log(deleteString(s))
	s = "aaabaab"
	t.Log(deleteString(s))
	s = "aaaaa"
	t.Log(deleteString(s))
}
