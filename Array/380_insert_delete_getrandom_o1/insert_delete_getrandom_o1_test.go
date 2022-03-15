package insertdeletegetrandomo1

import (
	"math/rand"
	"testing"
)

type RandomizedSet struct {
	Cache   []int
	Val2idx map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Cache:   make([]int, 0),
		Val2idx: make(map[int]int),
	}
}

func (p *RandomizedSet) Insert(val int) bool {
	if _, ok := p.Val2idx[val]; ok {
		return false
	}
	p.Val2idx[val] = len(p.Cache)
	p.Cache = append(p.Cache, val)
	return true
}

func (p *RandomizedSet) Remove(val int) bool {
	if _, ok := p.Val2idx[val]; !ok {
		return false
	}
	//val-idx
	idx := p.Val2idx[val]
	//end_idx
	end_idx := len(p.Cache) - 1
	//swap idx
	p.Val2idx[p.Cache[end_idx]] = idx
	//remove from Val2idx
	delete(p.Val2idx, val)
	//swap
	p.Cache[idx], p.Cache[end_idx] = p.Cache[end_idx], p.Cache[idx]
	//remove frome cache
	p.Cache = p.Cache[:len(p.Cache)-1]
	return true
}

func (p *RandomizedSet) GetRandom() int {
	return p.Cache[rand.Intn(len(p.Cache))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
func Test_insert_delete_getrandom_o1(t *testing.T) {
	c1 := []string{"RandomizedSet", "insert", "remove", "insert",
		"getRandom", "remove", "insert", "getRandom"}
	c2 := []int{-1, 1, 2, 2, -1, 1, 2, -1}
	var ran RandomizedSet
	for i, c := range c1 {
		switch c {
		case "RandomizedSet":
			ran = Constructor()
		case "insert":
			t.Log(ran.Insert(c2[i]))
		case "remove":
			t.Log(ran.Remove(c2[i]))
		case "getRandom":
			t.Log(ran.GetRandom())
		}
	}
}
