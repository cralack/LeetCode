package design_hashmap

import (
	"testing"
)

type MyHashMap struct {
	List []int
}

func Constructor() MyHashMap {
	list := make([]int, 1e6+10)
	for i := range list {
		list[i] = -1
	}
	return MyHashMap{
		list,
	}
}

func (p *MyHashMap) Put(key int, value int) {
	p.List[key] = value
}

func (p *MyHashMap) Get(key int) int {
	return p.List[key]
}

func (p *MyHashMap) Remove(key int) {
	p.List[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func Test_design_hashmap(t *testing.T) {
	tests := []struct {
		cmd  []string
		imcd [][]int
	}{
		{[]string{"MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"},
			[][]int{{}, {1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}, {2}}},
	}
	for _, tt := range tests {
		var sol MyHashMap
		for i, c := range tt.cmd {
			switch c {
			case "MyHashMap":
				sol = Constructor()
			case "put":
				sol.Put(tt.imcd[i][0], tt.imcd[i][1])
			case "get":
				t.Log(sol.Get(tt.imcd[i][0]))
			case "remove":
				sol.Remove(tt.imcd[i][0])
			}
		}
	}
}
