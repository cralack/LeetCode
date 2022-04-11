package superpow

import "testing"

const base = 1337

func superPow(a int, b []int) int {
	// 递归的 base case
	if len(b) == 0 {
		return 1
	} // 取出最后一个数
	last := b[len(b)-1]
	b = b[:len(b)-1]
	// 将原问题化简，缩小规模递归求解
	part1 := myPow(a, last)
	part2 := myPow(superPow(a, b), 10)
	// 每次乘法都要求模,合并出结果
	return (part1 * part2) % base
}

// 计算 a 的 k 次方然后与 base 求模的结果
func myPow(a, k int) int {
	if k == 0 {
		return 1
	}
	// 对因子求模
	a %= base
	// res := 1
	// for i := 0; i < k; i++ {
	// 	// 这里有乘法，是潜在的溢出点
	// 	res *= a
	// 	// 对乘法结果求模
	// 	res %= base
	// }
	// return res
	if k%2 == 1 { // k 是奇数
		return (a * myPow(a, k-1)) % base
	} else { // k 是偶数
		sub := myPow(a, k/2)
		return (sub * sub) % base
	}
}
func Test_super_pow(t *testing.T) {
	a, b := 2, []int{3}
	t.Log(superPow(a, b))
	a, b = 2, []int{1, 0}
	t.Log(superPow(a, b))
	a, b = 1, []int{4, 3, 3, 8, 5, 2}
	t.Log(superPow(a, b))
	a, b = 2147483647, []int{2, 0, 0}
	t.Log(superPow(a, b))
}
