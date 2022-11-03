package maximumfrequencystack

import (
	"LeetCode/WriteLog"
	"strconv"
	"strings"
	"testing"
)

type Node struct {
	//可封装任何结构
	Val int
}

type FreqStack struct {
	MaxFreq int
	Cnt     map[int]int
	Cache   map[int][]*Node
}

func Constructor() FreqStack {
	return FreqStack{
		MaxFreq: 0,
		Cnt:     make(map[int]int),
		Cache:   make(map[int][]*Node),
	}
}
func (p *FreqStack) Push(val int) {
	//val.freq
	p.Cnt[val]++
	freq := p.Cnt[val]
	//check is stack exist
	if _, ok := p.Cache[freq]; !ok {
		p.Cache[freq] = []*Node{}
	}
	//insert new node
	tmp := &Node{Val: val}
	p.Cache[freq] = append(p.Cache[freq], tmp)
	//check maxFreq
	if freq > p.MaxFreq {
		p.MaxFreq = freq
	}
}

func (p *FreqStack) Pop() int {
	//pop
	max_stack := p.Cache[p.MaxFreq]
	x := max_stack[len(max_stack)-1]
	p.Cache[p.MaxFreq] = max_stack[:len(max_stack)-1]
	//check maxfreq
	if len(p.Cache[p.MaxFreq]) == 0 {
		p.MaxFreq--
	}
	p.Cnt[x.Val]--
	return x.Val
}

// debug
func (p *FreqStack) log() string {
	sb := &strings.Builder{}
	for i := 1; i <= p.MaxFreq; i++ {
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString(":head->")
		// cur := p.Cache[i][0]
		// for cur != p.Cache[i].Tail {
		// 	cur = cur.Next
		// 	if cur != p.Cache[i].Tail {
		// 		sb.WriteString(strconv.Itoa(cur.Val))
		// 		sb.WriteString("->")
		// 	}
		// }
		for _, no := range p.Cache[i] {
			sb.WriteString(strconv.Itoa(no.Val))
			sb.WriteString("->")
		}
		sb.WriteString("tail\n")
	}
	return sb.String()
}

func Test_maximum_frequency_stack(t *testing.T) {
	// c1 := []string{"FreqStack", "push", "push",
	// 	"push", "push", "push", "push",
	// 	"pop", "pop", "pop", "pop"}
	// c2 := []int{-1, 5, 7,
	// 	5, 7, 4, 5,
	// 	-1, -1, -1, -1}
	c1 := []string{"FreqStack", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop", "pop"}
	c2 := []int{-1, 30, 40, 4, 25, 20, 40, 48, 21, 12, 44, 1, 16, 20, 9, 34, 26, 12, 21, 35, 16, 3, 23, 9, 31, 10, 6, -1, 45, -1, 16, -1, 14, -1, 27, -1, 35, -1, 34, -1, 40, -1, 13, -1, 21, -1, 18, -1, 26, -1, 29, -1, 32, -1, 3, -1, 18, -1, 36, -1, 1, -1, 38, -1, 34, -1, 20, -1, 22, -1, 13, -1, 37, -1, 24, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	var stack FreqStack
	for idx, str := range c1 {
		sb := &strings.Builder{}
		sb.WriteString(strconv.Itoa(idx))
		sb.WriteString(": c1:")
		sb.WriteString(str)
		sb.WriteString(" ")
		sb.WriteString(strconv.Itoa(c2[idx]))
		WriteLog.WriteFile("max_freq_stack", sb.String())
		switch str {
		case "FreqStack":
			stack = Constructor()
		case "push":
			stack.Push(c2[idx])
		case "pop":
			stack.Pop()
		}
		WriteLog.WriteFile("max_freq_stack", stack.log())
		// fmt.Println("\r")
	}
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
