package numbersatmostngivendigitset

import (
	"strconv"
	"testing"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([]int, m)
	for i := range dp {
		dp[i] = -1 // dp[i] = -1 表示 i 这个状态还没被计算出来
	}
	var f func(int, bool, bool) int
	f = func(i int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum { // 如果填了数字，则为 1 种合法方案
				return 1
			}
			return
		}
		if !isLimit && isNum { // 在不受到任何约束的情况下，返回记录的结果，避免重复运算
			dv := &dp[i]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum { // 前面不填数字，那么可以跳过当前数位，也不填数字
			// isLimit 改为 false，因为没有填数字，位数都比 n 要短，自然不会受到 n 的约束
			// isNum 仍然为 false，因为没有填任何数字
			res += f(i+1, false, false)
		}
		// 根据是否受到约束，决定可以填的数字的上限
		up := byte('9')
		if isLimit {
			up = s[i]
		}
		// 注意：对于一般的题目而言，如果此时 isNum 为 false，则必须从 1 开始枚举，由于本题 digits 没有 0，所以无需处理这种情况
		for _, d := range digits { // 枚举要填入的数字 d
			if d[0] > up { // d 超过上限，由于 digits 是有序的，后面的 d 都会超过上限，故退出循环
				break
			}
			// isLimit：如果当前受到 n 的约束，且填的数字等于上限，那么后面仍然会受到 n 的约束
			// isNum 为 true，因为填了数字
			res += f(i+1, isLimit && d[0] == up, true)
		}
		return
	}
	return f(0, true, false)
}
func Test_mostNgivendigitset(t *testing.T) {
	digits, n := []string{"1", "3", "5", "7"}, 100
	t.Log(atMostNGivenDigitSet(digits, n))
}
