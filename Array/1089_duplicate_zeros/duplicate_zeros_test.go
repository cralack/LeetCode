package duplicatezeros

import "testing"

func duplicateZeros(arr []int) {
	size := len(arr)
	i, j := 0, 0
	for j < size { // 确定有几个0
		if arr[i] == 0 {
			j++
		}
		i, j = i+1, j+1
	}
	i, j = i-1, j-1
	for i >= 0 { // 倒序重建数组
		if j < size {
			arr[j] = arr[i]
		}
		if arr[i] == 0 && j >= 1 {
			j--
			arr[j] = 0
		}
		i, j = i-1, j-1
	}
}
func Test_duplicate_zeros(t *testing.T) {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	t.Log(arr)

	arr = []int{1, 2, 3}
	duplicateZeros(arr)
	t.Log(arr)
}
