package platesbetweencandles

import (
	"runtime/debug"
	"testing"
)

func platesBetweenCandles1(s string, queries [][]int) []int {
	//*=42,|=124
	ans := []int{}
	for _, val := range queries {
		cnt := 0
		start := val[0]
		end := val[1]
		desStr := s[start : end+1]

		for _, v := range desStr {
			if v == 42 {
				start++
			} else {
				break
			}

		}
		for i := len(desStr) - 1; i > 0; i-- {
			if desStr[i] == 42 {
				end--
			} else {
				break
			}
		}

		if start >= end {
			ans = append(ans, 0)
			continue
		} else {
			desStr = s[start:end]
			for _, v := range desStr {
				if v == 42 {
					cnt++
				}
			}
			ans = append(ans, cnt)
		}
	}
	return ans
}

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	sum := make([]int, n+1) // sum[i] 表示 s[:i] 中盘子的个数
	left := make([]int, n)  // left[i] 表示 i 左侧最近蜡烛位置
	p := -1
	for i, b := range s {
		sum[i+1] = sum[i]
		if b == '|' {
			p = i
		} else {
			sum[i+1]++
		}
		left[i] = p
	}

	right := make([]int, n) // right[i] 表示 i 右侧最近蜡烛位置
	for i, p := n-1, n; i >= 0; i-- {
		if s[i] == '|' {
			p = i
		}
		right[i] = p
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := right[q[0]], left[q[1]] // 用最近蜡烛位置来代替查询的范围，从而去掉不符合要求的盘子
		if l < r {
			ans[i] = sum[r] - sum[l]
		}
	}
	return ans
}

func init() {
	debug.SetGCPercent(-1)
}

func CandleBinarySearch(slice []int, tar int) int {
	lo, hi := 0, len(slice)-1
	res := -1
	for lo <= hi {
		mi := (lo + hi) >> 1
		if slice[mi] >= tar {
			hi = mi - 1
			res = mi
		} else if slice[mi] < tar {
			lo = mi + 1
		} else {
			return mi
		}
	}
	return res
}

func TestPlatesBetweenCandles(t *testing.T) {
	s := "***|**|*****|**||**|*"
	var queries = [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}
	ans := platesBetweenCandles(s, queries)
	t.Log(ans)
}

func TestBinsearch(t *testing.T) {
	arr := []int{3, 6, 12, 15, 16, 19}
	tar := []int{1, 4, 14, 5, 15}
	for _, i := range tar {
		ret := CandleBinarySearch(arr, i)
		t.Log(arr[ret])
	}
}

func BenchmarkPlatesBetweenCandlesMine(b *testing.B) {
	s := "***|**|*****|**||**|*"
	var queries = [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		platesBetweenCandles1(s, queries)
	}
	b.StopTimer()
}
func BenchmarkPlatesBetweenCandles(b *testing.B) {
	s := "***|**|*****|**||**|*"
	var queries = [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		platesBetweenCandles(s, queries)
	}
	b.StopTimer()
}
