package addingtwonegabinarynumbers

import "testing"

func addNegabinary(arr1 []int, arr2 []int) (ans []int) {
	i, j := len(arr1)-1, len(arr2)-1
	for carry := 0; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		tmp := carry
		if i >= 0 {
			tmp += arr1[i]
		}
		if j >= 0 {
			tmp += arr2[j]
		}
		carry = 0
		if tmp >= 2 {
			tmp -= 2
			carry -= 1
		} else if tmp == -1 {
			tmp = 1
			carry += 1
		}
		ans = append(ans, tmp)
	}
	for len(ans) > 1 && ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	for i, j = 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func Test_adding_two_negabinary_numbers(t *testing.T) {
	tests := []struct {
		arr1, arr2 []int
	}{
		{[]int{1, 1, 1, 1, 1}, []int{1, 0, 1}},
		{[]int{0}, []int{0}},
		{[]int{0}, []int{1}},
	}
	for _, tt := range tests {
		t.Log(addNegabinary(tt.arr1, tt.arr2))
	}
}
