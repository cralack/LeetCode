package goSort

import "testing"

/****    基数排序    ****/

// 以计数排序为基础
func radix_sort_v1(arr []int) []int {
	n, mx := len(arr), abs(arr[0])
	// 找到 arr 中绝对值最大者
	for _, num := range arr[1:] {
		mx = max(mx, abs(num))
	}
	//// 求最大数的位数
	width, base := 0, 10
	for mx != 0 {
		width++
		mx /= base
	}
	sortedArr := make([]int, n)
	// 每一轮都对当前位(基数)执行一次计数排序
	for i := 0; i < width; i++ {
		// 为应对有负数的情况，countArr范围为[-9, 9]
		cntArr := make([]int, 19)
		// 根据每一个数字当前位的数字，累计相应位置的计数
		for _, num := range arr {
			// 求当前位上的 (+9处理负数)
			bucketIdx := (num%base)/(base/10) + 9
			cntArr[bucketIdx]++
		} // 变形
		for j := 1; j < len(cntArr); j++ {
			cntArr[j] += cntArr[j-1]
		} // 逆序输出保持稳定性
		for j := n - 1; j >= 0; j-- {
			// arr[j] 元素当前位对应 countArr 中的下标
			cntIdx := (arr[j]%base)/(base/10) + 9
			// 在排序后数组中的下标
			sortedIdx := cntArr[cntIdx] - 1
			// 在排序后数组中填入值
			sortedArr[sortedIdx] = arr[j]
			// countArr[countIdx] 已排序一位，下一个该位置的数的排位要靠前一位
			cntArr[cntIdx]--
		}
		copy(arr, sortedArr)
		base *= 10
	}
	return arr
}

// 不以计数排序为基础
func radix_sort_v2(arr []int) []int {
	n, mx := len(arr), abs(arr[0])
	// 找到 arr 中绝对值最大者
	for _, num := range arr[1:] {
		mx = max(mx, abs(num))
	}
	//// 求最大数的位数
	width, base := 0, 10
	for mx != 0 {
		width++
		mx /= base
	}
	// 19个桶处理负数，n+1 的用意是让每个桶的第一个位置保存该桶的元素个数
	buckets := make([][]int, 19)
	for i := range buckets {
		buckets[i] = make([]int, n+1)
	}
	//在每一位上将数组中所有具有该位的数字装入对应桶中
	for i := 0; i < width; i++ {
		for _, num := range arr {
			idx := (num%base)/(base/10) + 9
			cnt := buckets[idx][0]
			buckets[idx][cnt+1] = num
			buckets[idx][0]++
		}
		arrIdx := 0
		// 将当前所有桶的数按桶序，桶内从低到高输出为本轮排序结果
		for j := 0; j < len(buckets); j++ {
			for k := 1; k <= buckets[j][0]; k++ {
				arr[arrIdx] = buckets[j][k]
				arrIdx++
			}
		}
		// 每一轮过后将桶计数归零
		for _, bucket := range buckets {
			bucket[0] = 0
		}
		base *= 10
	}
	return arr
}
func Test_radix_sort(t *testing.T) {
	t.Log(radix_sort_v1(Knuth_shuffle(15)))
	t.Log(radix_sort_v2(Knuth_shuffle(15)))
}
func Benchmark_radix_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	b.Run("以计排为基", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			radix_sort_v1(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	b.Run("不以计排为基", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			radix_sort_v2(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
}
