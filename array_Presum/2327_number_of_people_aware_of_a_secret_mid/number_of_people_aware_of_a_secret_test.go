package number_of_people_aware_of_a_secret_mid

import (
	"testing"
)

func peopleAwareOfSecret(n, delay, forget int) (ans int) {
	const mod = 1e9 + 7
	// known[i] 表示恰好在第 i 天得知秘密的人数
	known := make([]int, n+1)
	known[1] = 1

	for i := 1; i <= n; i++ {
		known[i] %= mod
		// 统计在第 n 天没有忘记秘密的人数
		if i+forget-1 >= n {
			ans += known[i]
		}
		// 恰好在第 i 天得知秘密的人，会在第 [i+delay, i+forget-1] 天分享秘密
		for j := i + delay; j <= min(i+forget-1, n); j++ {
			known[j] += known[i]
		}
	}

	return ans % mod
}

func Test_number_of_people_aware_of_a_secret(t *testing.T) {
	tests := []struct {
		n, delay, forget int
	}{
		{n: 6, delay: 2, forget: 4},
		{n: 4, delay: 1, forget: 3},
	}
	for _, test := range tests {
		t.Log(peopleAwareOfSecret(test.n, test.delay, test.forget))
	}
}
