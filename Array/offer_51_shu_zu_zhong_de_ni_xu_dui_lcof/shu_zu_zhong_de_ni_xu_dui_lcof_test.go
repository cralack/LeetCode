package offer51shuzuzhongdenixuduilcof

import "testing"

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)>>1
	cnt := mergeSort(nums, start, mid) +
		mergeSort(nums, mid+1, end) //cnt包括整体+左半+右半
	tmp := []int{}             //辅助数组,用于存储排好序的数
	i, j := start, mid+1       //左指针从左半部分的左端开始遍历 右指针从右半部分的左端开始遍历
	for i <= mid && j <= end { //指针超出自己的那部分则退出循环
		if nums[i] <= nums[j] {
			//当左元素小于等于右元素 说明在右半部分的右元素的左边的元素均小于左元素
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1) //因此cnt需要加上右元素左边部分的元素数量
			i++
		} else {
			tmp = append(tmp, nums[j]) //当左元素大于右元素 不能直接开始计数
			j++                        //因为不确定右元素右边还有没有比左元素更小的元素
		}
	}
	for ; i <= mid; i++ { //处理另一半未遍历完的数据
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}
	for ; j <= end; j++ { //同上
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ { //将排序好的数组元素依次赋值给nums数组
		nums[i] = tmp[i-start]
	}
	return cnt
}

func Test_shu_zu_zhong_de_ni_xu_dui_lcof(t *testing.T) {
	arr := []int{7, 5, 6, 4}
	t.Log(reversePairs(arr))
}
