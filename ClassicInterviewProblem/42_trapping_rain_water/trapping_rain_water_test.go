package trappingrainwater

import "testing"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	n, res := len(height), 0
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	for i := 1; i < n-1; i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}
func trap_double_pointer(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	res := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		// res += min(l_max, r_max) - height[i]
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}
func Test_trapping_rain_water(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	t.Log(trap_double_pointer(height))
	t.Log(trap(height))
	height = []int{4, 2, 0, 3, 2, 5}
	t.Log(trap_double_pointer(height))
	t.Log(trap(height))
}

func Benchmark_trap(b *testing.B) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	b.Run("memo", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			trap(height)
		}
		b.StopTimer()
	})
	b.Run("double_pointer", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			trap_double_pointer(height)
		}
		b.StopTimer()
	})
}
