package goSort

import "testing"

/****    堆排序    ****/
func heap_sort_v1(arr []int) []int {
	// 构建大顶堆
	heapify(arr, len(arr)-1)
	// i > 0 即可，无需写成 i >= 0，当 N-1 个元素排序时，最后一个元素也已排序
	for i := len(arr) - 1; i > 0; i-- {
		// 交换堆顶和当前未排序部分最后一个元素
		swap(arr, 0, i)
		// i - 1 是未排序部分最后一个元素下标，传入此参数确保下滤不会超过此范围
		percolateDown(arr, 0, i-1)
	}
	return arr
}

// 堆化方法
func heapify(arr []int, n int) {
	lastInternal := n>>1 - 1
	for i := lastInternal; i >= 0; i-- {
		percolateDown(arr, i, n)
	}
}

// 下滤方法
func percolateDown(arr []int, i, n int) {
	// target 是要下滤的结点的值
	tar := arr[i]
	// 赋i左孩子下标,rChild: (1+i)<<1
	child := 1 + i<<1
	for child <= n {
		// child<r表示i有右孩子
		if child < n && arr[child+1] > arr[child] {
			// 且右孩子更大,则令 child 为右孩子下标
			// 此时childe为i的孩子中值较大者的下标
			child++
		}
		// 若 child 大于 target
		if arr[child] > tar {
			// 则 arr[child] 上移到下标i处
			arr[i] = arr[child]
			// i 更新为 child (下滤)
			i = child
			// 更新 child
			child = i*2 + 1 // child = child * 2 + 1
		} else {
			// 若 arr[child] <= target ，说明下标 i 处已经满足堆序，退出循环
			break
		}
	}
	// 将 target 填入i中
	arr[i] = tar
}
func Test_heap_sort(t *testing.T) {
	arr := []int{4, 6, 2, 1, 7, 9, 5, 8, 3}
	t.Log(heap_sort_v1(arr))
}
func Benchmark_heap_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	b.Run("basic", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			heap_sort_v1(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
}
