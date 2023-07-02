package maxchunkstomakesortediihard

import (
	"testing"
)

func maxChunksToSorted(arr []int) int {
	stack := []int{}
	for _, val := range arr {
		if len(stack) == 0 || stack[len(stack)-1] <= val {
			stack = append(stack, val) // 单调递增入栈
		} else {
			max := stack[len(stack)-1] // pop
			stack = stack[:len(stack)-1]
			for len(stack) > 0 && stack[len(stack)-1] > val {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, max)
		}
	}
	return len(stack)
}

func maxChunksToSorted_v2(arr []int) int {
	stack := make([]int, 0, len(arr))
	for _, val := range arr {
		if len(stack) == 0 || stack[len(stack)-1] <= val {
			stack = append(stack, val) // 单调递增入栈
		} else {
			max := stack[len(stack)-1] // pop
			stack = stack[:len(stack)-1]
			for len(stack) > 0 && stack[len(stack)-1] > val {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, max)
		}
	}
	return len(stack)
}
func Test_max_chunks_to_make_sorted_ii(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	t.Log(maxChunksToSorted(arr))
	arr = []int{2, 1, 3, 4, 4}
	t.Log(maxChunksToSorted(arr))
	arr = []int{1, 2, 3, 4, 5}
	t.Log(maxChunksToSorted(arr))
	arr = []int{2, 1, 5, 2, 3, 4, 7, 8}
	t.Log(maxChunksToSorted_v2(arr))
}

func Benchmark_stack(b *testing.B) {
	arr := []int{2, 1, 5, 2, 3, 4, 7, 8}
	b.Run("stack_v1", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maxChunksToSorted(arr)
		}
		b.StopTimer()
	})

	b.Run("stack_v2", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maxChunksToSorted_v2(arr)
		}
		b.StopTimer()
	})
}
