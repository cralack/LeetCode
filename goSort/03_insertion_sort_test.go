package goSort

import (
	"sort"
	"testing"
)

/****    插入排序    ****/

// 简单插入排序
func insertion_sort_v1(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		tar := arr[i]
		for j >= 0 && arr[j] > tar {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tar
	}
	return arr
}

// 折半插入排序
func insertion_sort_v2(arr []int) []int {
	// binSearch := func(left, right, tar int) int {
	// 	for left <= right {
	// 		mid := left + (right-left)>>1
	// 		if arr[mid] <= tar {
	// 			left = mid + 1
	// 		} else {
	// 			right = mid - 1
	// 		}
	// 	}
	// 	return left
	// }
	// n-1 轮次执行
	for i := 1; i < len(arr); i++ {
		// 通过二分查找得到插入位置
		tar := arr[i]
		// arr[0,i)是排序状态
		// pos := binSearch(0, i-1, tar)
		pos := sort.Search(i, func(x int) bool {
			return arr[x] > tar
		})
		for j := i; j > pos; j-- {
			arr[j] = arr[j-1]
		}
		arr[pos] = tar
	}
	return arr
}
func Benchmark_insertion_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	b.Run("简单插排", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			insertion_sort_v1(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	b.Run("折半插排", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			insertion_sort_v2(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
}
