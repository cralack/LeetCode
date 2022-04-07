package partitiontokequalsumsubsets

import (
	"sort"
	"testing"
)

var (
	numberz []int
	tar     int
	uzed    int
	memo    map[int]bool
)

func canPartitionKSubsets_v1(nums []int, k int) bool {
	if k == 0 {
		return true
	}
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}
	sort.Ints(nums)
	// k 号桶初始什么都没装，从 nums[0] 开始做选择
	tar = sum / k
	uzed = 0
	memo = make(map[int]bool, 0)
	numberz = nums
	return backtrack(k, 0, len(numberz)-1)
}

func backtrack(k, bucket, start int) bool {
	//base case
	if k == 0 {
		// 所有桶都被装满了，而且 nums 一定全部用完了
		// 因为 target == sum / k
		return true
	}

	if bucket == tar {
		// 装满了当前桶，递归穷举下一个桶的选择
		// 让下一个桶从 nums[0] 开始选数字
		res := backtrack(k-1, 0, len(numberz)-1)
		// 缓存结果
		memo[uzed] = res
		return res
		// return backtrack(k-1, 0, len(numberz)-1)
	}
	// 避免冗余计算
	if ans, ok := memo[uzed]; ok {
		return ans
	}
	for i := start; i >= 0; i-- {
		if ((uzed >> i) & 1) == 1 { //剪枝
			continue // nums[i] 已经被装入别的桶中
		}
		if numberz[i]+bucket > tar {
			continue // 当前桶装不下 nums[i]
		}
		// 做选择，将 nums[i] 装入当前桶中
		uzed |= 1 << i // 将第 i 位置为 1
		bucket += numberz[i]
		// 递归穷举下一个数字是否装入当前桶
		if backtrack(k, bucket, i-1) {
			return true
		}
		// 撤销选择
		uzed ^= 1 << i // 使用异或运算将第 i 位恢复 0
		bucket -= numberz[i]
		for i > 0 && numberz[i-1] == numberz[i] {
			i--
		}
	} // 穷举了所有数字，都无法装满当前桶
	return false
}

var (
	used    []bool
	numbers []int
	target  int
)

func canPartitionKSubsets_v2(nums []int, k int) bool {
	var sum int = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target = sum / k

	sort.Ints(nums)
	if nums[len(nums)-1] > target {
		return false
	}

	numbers = nums
	used = make([]bool, len(nums))

	return traverse(k, len(nums)-1, 0)
}

func traverse(k int, start int, bucket int) bool {
	if k == 0 {
		return true
	}
	if bucket == target {
		return traverse(k-1, len(numbers)-1, 0)
	}
	for i := start; i >= 0; i-- {
		if used[i] {
			continue
		}
		if numbers[i]+bucket > target {
			continue
		}

		bucket, used[i] = bucket+numbers[i], true
		if traverse(k, i-1, bucket) {
			return true
		}
		bucket, used[i] = bucket-numbers[i], false
		for i > 0 && numbers[i-1] == numbers[i] {
			i--
		}
	}
	return false
}
func Test_partition_to_k_equal_sum_subsets(t *testing.T) {
	nums, k := []int{4, 3, 2, 3, 5, 2, 1}, 4
	t.Log(canPartitionKSubsets_v1(nums, k))
	nums, k = []int{1, 2, 3, 4}, 3
	t.Log(canPartitionKSubsets_v1(nums, k))
}

func Benchmark_partk(b *testing.B) {
	nums := []int{}
	for i := 1; i <= 20; i++ {
		nums = append(nums, i)
	}
	k := 5
	b.Run("v1", func(b *testing.B) { //best 32ms
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			canPartitionKSubsets_v1(nums, k)
		}
		b.StopTimer()
	})
	b.Run("v2", func(b *testing.B) { //best 0ms
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			canPartitionKSubsets_v2(nums, k)
		}
		b.StopTimer()
	})
}
