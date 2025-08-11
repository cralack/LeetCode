package rangeproductqueriesofpowersmid

import "testing"

const mod = 1e9 + 7

func productQueries(n int, queries [][]int) (ans []int) {
	powers := []int{}
	for n > 0 {
		lowbit := n & -n
		powers = append(powers, lowbit)
		n ^= lowbit
	}

	ans = make([]int, len(queries))
	for i, q := range queries {
		mul := 1
		for _, x := range powers[q[0] : q[1]+1] {
			mul = (mul * x) % mod
		}
		ans[i] = mul
	}

	return
}

func Test_range_product_queries_of_powers(t *testing.T) {
	tests := []struct {
		n       int
		queries [][]int
	}{
		{n: 15, queries: [][]int{{0, 1}, {2, 2}, {0, 3}}},
		{n: 2, queries: [][]int{{0, 0}}},
	}
	for _, tt := range tests {
		result := productQueries(tt.n, tt.queries)
		t.Logf("n: %d, queries: %v, result: %v", tt.n, tt.queries, result)
	}
}
