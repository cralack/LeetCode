package goSort

import "testing"

/****    选择排序    ****/

// 单元选择排序
func selection_sort_v1(arr []int) []int {
	//当前N-1个元素排好后，最后一个元素无需执行，故i<arr.length-1
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		//找到本轮执行中最小的元素，将最小值下标赋值给minIdx
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		swap(arr, i, minIdx)
	}
	return arr
}

// 双元选择排序(同时寻找min、max)
func selection_sort_v2(arr []int) []int {
	n := len(arr)
	//每轮确定两个数字，因此上界也会动态变化
	for i := 0; i < n-1-i; i++ {
		minIdx, maxIdx := i, i
		for j := i + 1; j < n-i; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		//若本轮max==min，说明未排序部分所有元素相等，无需再排序
		if minIdx == maxIdx {
			break
		}
		//<--min
		swap(arr, i, minIdx)
		//在交换i和minIdx时，有可能出现i即maxIdx的情况(特例1)
		//由于先将minIdx与i(maxIdx)进行了交换,此时需要修改maxIdx为minIdx
		if maxIdx == i {
			maxIdx = minIdx
		}
		//max-->
		swap(arr, n-1-i, maxIdx)
	}
	return arr
}
func Test_selcetion_v2(t *testing.T) {
	arr := []int{1, 2, 3, 3, 7, 4, 4, 5, 9}
	t.Log(selection_sort_v2(arr))
	//(特例1),i==3
	arr = []int{4, 3, 9, 6, 2, 5, 8, 7, 1}
	t.Log(selection_sort_v2(arr))
}
func Benchmark_selection_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	b.Run("单元选排", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			selection_sort_v1(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	b.Run("双元选排", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			selection_sort_v2(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
}
