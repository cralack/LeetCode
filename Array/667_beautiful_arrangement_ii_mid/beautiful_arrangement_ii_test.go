package beautifularrangementiimid

import "testing"

func constructArray(n int, k int) (ans []int) {
	ans = make([]int, n)
	i := 0
	//p从小到大 q从大到小
	p, q := 1, n
	//构造前k个数组 k-1个不同的整数
	//奇数位从大到小，偶数位从小到大
	for j := 0; j < k; j++ {
		if j%2 == 0 {
			ans[i] = p
			p++
		} else {
			ans[i] = q
			q--
		}
		i++
	}
	//构造剩下的绝对值为1的序列
	if k%2 == 0 { //偶数时，降序
		for i < n {
			ans[i] = q
			q--
			i++
		}
	} else { //奇数时，升序
		for i < n {
			ans[i] = p
			p++
			i++
		}
	}
	return
}
func Test_beautiful_arrangement_ii(t *testing.T) {
	n, k := 3, 1
	t.Log(constructArray(n, k))
	n, k = 3, 2
	t.Log(constructArray(n, k))
	n, k = 6, 5
	t.Log(constructArray(n, k))
	n, k = 6, 4
	t.Log(constructArray(n, k))
	n, k = 6, 3
	t.Log(constructArray(n, k))
	n, k = 6, 2
	t.Log(constructArray(n, k))
}
