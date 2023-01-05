package countpairswithxorinarange

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
)

// brute force
func countPairs(nums []int, low int, high int) (ans int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			cur := nums[i] ^ nums[j]
			if low <= cur && cur <= high {
				ans++
			}
		}
	}
	return
}

// trie
type Trie struct {
	children [2]*Trie
	cnt      int
}

func newTrie() *Trie {
	return &Trie{}
}

func (this *Trie) insert(x int) {
	node := this
	for i := 15; i >= 0; i-- {
		v := (x >> i) & 1
		if node.children[v] == nil {
			node.children[v] = newTrie()
		}
		node = node.children[v]
		node.cnt++
	}
}

func (this *Trie) search(x, limit int) (ans int) {
	node := this
	for i := 15; i >= 0 && node != nil; i-- {
		v := (x >> i) & 1
		if (limit >> i & 1) == 1 {
			if node.children[v] != nil {
				ans += node.children[v].cnt
			}
			node = node.children[v^1]
		} else {
			node = node.children[v]
		}
	}
	return
}

func countPairs_trie(nums []int, low int, high int) (ans int) {
	tree := newTrie()
	for _, x := range nums {
		ans += tree.search(x, high+1) - tree.search(x, low)
		tree.insert(x)
	}
	return
}

// hash
func countPairs_hash(nums []int, low, high int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for high++; high > 0; high >>= 1 {
		nxt := map[int]int{}
		for x, c := range cnt {
			ans += c * (high%2*cnt[(high-1)^x] - low%2*cnt[(low-1)^x])
			nxt[x>>1] += c
		}
		cnt = nxt
		low >>= 1
	}
	return ans / 2
}

func Test_count_pairs_with_xor_in_a_range(t *testing.T) {
	nums := []int{1, 4, 2, 7}
	low, high := 2, 6
	t.Log(countPairs(nums, low, high))
	nums = []int{9, 8, 4, 2, 1}
	low, high = 5, 14
	t.Log(countPairs(nums, low, high))
	nums = ReadTxt()
	low, high = 18678, 18727
	t.Log(countPairs(nums, low, high))
}

func ReadTxt() []int {
	bytes, err := ioutil.ReadFile("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(bytes)
	strs := strings.Split(str, ",")
	ans := make([]int, len(strs))
	for i := range strs {
		ans[i], _ = strconv.Atoi(strs[i])
	}
	return ans
}

// Benchmark
func Benchmark_count_pairs(b *testing.B) {
	nums := ReadTxt()
	low, high := 18678, 18727
	b.Run("brute force", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < 30*b.N; i++ {
			countPairs(nums, low, high)
		}
		b.StopTimer()
	})
	b.Run("trie", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < 30*b.N; i++ {
			countPairs_trie(nums, low, high)
		}
		b.StopTimer()
	})
	b.Run("hash", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < 30*b.N; i++ {
			countPairs_hash(nums, low, high)
		}
		b.StopTimer()
	})
}
