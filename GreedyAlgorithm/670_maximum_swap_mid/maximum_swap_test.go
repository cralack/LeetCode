package maximumswapmid

import (
	"strconv"
	"testing"
)

func maximumSwap(num int) int {
	str := []byte(strconv.Itoa(num))
	n := len(str)
	maxIdx, idx1, idx2 := n-1, -1, -1
	for i := n - 1; i >= 0; i-- {
		if str[i] > str[maxIdx] {
			maxIdx = i
		} else if str[i] < str[maxIdx] {
			idx1, idx2 = i, maxIdx
		}
	}
	if idx1 < 0 {
		return num
	}
	str[idx1], str[idx2] = str[idx2], str[idx1]
	num, _ = strconv.Atoi(string(str))
	return num
}

func Test_maximum_swap(t *testing.T) {
	num := 2736
	t.Log(maximumSwap(num))
	num = 9973
	t.Log(maximumSwap(num))
	num = 98368
	t.Log(maximumSwap(num))
}
