package randompickwithweight

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"LeetCode/util/GetPrime"
)

type Solution struct {
	Max    int
	PreSum []int
}

func Constructor(w []int) Solution {
	// 构建前缀和数组，偏移一位留给 preSum[0]
	pSum := make([]int, len(w)+1)
	for i := range w {
		// preSum[i] = sum(w[0..i-1])
		pSum[i+1] = pSum[i] + w[i]
	}
	return Solution{PreSum: pSum, Max: pSum[len(w)]}
}

func (p *Solution) PickIndex() int {
	if len(p.PreSum) == 0 {
		return -1
	}
	// rand.Seed(int64(time.Now().Nanosecond()))
	// 在闭区间 [1, max] 中选择一个数字
	// 前缀和数组中 0 本质上是个占位符
	tar := rand.Intn(p.Max) + 1
	lo, hi := 0, len(p.PreSum)
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if p.PreSum[mi] >= tar {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
func Test_random_pick_with_weight(t *testing.T) {
	arr := GetPrime.PrimeArr(1e+5)
	sol := Constructor(arr)
	var waiter sync.WaitGroup
	n := 15 // thread num

	waiter = sync.WaitGroup{}
	waiter.Add(n)
	t.Logf("MAX nums:%d", len(sol.PreSum)-1)
	for i := 0; i < n; i++ {
		go func(idx int) {
			t.Log(sol.PickIndex())
			waiter.Done()
		}(i)
		time.Sleep(1 * time.Nanosecond)
	}
	waiter.Wait()
}
