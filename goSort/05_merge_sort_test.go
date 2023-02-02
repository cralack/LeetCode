package goSort

import "testing"

/****    归并排序    ****/

// 自顶向下非原地归并
func merge_sort_v1(nums []int) []int {
	//mergeSort 递归方法
	var mergeSort func([]int, []int, int, int)
	mergeSort = func(arr, tmpArr []int, left, right int) {
		if left < right {
			mid := left + (right-left)>>1
			mergeSort(arr, tmpArr, left, mid)
			mergeSort(arr, tmpArr, mid+1, right)
			merge_Extra_Space(arr, tmpArr, left, right, mid)
		}
	}

	tmpNums := make([]int, len(nums))
	mergeSort(nums, tmpNums, 0, len(nums)-1)
	return nums
}

// 非原地合并方法
func merge_Extra_Space(arr, tmpArr []int, left, right, mid int) {
	leftP, rightP, i := left, mid+1, left
	for ; leftP <= mid && rightP <= right; i++ {
		if arr[leftP] <= arr[rightP] {
			tmpArr[i] = arr[leftP]
			leftP++
		} else {
			tmpArr[i] = arr[rightP]
			rightP++
		}
	}
	// 左半边还有剩余，加入 tmpArr 末尾
	for ; leftP <= mid; i++ {
		tmpArr[i] = arr[leftP]
		leftP++
	}
	// 右半边还有剩余，加入 tmpArr 末尾
	for ; rightP <= right; i++ {
		tmpArr[i] = arr[rightP]
		rightP++
	}
	// 将 tmpArr 拷回 arr 中
	for ; left <= right; left++ {
		arr[left] = tmpArr[left]
	}
}

// 自顶向下原地归并
func merge_sort_v2(nums []int) []int {
	//mergeSort 递归方法
	var mergeSort func([]int, int, int)
	mergeSort = func(arr []int, left, right int) {
		if left < right {
			mid := left + (right-left)>>1
			mergeSort(arr, left, mid)
			mergeSort(arr, mid+1, right)
			merge_no_Extra_Space(arr, left, right, mid)
		}
	}
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

// 原地翻转
func reverse(arr []int, left, right int) {
	for ; left < right; left, right = left+1, right-1 {
		swap(arr, left, right)
	}
}

// 三次翻转实现交换
func exchange(arr []int, left, mid, right int) {
	reverse(arr, left, mid)
	reverse(arr, mid+1, right)
	reverse(arr, left, right)
}

// 原地归并(手摇算法)
func merge_no_Extra_Space(arr []int, left, right, mid int) {
	i, j := left, mid+1 //#1
	for i <= mid && j <= right {
		for ; i < j && arr[i] <= arr[j]; i++ { //#2
		}
		index := j //#3
		//#4注意是arr[j]<arr[i]，即找到j使得arr[j]为第一个大于等于arr[i]值
		for ; j <= right && arr[j] < arr[i]; j++ {
		}
		exchange(arr, i, index-1, j-1) //#5
	}
}

// 自底向上非原地
func merge_sort_v3(arr []int) []int {
	n := len(arr)
	tmpArr := make([]int, n)
	// 基本分区合并(随着间隔的成倍增长，一一合并，二二合并，四四合并...)
	for gap := 1; gap < n; gap *= 2 {
		// leftEnd = left+gap-1; rightEnd = left+2*gap-1;
		for left := 0; left < n-gap; left += 2 * gap {
			// 防止最后一次合并越界
			right := min(left+2*gap-1, n-1)
			merge_Extra_Space(arr, tmpArr, left, right, left+gap-1)
		}
	}
	return arr
}

// 自底向上原地
func merge_sort_v4(arr []int) []int {
	n := len(arr)
	// 基本分区合并(随着间隔的成倍增长，一一合并，二二合并，四四合并...)
	for gap := 1; gap < n; gap *= 2 {
		// leftEnd = left+gap-1; rightEnd = left+2*gap-1;
		for left := 0; left < n-gap; left += 2 * gap {
			// 防止最后一次合并越界
			right := min(left+2*gap-1, n-1)
			merge_no_Extra_Space(arr, left, right, left+gap-1)
		}
	}
	return arr
}

func Benchmark_merge_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	b.Run("T2B非原地", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			merge_sort_v1(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	// b.Run("T2b原地", func(b *testing.B) {
	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		merge_sort_v2(append([]int{}, TestArr...))
	// 	}
	// 	b.StopTimer()
	// })
	b.Run("B2T非原地", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			merge_sort_v3(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	// b.Run("B2T原地", func(b *testing.B) {
	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		merge_sort_v4(append([]int{}, TestArr...))
	// 	}
	// 	b.StopTimer()
	// })
}
