package detonate_the_maximum_bombs_mid

import (
	"math/bits"
	"testing"
)

func maximumDetonation(bombs [][]int) (ans int) {
	n := len(bombs)
	graph := make([][]int, n)
	for from, cir1 := range bombs {
		x, y, r := cir1[0], cir1[1], cir1[2]
		for to, cir2 := range bombs {
			dx := x - cir2[0]
			dy := y - cir2[1]
			if to != from && dx*dx+dy*dy <= r*r {
				graph[from] = append(graph[from], to)
			}
		}
	}

	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(cur int) int {
		vis[cur] = true
		cnt := 1
		for _, next := range graph[cur] {
			if !vis[next] {
				cnt += dfs(next)
			}
		}
		return cnt
	}
	for i := range graph {
		clear(vis)
		ans = max(ans, dfs(i))
	}
	return
}

// bitset+floyd
func maximumDetonationBF(bombs [][]int) (ans int) {
	n := len(bombs)
	f := make([]bitset, n)
	for from, cir1 := range bombs {
		x, y, r := cir1[0], cir1[1], cir1[2]
		f[from] = newBitset(n)
		for to, cir2 := range bombs {
			dx := x - cir2[0]
			dy := y - cir2[1]
			if dx*dx+dy*dy <= r*r {
				f[from].set(to)
			}
		}
	}

	for k, fk := range f {
		for _, fi := range f {
			if fi.has(k) { // i 可以到达 k
				fi.or(fk) // i 也可以到 k 可以到达的点
			}
		}
	}

	for _, s := range f {
		ans = max(ans, s.onesCount()) // 集合大小的最大值
	}
	return
}

const w = bits.UintSize

type bitset []uint

func newBitset(n int) bitset {
	return make(bitset, (n+w-1)/w) // 需要 ceil(n/w) 个 w 位整数
}

func (b bitset) set(p int) {
	b[p/w] |= 1 << (p % w)
}

func (b bitset) has(p int) bool {
	return b[p/w]&(1<<(p%w)) != 0
}

func (b bitset) or(other bitset) {
	for i, x := range other {
		b[i] |= x
	}
}

func (b bitset) onesCount() (c int) {
	for _, x := range b {
		c += bits.OnesCount(x)
	}
	return
}

func Test_detonate_the_maximum_bombs(t *testing.T) {
	tests := []struct {
		bombs [][]int
		want  int
	}{
		{bombs: [][]int{{2, 1, 3}, {6, 1, 4}}, want: 2},
		{bombs: [][]int{{1, 1, 5}, {10, 10, 5}}, want: 1},
		{bombs: [][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}, want: 5},
	}
	for _, test := range tests {
		t.Log(maximumDetonation(test.bombs))
		t.Log(maximumDetonationBF(test.bombs) == test.want)
	}
	t.Log(w)
}
