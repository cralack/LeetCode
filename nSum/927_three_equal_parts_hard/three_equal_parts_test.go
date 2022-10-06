package threeequalparts

import (
	"testing"
)

func threeEqualParts(arr []int) []int {
	cnt := 0
	for _, n := range arr {
		if n == 1 {
			cnt++
		}
	}
	if cnt == 0 {
		return []int{0, 2}
	}
	if cnt%3 != 0 {
		return []int{-1, -1}
	}
	tar := cnt / 3
	first, second, third, cur := 0, 0, 0, 0
	//三指针对齐至对应1倍数位置
	for i, x := range arr {
		if x == 1 {
			if cur == 0 {
				first = i
			} else if cur == tar {
				second = i
			} else if cur == 2*tar {
				third = i
			}
			cur++
		}
	}
	n := len(arr)
	//1后置0长度
	length := n - third
	if first+length <= second && second+length <= third {
		i := 0
		for third+i < n {
			if arr[first+i] != arr[second+i] || arr[first+i] != arr[third+i] {
				return []int{-1, -1}
			}
			i++
		}
		return []int{first + length - 1, second + length}
	}
	return []int{-1, -1}
}

func Test_three_equal_parts(t *testing.T) {
	arr := []int{1, 0, 1, 0, 1}
	t.Log(threeEqualParts(arr))
	arr = []int{1, 1, 0, 1, 1}
	t.Log(threeEqualParts(arr))
	arr = []int{1, 1, 0, 0, 1}
	t.Log(threeEqualParts(arr))
}
