package goSort

import "testing"

/****    冒泡排序    ****/

// basic algorithm
func buble_sort_basic(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		// 每轮循环，通过依次向右比较两个数，将本轮循环中最大的数放到最右
		for j := 1; j < len(arr)-i; j++ {
			// 若左大于右则交换，并将swapped置为true
			if arr[j-1] > arr[j] {
				swap(arr, j-1, i)
			}
		}
	}
	return arr
}

// 提前结束优化
func buble_sort_v1(arr []int) []int {
	// n-1轮次执行，当前n-1个元素排好后，最后一个元素无需执行
	for i := 0; i < len(arr)-1; i++ {
		// 本轮执行是否有交换
		swaped := false
		// 每轮循环，通过依次向右比较两个数，将本轮循环中最大的数放到最右
		for j := 1; j < len(arr)-i; j++ {
			// 若左大于右则交换，并将swapped置为true
			if arr[j-1] > arr[j] {
				swap(arr, j-1, i)
				swaped = true
			}
		}
		// 若无交换，表示当前数组已完全排序，退出大循环
		if !swaped {
			break
		}
	}
	return arr
}

// 冒泡界优化
func buble_sort_v2(arr []int) []int {
	// 上一轮最后交换的下标
	lastSwappedIdx := len(arr) - 1
	for lastSwappedIdx > 0 {
		curSwappedIdx := 0
		// 每一轮 i 只需要考察到 lastSwappedIdx 前一位
		for i := 0; i < lastSwappedIdx; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				curSwappedIdx = i
			}
		}
		// 若当前轮次未发生交换，则lastSwappedIdx=0，则在下一次while判断时退出
		lastSwappedIdx = curSwappedIdx
	}
	return arr
}

func Benchmark_buble_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	b.Run("basic", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buble_sort_basic(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	b.Run("提前结束优化", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buble_sort_v1(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	b.Run("冒泡界优化", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buble_sort_v2(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
}
