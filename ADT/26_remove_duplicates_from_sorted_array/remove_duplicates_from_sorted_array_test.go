package removeduplicatesfromsortedarray

import "testing"

func removeDuplicates_v1(nums []int) int {
	slow, fast := 0, 0
	for fast = range nums {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]

		}
	}
	return slow + 1
}
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]

		}
		fast++
	}
	return slow + 1
}
func Test_remove_duplicates_from_sorted_array(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(removeDuplicates(arr))
}

func Benchmark_remove(b *testing.B) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	b.Run("V1", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			removeDuplicates_v1(arr)
		}
		b.StopTimer()
	})
	b.Run("V2", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			removeDuplicates(arr)
		}
		b.StopTimer()
	})
	arr = []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2,
		3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5,
		6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10}
	b.Run("V1_long", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			removeDuplicates_v1(arr)
		}
		b.StopTimer()
	})
	b.Run("V2_long", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			removeDuplicates(arr)
		}
		b.StopTimer()
	})
	// basically no difference
}
