package randompickindex

import (
	"math/rand"
	"testing"
)

type Solution struct {
	Nums []int
}

func Constructor(nums []int) Solution {
	return Solution{Nums: nums}
}

func (p *Solution) Pick1(target int) int {
	pool := []int{}
	for i, n := range p.Nums {
		if target == n {
			pool = append(pool, i)
		}
	}
	if len(pool) == 1 {
		return pool[0]
	}
	return pool[rand.Intn(len(pool))]
}
func (p *Solution) Pick(target int) int {
	res, index := -1, 1
	for i := 0; i < len(p.Nums); i++ {
		if p.Nums[i] == target {
			if rand.Intn(index) == 0 {
				res = i
			}
			index++
		}
	}
	return res
}
func Test_random_pick_index(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3}
	sol := Constructor(arr)
	for i := 0; i < 20; i++ {
		t.Log(sol.Pick(1))
	}
}

func Benchmark_rand_pick(b *testing.B) {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	sol := Constructor(arr)
	b.Run("pick1", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sol.Pick1(4)
		}
		b.StopTimer()
	})
	b.Run("pick1", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sol.Pick(4)
		}
		b.StopTimer()
	})
}
