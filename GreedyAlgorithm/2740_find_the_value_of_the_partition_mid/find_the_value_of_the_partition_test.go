package find_the_value_of_the_partition_mid

import (
	"slices"
	"sort"
	"testing"
)

func findValueOfPartition(nums []int) (ans int) {
	slices.Sort(nums)
	ans = 1<<31 - 1
	for i := 0; i < len(nums)-1; i++ {
		ans = min(ans, nums[i+1]-nums[i])
	}
	return
}

func Test_find_the_value_of_the_partition(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{nums: []int{1, 3, 2, 4}},
		{nums: []int{100, 1, 10}},
	}
	for _, test := range tests {
		t.Log(findValueOfPartition(test.nums))
	}
}

type Person struct {
	Name string
	Age  int
}

func Test_slice_sort(t *testing.T) {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}
