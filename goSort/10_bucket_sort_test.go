package goSort

import "testing"

/****    用桶排序    ****/
func bucket_sort_v1(arr []int) []int {
	k := len(arr) / 3 // n/3 个桶
	mn, mx := arr[0], arr[0]
	// 确定 min 和 max
	for _, num := range arr[1:] {
		mn = min(mn, num)
		mx = max(mx, num)
	}
	buckets := make([][]int, k) // k 个桶
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	// 桶间隔
	var interval float64 = (float64(mx-mn) / float64(k-1))
	// 遍历arr，根据元素值将所有元素装入对应值区间的桶中
	for _, num := range arr {
		// arr[i] (num) 元素应该装入的桶的下标
		bucketIdx := int(float64(num-mn) / interval)
		// 装入对应桶中
		buckets[bucketIdx] = append(buckets[bucketIdx], num)
	}
	for i := range buckets {
		// 桶内排序(调用库函数，从小到大)
		quick_sort_v1(buckets[i])
	}
	index := 0
	for _, bucket := range buckets {
		for _, sortedNum := range bucket {
			arr[index] = sortedNum
			index++
		}
	}
	return arr
}

func Test_bucket_sort(t *testing.T) {
	//稳定性：取决于桶内排序方法的稳定性
	t.Log(bucket_sort_v1(Knuth_shuffle(15)))
}

func Benchmark_bucket_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	b.Run("以快排为基", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bucket_sort_v1(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
}
