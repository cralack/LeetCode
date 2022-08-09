package designanorderedstreamez

import "testing"

type OrderedStream struct {
	Ptr   int
	Cache []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		Ptr:   1,
		Cache: make([]string, n+1),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) (ans []string) {
	this.Cache[idKey] = value
	if idKey != this.Ptr {
		return []string{}
	}
	for ; this.Ptr != len(this.Cache) &&
		this.Cache[this.Ptr] != ""; this.Ptr++ {
		ans = append(ans, this.Cache[this.Ptr])
	}
	return
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
func Test_design_an_ordered_stream(t *testing.T) {
	keys := []int{3, 1, 2, 5, 4}
	vals := []string{"ccccc", "aaaaa", "bbbbb", "eeeee", "ddddd"}
	sol := Constructor(5)
	for i, key := range keys {
		t.Log(sol.Insert(key, vals[i]))
	}
}
