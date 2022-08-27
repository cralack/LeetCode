package finalpricesez

import (
	"Learning/LeetCode/GetPrime"
	"math/rand"
	"testing"
)

func finalPrices(prices []int) []int {
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}

func finalPrices_v2(prices []int) []int {
	stack := make([]int, 0, len(prices)/2)
	for i := len(prices) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] > prices[i] {
			stack = stack[:len(stack)-1]
		}
		var res int
		if len(stack) == 0 {
			res = prices[i]
		} else {
			res = prices[i] - stack[len(stack)-1]
		}
		stack = append(stack, prices[i])
		prices[i] = res
	}
	return prices
}

func Test_final_prices(t *testing.T) {
	arr := []int{8, 4, 6, 2, 3}
	t.Log(finalPrices_v2(arr))
	arr = []int{1, 2, 3, 4, 5}
	t.Log(finalPrices_v2(arr))
}

func Benchmark_final_prices(b *testing.B) {
	n := 10000
	arr := GetPrime.PrimeArr(n)
	for i := 0; i < len(arr); i++ {
		idx := rand.Intn(len(arr))
		arr[i], arr[idx] = arr[idx], arr[i]
	}
	b.Run("brute force", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			finalPrices(arr)
		}
		b.StopTimer()
	})
	b.Run("monotonic stack", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			finalPrices_v2(arr)
		}
		b.StopTimer()
	})
}
