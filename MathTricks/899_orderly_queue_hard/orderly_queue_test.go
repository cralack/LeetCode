package orderlyqueuehard

import (
	"sort"
	"testing"
)

func orderlyQueue(s string, k int) string {
	if k == 1 { //字符串即为环,遍历求最小字典序str
		ans := s
		for i := 0; i < len(s); i++ {
			cur := s[i:] + s[:i]
			if cur < ans {
				ans = cur
			}
		}
		return ans
	} else { //总能有序,排序
		byt := []byte(s)
		sort.Slice(byt, func(i, j int) bool { return byt[i] < byt[j] })
		return string(byt)
	}
}
func Test_orderly_queue(t *testing.T) {
	s, k := "cba", 1
	t.Log(orderlyQueue(s, k))
	s, k = "baaca", 3
	t.Log(orderlyQueue(s, k))
	s, k = "v", 1
	t.Log(orderlyQueue(s, k))
	s, k = "xisxr", 1
	t.Log(orderlyQueue(s, k))
}
