package BinHeap

import "fmt"

type Node struct {
	Val int
}
type BinHeap struct {
	// 大顶堆 Less(){return i<j}
	Size, Cap int
	Heap      []*Node // start from 1
}

func (b *BinHeap) parent(root int) int { return root / 2 }
func (b *BinHeap) left(root int) int   { return root * 2 }
func (b *BinHeap) right(root int) int  { return root*2 + 1 }
func (b *BinHeap) lastInternal() int   { return b.Size / 2 }
func (b *BinHeap) NewHeap(cap int)     { b.Heap = make([]*Node, cap+1); b.Cap = cap }
func (b *BinHeap) Peek() *Node         { return b.Heap[1] }
func (b *BinHeap) Less(i, j int) bool  { return b.Heap[i].Val < b.Heap[j].Val }
func (b *BinHeap) Swap(i, j int)       { b.Heap[i], b.Heap[j] = b.Heap[j], b.Heap[i] }
func (b *BinHeap) Swim(k int) {
	for k > 1 && b.Less(b.parent(k), k) {
		b.Swap(b.parent(k), k)
		k = b.parent(k)
	}
}
func (b *BinHeap) Sink(k int) {
	// 如果沉到堆底，就沉不下去了
	for b.left(k) <= b.Size {
		// 先假设左边节点较大
		old := b.left(k)
		// 如果右边节点存在，比一下大小
		if b.right(k) <= b.Size && b.Less(old, b.right(k)) {
			old = b.right(k)
		}
		// 结点 k 比俩孩子都大，就不必下沉了
		if b.Less(old, k) {
			break
		}
		// 否则，不符合最大堆的结构，下沉 k 结点
		b.Swap(k, old)
		k = old
	}
}
func (b *BinHeap) Insert(x *Node) {
	b.Size++
	// 先把新元素加到最后
	b.Heap[b.Size] = x
	// 然后让它上浮到正确的位置
	b.Swim(b.Size)
}
func (b *BinHeap) DelMax() *Node {
	max := b.Heap[1]
	b.Swap(1, b.Size)
	b.Heap[b.Size] = nil
	b.Size--
	// 让 pq[1] 下沉到正确位置
	b.Sink(1)
	return max
}

func (b *BinHeap) Init(arr []int) {
	b.NewHeap(len(arr))
	b.Size = b.Cap
	for i := 1; i < b.Size+1; i++ {
		b.Heap[i] = &Node{Val: arr[i-1]}
	}
	last := b.lastInternal()
	for i := last; i >= 1; i-- {
		b.Sink(i)
	}
}
func (b *BinHeap) Show() {
	i := 1
	h := 1
	for ; 1<<h <= b.Size; h++ {
		// fmt.Printf("h:%d 1<<h:%d", h, 1<<h)
		n := (b.Size - i) / 2
		for j := 0; j < n; j++ {
			fmt.Printf("\t")
		}
		for ; i <= 1<<h-1; i++ {
			fmt.Printf("val:%d\t", b.Heap[i].Val)
		}
		fmt.Printf("\r\n")
	}
	for ; i <= b.Size; i++ {
		fmt.Printf("val:%d\t", b.Heap[i].Val)
	}
	fmt.Printf("\r\n")
}
