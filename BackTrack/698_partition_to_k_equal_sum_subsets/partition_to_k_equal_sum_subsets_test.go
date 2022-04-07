package partitiontokequalsumsubsets

import (
	"testing"
)

func canPartitionKSubsets(nums []int, k int) bool {
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
	// k 号桶初始什么都没装，从 nums[0] 开始做选择
	tar := sum / k
	used := 0
	memo := make(map[int]bool, 0)
	return backtrack(k, 0, 0, tar, used, nums, memo)
}

func backtrack(k, bucket, start, target, used int, nums []int, memo map[int]bool) bool {
	//base case
	if k == 0 {
		// 所有桶都被装满了，而且 nums 一定全部用完了
		// 因为 target == sum / k
		return true
	}
	// 避免冗余计算
	if ans, ok := memo[used]; ok {
		return ans
	}
	if bucket == target {
		// 装满了当前桶，递归穷举下一个桶的选择
		// 让下一个桶从 nums[0] 开始选数字
		res := backtrack(k-1, 0, 0, target, used, nums, memo)
		// 缓存结果
		memo[used] = res
		return res
	}

	for i := start; i < len(nums); i++ {
		if ((used >> i) & 1) == 1 { //剪枝
			continue // nums[i] 已经被装入别的桶中
		}
		if nums[i]+bucket > target {
			continue // 当前桶装不下 nums[i]
		}
		// 做选择，将 nums[i] 装入当前桶中
		used |= 1 << i // 将第 i 位置为 1
		bucket += nums[i]
		// 递归穷举下一个数字是否装入当前桶
		if backtrack(k, bucket, i+1, target, used, nums, memo) {
			return true
		}
		// 撤销选择
		used ^= 1 << i // 使用异或运算将第 i 位恢复 0
		bucket -= nums[i]
	} // 穷举了所有数字，都无法装满当前桶
	return false
}
func Test_partition_to_k_equal_sum_subsets(t *testing.T) {
	nums, k := []int{4, 3, 2, 3, 5, 2, 1}, 4
	t.Log(canPartitionKSubsets(nums, k))
	nums, k = []int{1, 2, 3, 4}, 3
	t.Log(canPartitionKSubsets(nums, k))
}
