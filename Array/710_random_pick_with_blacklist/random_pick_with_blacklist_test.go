package randompickwithblacklist

import (
	"math/rand"
	"testing"
)

type Solution struct {
	Size      int
	Black2idx map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	size := n - len(blacklist)
	black := make(map[int]int, 0)
	// 先将所有黑名单数字加入 map
	for _, b := range blacklist {
		// 这里赋值多少都可以
		// 目的仅仅是把键存进哈希表
		// 方便快速判断数字是否在黑名单内
		black[b] = -1
	}
	last := n - 1
	for _, b := range blacklist {
		// 如果 b 已经在区间 [sz, N)
		// 可以直接忽略
		if b >= size {
			continue
		}
		// 跳过所有黑名单中的数字
		for {
			if _, ok := black[last]; ok {
				last--
			} else {
				break
			}
		}
		// 将黑名单中的索引映射到合法数字
		black[b] = last
		last--
	}
	return Solution{
		Size:      size,
		Black2idx: black,
	}
}

func (p *Solution) Pick() int {
	// 随机选取一个索引
	idx := rand.Intn(p.Size)
	// 这个索引命中了黑名单，
	// 需要被映射到其他位置
	if val, ok := p.Black2idx[idx]; ok {
		return val
	}
	// 若没命中黑名单，则直接返回
	return idx
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
func Test_random_pick_with_blacklist(t *testing.T) {
	c1 := []string{"Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"}
	n, blacklist := 7, []int{2, 3, 5}
	var sol Solution
	for _, c := range c1 {
		switch c {
		case "Solution":
			sol = Constructor(n, blacklist)
		case "pick":
			t.Log(sol.Pick())
		}

	}
}
