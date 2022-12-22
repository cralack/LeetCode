package maximizescoreafternoperations

import (
	"math/bits"
	"testing"
)

func maxScore(nums []int) int {
	m := len(nums)
	g := make([][]int, m)
	for i := 0; i < m; i++ {
		g[i] = make([]int, m)
		for j := i + 1; j < m; j++ {
			g[i][j] = gcd(nums[i], nums[j])
		}
	}
	//假设 m 为数组 nums 中的元素个数，那么状态一共有 2^m种
	dp := make([]int, 1<<m)
	for k := 0; k < 1<<m; k++ {
		//先判断此状态中二进制位的个数 cntcnt 是否为偶数
		cnt := bits.OnesCount(uint(k))
		if cnt%2 == 0 {
			//枚举 k 中二进制位为 1 的位置
			for i := 0; i < m; i++ {
				if k>>i&1 == 1 {
					for j := i + 1; j < m; j++ {
						if k>>j&1 == 1 {
							//1<<i,1<<j为分别枚举出的i,j对应2进制位置
							// k^(1<<i)^(1<<j)即为未取gcd(i,j)之前的状态
							dp[k] = max(dp[k], dp[k^(1<<i)^(1<<j)]+
								cnt/2*g[i][j])
						}
					}
				}
			}
		}
	}
	return dp[1<<m-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func Test_maximize_score_after_n_operations(t *testing.T) {
	nums := []int{1, 2}
	t.Log(maxScore(nums))
	nums = []int{3, 4, 6, 8}
	t.Log(maxScore(nums))
	nums = []int{1, 2, 3, 4, 5, 6}
	t.Log(maxScore(nums))
}
