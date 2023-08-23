package minimumnumberofoperations

import "testing"

func minOperations(boxes string) (ans []int) {
	n := len(boxes)
	ans = make([]int, n)
	// 预处理出每个位置 i 左边的小球移动到 i 的操作数leftPre
	// 每个位置 i 邮编的小球移动到 i 的操作数rightPre
	// leftPre, rightPre := make([]int, n), make([]int, n)

	for i, cnt := 1, 0; i < n; i++ {
		if boxes[i-1] == '1' {
			cnt++
		}
		// leftPre[i] = leftPre[i-1] + cnt
		// 进一步优化空间复杂度
		ans[i] = ans[i-1] + cnt
	}
	for i, cnt, sum := n-2, 0, 0; i >= 0; i-- {
		if boxes[i+1] == '1' {
			cnt++
		}
		// rightPre[i] = rightPre[i+1] + cnt
		// 进一步优化空间复杂度
		sum += cnt
		ans[i] += sum
	}
	// for i := range ans {
	// 	ans[i] = leftPre[i] + rightPre[i]
	// }
	return
}
func Test_minimum_number_of_operations(t *testing.T) {
	boxes := "110"
	t.Log(minOperations(boxes))
	boxes = "001011"
	t.Log(minOperations(boxes))
}
