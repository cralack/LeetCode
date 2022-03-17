package alloonedatastructure

import (
	"fmt"
	"testing"
)

type Node struct {
	Cnt        int
	Keys       map[string]bool
	Prev, Next *Node
}
type AllOne struct {
	Cnt        map[string]int
	Cnt2Node   map[int]*Node
	Head, Tail *Node
}

func Constructor() AllOne {
	head := &Node{Cnt: -1}
	tail := &Node{Cnt: -2}
	head.Next = tail
	tail.Prev = head
	return AllOne{
		Cnt2Node: make(map[int]*Node),
		Cnt:      make(map[string]int),
		Head:     head,
		Tail:     tail,
	}
}
func (p *AllOne) AddNode(cur, tar *Node) {
	tar.Next = cur.Next
	tar.Prev = cur
	cur.Next.Prev = tar
	cur.Next = tar
	p.Cnt2Node[tar.Cnt] = tar
}
func (p *AllOne) DelNode(tar *Node) {
	tar.Next.Prev = tar.Prev
	tar.Prev.Next = tar.Next
	tar.Next = nil
	tar.Prev = nil
}

func (p *AllOne) Inc(key string) {
	cnt, ok := p.Cnt[key]
	//key exist
	if ok {
		oldNode := p.Cnt2Node[cnt]
		//need create new node
		if oldNode.Next.Cnt != cnt+1 {
			p.AddNode(oldNode, &Node{Cnt: cnt + 1, Keys: make(map[string]bool)})
		}
		newNode := oldNode.Next
		newNode.Keys[key] = true
		p.Cnt[key]++

		delete(oldNode.Keys, key)
		if len(oldNode.Keys) == 0 {
			delete(p.Cnt2Node, oldNode.Cnt)
			p.DelNode(oldNode)
		}
		return
	}
	//key doesnt exist
	//cnt1 node doesnt exist
	if p.Head.Next.Cnt != 1 {
		cur := p.Head
		p.AddNode(cur, &Node{
			Cnt:  1,
			Keys: make(map[string]bool),
		})
	}
	cur := p.Head.Next
	cur.Keys[key] = true
	p.Cnt[key]++
}

func (p *AllOne) Dec(key string) {
	cnt := p.Cnt[key]
	oldNode := p.Cnt2Node[cnt]
	//check prev node
	if cnt != 1 {
		//need create prev node
		if oldNode.Prev.Cnt != cnt-1 {
			p.AddNode(oldNode.Prev, &Node{Cnt: cnt - 1, Keys: make(map[string]bool)})
		}
		newNode := oldNode.Prev
		newNode.Keys[key] = true
	}
	delete(oldNode.Keys, key)
	if len(oldNode.Keys) == 0 {
		p.DelNode(oldNode)
		delete(p.Cnt2Node, cnt)
	}
	p.Cnt[key]--
	if p.Cnt[key] == 0 {
		delete(p.Cnt, key)
	}
}

func (p *AllOne) GetMaxKey() string {
	for str := range p.Tail.Prev.Keys {
		return str
	}
	return ""
}

func (p *AllOne) GetMinKey() string {
	for str := range p.Head.Next.Keys {
		return str
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
func Test_all_oone_data_structure(t *testing.T) {
	var alo AllOne
	c1 := []string{"AllOne", "inc", "inc", "inc", "inc",
		"inc", "inc", "getMaxKey", "inc", "dec",
		"getMaxKey", "dec", "inc", "getMaxKey", "inc",
		"inc", "dec", "dec", "dec", "dec",
		"getMaxKey", "inc", "inc", "inc", "inc",
		"inc", "inc", "getMaxKey", "getMinKey"}
	c2 := []string{"", "hello", "world", "leet", "code",
		"ds", "leet", "", "ds", "leet",
		"", "ds", "hello", "", "hello",
		"hello", "world", "leet", "code", "ds",
		"", "new", "new", "new", "new",
		"new", "new", "", ""}
	for i, c := range c1 {
		fmt.Printf("%02d:  c1:%s  c2:%s\r", i, c, c2[i])
		switch c {
		case "AllOne":
			alo = Constructor()
		case "inc":
			alo.Inc(c2[i])
		case "dec":
			alo.Dec(c2[i])
		case "getMaxKey":
			fmt.Printf("\rMax:%s", alo.GetMaxKey())
		case "getMinKey":
			fmt.Printf("\rMin:%s", alo.GetMinKey())
		}
		fmt.Println("\r")
	}
}
