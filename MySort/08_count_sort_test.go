package goSort

import "testing"

/****    计数排序    ****/

// 不稳定计数排序
func count_sort_v1(arr []int) []int {
	n := len(arr)
	// 确定 min 和 max
	mn, mx := arr[0], arr[1]
	for _, num := range arr[1:] {
		mn = min(mn, num)
		mx = max(mx, num)
	}
	// arr 最多有 mx-mn+1 种数字
	cntArr := make([]int, mx-mn+1)
	// 计数
	for i := 0; i < n; i++ {
		cntArr[arr[i]-mn]++
	}
	index := 0
	// 依次取出
	for i := 0; i < len(cntArr); i++ {
		for j := 0; j < cntArr[i]; j++ {
			// 排序
			arr[index] = i + mn
			index++
		}
	}
	return arr
}

// 稳定计数排序
func count_sort_v2(arr []int) []int {
	n := len(arr)
	// 确定 min 和 max
	mn, mx := arr[0], arr[1]
	for i := 1; i < n; i++ {
		mn = min(mn, arr[i])
		mx = max(mx, arr[i])
	}
	// arr 最多有 mx-mn+1 种数字
	cntArr := make([]int, mx-mn+1)
	// 计数
	for i := 0; i < n; i++ {
		cntArr[arr[i]-mn]++
	}

	// 变形
	for i := 1; i < len(cntArr); i++ {
		cntArr[i] += cntArr[i-1]
	}
	// 根据 sortedArr, arr, countArr 三者关系完成 sortedArr 的输出
	sortedArr := make([]int, n)
	// 逆序输出保持稳定性
	for i := n - 1; i >= 0; i-- {
		// arr[i] 元素对应 countArr 中的下标
		cntIdx := arr[i] - mn
		// 在排序后数组中的下标
		sortedIdx := cntArr[cntIdx] - 1
		// 在排序后数组中填入值
		sortedArr[sortedIdx] = arr[i]
		// countArr[countIdx] 已排序一位，下一个该位置的数的排位要靠前一位
		cntArr[cntIdx]--
	}
	return sortedArr
}
func Test_count_sort(t *testing.T) {
	arr := []int{4, 6, 2, 1, 7, 9, 5, 8, 3, 1, 1}
	t.Log(count_sort_v2(arr))
}
func Benchmark_count_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	b.Run("不稳定计排", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			count_sort_v1(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	b.Run("稳定计排", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			count_sort_v2(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
}
