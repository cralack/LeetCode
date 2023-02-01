package Mysort

import (
	"math"
	"testing"
)

/****    希尔排序    ****/

// Shell增量 2^k
func shell_sort_v1(arr []int) []int {
	n := len(arr)
	// gap 初始为 n/2，缩小gap直到1
	for gap := n / 2; gap > 0; gap /= 2 {
		// 步长增量是gap，当前增量下需要对gap组序列进行简单插入排序
		for start := 0; start < gap; start++ {
			// 此for及下一个for对当前增量序列执行简单插入排序
			for i := start + gap; i < n; i += gap {
				tar, j := arr[i], i-gap
				for j >= 0 && arr[j] > tar {
					arr[j+gap] = arr[j]
					j -= gap
				}
				arr[j+gap] = tar
			}
		}
	}
	return arr
}

// Hibbard增量 2^k-1 {1, 3, 7, 15,...}
func shell_sort_v2(arr []int) []int {
	n, gap := len(arr), 1
	// 初始化gap (Hibbard增量序列)
	for gap < n/2 {
		gap = gap*2 + 1
	}
	// 缩小gap直到1
	for ; gap > 0; gap /= 2 {
		// 步长增量是gap，当前增量下需要对gap组序列进行简单插入排序
		for start := 0; start < gap; start++ {
			// 此for及下一个for对当前增量序列执行简单插入排序
			for i := start + gap; i < n; i += gap {
				tar, j := arr[i], i-gap
				for j >= 0 && arr[j] > tar {
					arr[j+gap] = arr[j]
					j -= gap
				}
				arr[j+gap] = tar
			}
		}
	}
	return arr
}

// Knuth增量 3^k-1 {1, 4, 13, 40,...}
func shell_sort_v3(arr []int) []int {
	n, gap := len(arr), 1
	// 初始化gap (Knuth增量序列)
	for gap < n/3 {
		gap = gap*3 + 1
	}

	for ; gap > 0; gap /= 3 {
		// 步长增量是gap，当前增量下需要对gap组序列进行简单插入排序
		for start := 0; start < gap; start++ {
			// 此for及下一个for对当前增量序列执行简单插入排序
			for i := start + gap; i < n; i += gap {
				tar, j := arr[i], i-gap
				for j >= 0 && arr[j] > tar {
					arr[j+gap] = arr[j]
					j -= gap
				}
				arr[j+gap] = tar
			}
		}
	}
	return arr
}

// Sedgewick增量 4^k + 3* 2^(k-1) {1, 4, 13, 40,...}
func shell_sort_v4(arr []int) []int {
	n, gap := len(arr), 1
	gaps := make([]int, 0)
	gaps = append(gaps, gap)
	// 初始化gap (Sedgewick增量序列)
	for k := 1; gap < n; k++ {
		gap = int(math.Pow(4, float64(k)) + 3*math.Pow(2, float64(k-1)) + 1)
		gaps = append(gaps, gap)
	}
	// 缩小gap直到1
	for idx := len(gaps) - 1; idx >= 0; idx-- {
		gap = gaps[idx]
		for start := 0; start < gap; start++ {
			// 此for及下一个for对当前增量序列执行简单插入排序
			for i := start + gap; i < n; i += gap {
				tar, j := arr[i], i-gap
				for j >= 0 && arr[j] > tar {
					arr[j+gap] = arr[j]
					j -= gap
				}
				arr[j+gap] = tar
			}
		}
	}
	return arr
}
func Benchmark_shell_sort(b *testing.B) {
	arr := Knuth_shuffle(MaxN)
	b.Run("Shell增量 2^k", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			shell_sort_v1(append([]int{}, arr...))
		}
		b.StopTimer()
	})
	b.Run("Hibbard增量 2^k-1", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			shell_sort_v2(append([]int{}, arr...))
		}
		b.StopTimer()
	})
	b.Run("Knuth增量 3^k-1", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			shell_sort_v3(append([]int{}, arr...))
		}
		b.StopTimer()
	})
	b.Run("Sedgewick增量 4^k + 3* 2^(k-1)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			shell_sort_v4(append([]int{}, arr...))
		}
		b.StopTimer()
	})
}
