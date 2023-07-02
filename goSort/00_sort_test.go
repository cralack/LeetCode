package goSort

import "testing"

// Benchmark中Knuth_shuffle生成的测试用例大小参数
const MaxN int = 1e4

var TestArr []int

func Benchmark_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	// init test arr
	TestArr = Knuth_shuffle(MaxN)

	b.Run("01.冒泡排序", func(b *testing.B) {
		Benchmark_buble_sort(b)
	})
	b.Run("02.选择排序", func(b *testing.B) {
		Benchmark_selection_sort(b)
	})
	b.Run("03.插入排序", func(b *testing.B) {
		Benchmark_insertion_sort(b)
	})
	b.Run("04.shell排序", func(b *testing.B) {
		Benchmark_shell_sort(b)
	})
	b.Run("05.归并排序", func(b *testing.B) {
		Benchmark_merge_sort(b)
	})
	b.Run("06.快速排序", func(b *testing.B) {
		Benchmark_quick_sort(b)
	})
	b.Run("07.堆排序", func(b *testing.B) {
		Benchmark_heap_sort(b)
	})
	b.Run("08.计数排序", func(b *testing.B) {
		Benchmark_count_sort(b)
	})
	b.Run("09.基数排序", func(b *testing.B) {
		Benchmark_radix_sort(b)
	})
	b.Run("10.桶排序", func(b *testing.B) {
		Benchmark_bucket_sort(b)
	})

}

// math func
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
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
