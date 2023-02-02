package goSort

import (
	"math/rand"
	"testing"
	"time"
)

/****    快速排序    ****/

// core func
// partition方法执行后，主轴左边元素均小于主轴，
// 主轴右边元素均大等于主轴元素 ，返回主轴下标p
func partition(arr []int, left, right int) int {
	j := left + 1
	for i := j; i <= right; i++ {
		if arr[i] < arr[left] {
			// 交换后的 arr[j] 为当前最后一个小于主轴元素的元素
			swap(arr, i, j)
			j++
		}
	}
	// 主轴元素归位
	swap(arr, left, j-1)
	return j - 1
}

// 单轴快排递归实现
func quick_sort_v1(arr []int) []int {
	var quick_sort func(int, int)
	quick_sort = func(left, right int) {
		//若left==right，表示此时arr只有一个元素，即为基准情形，完成递归
		if left < right {
			p := partition(arr, left, right)
			quick_sort(left, p-1)
			quick_sort(p+1, right)
		}
	}
	quick_sort(0, len(arr)-1)
	return arr
}

// 将left, mid, right下标三个数中，大小居中者放到left下标处
func median3(arr []int, left, right int) {
	mid := left + (right-left)>>1
	// 左中，大者居中
	if arr[left] > arr[mid] {
		swap(arr, left, mid)
	} // 中右，大者居右，此时最大者居右
	if arr[mid] > arr[right] {
		swap(arr, mid, right)
	} // 左中，大者居左，此时中者居左
	if arr[mid] > arr[left] {
		swap(arr, left, mid)
	}
}

// 单轴快排递归实现,三数取中轴
func quick_sort_v2(arr []int) []int {
	var quick_sort_median3 func(int, int)
	quick_sort_median3 = func(left, right int) {
		// 执行median3将左，中，右三数中值放到left位置上
		if left < right {
			median3(arr, left, right)
			p := partition(arr, left, right)
			quick_sort_median3(left, p-1)
			quick_sort_median3(p+1, right)
		}
	}

	quick_sort_median3(0, len(arr)-1)
	return arr
}

// 单轴快排递归实现,随机轴
func quick_sort_v3(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	var quick_sort_rand func(left, right int)
	quick_sort_rand = func(left, right int) {
		if left < right {
			randIdx := rand.Intn(right-left) + left + 1
			swap(arr, left, randIdx)
			p := partition(arr, left, right)
			quick_sort_rand(left, p-1)
			quick_sort_rand(p+1, right)
		}
	}
	quick_sort_rand(0, len(arr)-1)
	return arr
}

// 单轴快排迭代实现(利用栈),首位为轴
func quick_sort_v4(arr []int) []int {
	// 保存区间左右边界的栈，按right到left的顺序将初始区间界入栈
	stack := &Stack{}
	stack.push(len(arr) - 1)
	stack.push(0)
	// 判断栈是否空，不空则弹出一对left，right界
	for !stack.isEmpty() {
		left, right := stack.pop(), stack.pop()
		// 执行partition的前提是 left 小于 right
		if left < right {
			p := partition(arr, left, right)
			stack.push(p - 1)
			stack.push(left)
			stack.push(right)
			stack.push(p + 1)
		}
	}
	return arr
}

// 单轴快排迭代实现(利用栈),三数取中
func quick_sort_v5(arr []int) []int {
	// 保存区间左右边界的栈，按right到left的顺序将初始区间界入栈
	stack := &Stack{}
	stack.push(len(arr) - 1)
	stack.push(0)
	// 判断栈是否空，不空则弹出一对left，right界
	for !stack.isEmpty() {
		left, right := stack.pop(), stack.pop()
		// 执行partition的前提是 left 小于 right
		if left < right {
			median3(arr, left, right)
			p := partition(arr, left, right)
			stack.push(p - 1)
			stack.push(left)
			stack.push(right)
			stack.push(p + 1)
		}
	}
	return arr
}

// 单轴快排迭代实现（利用栈）,随机轴
func quick_sort_v6(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	// 保存区间左右边界的栈，按right到left的顺序将初始区间界入栈
	stack := &Stack{}
	stack.push(len(arr) - 1)
	stack.push(0)
	// 判断栈是否空，不空则弹出一对left，right界
	for !stack.isEmpty() {
		left, right := stack.pop(), stack.pop()
		// 执行partition的前提是 left 小于 right
		if left < right {
			randIdx := rand.Intn(right-left) + left + 1
			swap(arr, left, randIdx)
			p := partition(arr, left, right)
			stack.push(p - 1)
			stack.push(left)
			stack.push(right)
			stack.push(p + 1)
		}
	}
	return arr
}

// 双轴快排
func quick_sort_v7(arr []int) []int {
	dual_pivot_quick_sort(arr, 0, len(arr)-1)
	return arr
}

/*
 *     区间1               区间2                         区间3
 * +------------------------------------------------------------+
 * |  < pivot1  | pivot1 <= && <= pivot2  |    ?    | > pivot2  |
 * +------------------------------------------------------------+
 *              ^                         ^         ^
 *              |                         |         |
 *            lower                     idx      upper
 */
func dual_pivot_quick_sort(arr []int, left, right int) {
	if left < right {
		// 排序对象是right大于left的区间（即大小大于等于2的区间）
		// 令左右两端元素中较小者居左，以left为初始pivot1，right为初始pivot2
		// 即arr[left]为选定的左轴值，arr[right]为选定的右轴值
		if arr[left] > arr[right] {
			swap(arr, left, right)
		}
		// 当前考察元素下标
		idx := left + 1
		// 用于推进到pivot1最终位置的动态下标,总有[left,lower)确定在区间1中
		lower := left + 1
		// 用于推进到pivot2最终位置的动态下标,总有(upper,right]确定在区间3中
		upper := right - 1
		// [lower, idx)确定在区间2中，[idx, upper]为待考察区间

		// upper以右（不含upper）的元素都是确定在区间3的元素，所以考察元素的右界是upper
		for idx <= upper {
			// 若arr[idx] < arr[left]，即arr[idx]小于左轴值，则arr[idx]位于区间1
			if arr[idx] < arr[left] {
				// 交换arr[idx]和arr[lower]，配合后一条lower++，保证arr[idx]位于区间1
				swap(arr, idx, lower)
				// lower++，扩展区间1，lower位置向右一位靠近pivot1的最终位置
				lower++

				// 若arr[idx] > arr[right]，即arr[idx]大于右轴值，则arr[idx]位于区间3
			} else if arr[idx] > arr[right] {
				// 先扩展区间3，使得如下while结束后upper以右（不含upper）的元素都位于区间3
				// 区间3左扩不可使idx == upper，否则之后的第二条upper--将导致upper为一个已经确定了区间归属的元素的位置（即idx - 1）
				for arr[upper] > arr[right] && idx < upper {
					upper--
				}
				// 交换arr[idx]和arr[upper]，配合后一条upper--，保证arr[idx]位于区间3
				swap(arr, idx, upper)
				upper--
				// 上述交换后，idx上的数字已经改变，只知道此时arr[idx] ≤ arr[right]，arr[idx]有可能在区间1或区间2，
				// 若arr[idx] < arr[left]，即arr[idx]小于左轴值，则arr[idx]位于区间1
				if arr[idx] < arr[left] {
					// 交换arr[idx]和arr[lower]，配合后一条lower++，保证arr[idx]位于区间1
					swap(arr, idx, lower)
					// lower++，扩展区间1，lower位置向右一位靠近pivot1的最终位置
					lower++
				}
			}
			// 考察下一个数字
			idx++
		}
		// while(idx <= upper)结束后最后一个确定在区间1的元素的下标是lower--，
		// 最后一个确定在区间3的元素下标是upper++。
		lower--
		upper++
		// 双轴归位。此时的lower，upper即分别为最终pivot1(初始时为left)，最终pivot2(初始时为right)。
		swap(arr, left, lower)
		swap(arr, upper, right)
		// 对三个子区间分别执行双轴快排
		dual_pivot_quick_sort(arr, left, lower-1)    // 区间1
		dual_pivot_quick_sort(arr, lower+1, upper-1) // 区间2
		dual_pivot_quick_sort(arr, upper+1, right)   // 区间3

	}
}

// implement stack
type Stack []int

func (s *Stack) isEmpty() bool { return len(*s) == 0 }
func (s *Stack) push(x int)    { *s = append(*s, x) }
func (s *Stack) pop() int {
	old := *s
	x := old[len(old)-1]
	old = old[:len(old)-1]
	*s = old
	return x
}
func Test_quick_sort(t *testing.T) {
	arr := Knuth_shuffle(30)
	t.Log(quick_sort_v1(append([]int{}, arr...)))
	t.Log(quick_sort_v2(append([]int{}, arr...)))
	t.Log(quick_sort_v3(append([]int{}, arr...)))
	t.Log(quick_sort_v4(append([]int{}, arr...)))
	t.Log(quick_sort_v5(append([]int{}, arr...)))
	t.Log(quick_sort_v6(append([]int{}, arr...)))
	t.Log(quick_sort_v7(append([]int{}, arr...)))
}

func Benchmark_quick_sort(b *testing.B) {
	if TestArr == nil {
		TestArr = Knuth_shuffle(MaxN)
	}
	b.Run("单轴快排递归实现", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			quick_sort_v1(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	// b.Run("单轴快排递归实现,三数取中轴", func(b *testing.B) {
	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		quick_sort_v2(append([]int{}, TestArr...))
	// 	}
	// 	b.StopTimer()
	// })
	// b.Run("单轴快排递归实现,随机轴", func(b *testing.B) {
	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		quick_sort_v3(append([]int{}, TestArr...))
	// 	}
	// 	b.StopTimer()
	// })
	b.Run("单轴快排迭代实现(利用栈),首位为轴", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			quick_sort_v4(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
	// b.Run("单轴快排迭代实现(利用栈),三数取中", func(b *testing.B) {
	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		quick_sort_v5(append([]int{}, TestArr...))
	// 	}
	// 	b.StopTimer()
	// })
	// b.Run("单轴快排迭代实现（利用栈）,随机轴", func(b *testing.B) {
	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		quick_sort_v6(append([]int{}, TestArr...))
	// 	}
	// 	b.StopTimer()
	// })
	b.Run("双轴快排", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			quick_sort_v7(append([]int{}, TestArr...))
		}
		b.StopTimer()
	})
}
