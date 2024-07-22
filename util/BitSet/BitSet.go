package BitSet

import (
	"math/bits"
)

const w = bits.UintSize

type Bitset []uint

func NewBitset(n int) Bitset {
	return make(Bitset, (n+w-1)/w) // 需要 ceil(n/w) 个 w 位整数
}

func (b Bitset) Set(pos int) {
	b[pos/w] |= 1 << (pos % w)
}

func (b Bitset) Has(pos int) bool {
	return b[pos/w]&(1<<(pos%w)) != 0
}

func (b Bitset) Or(other Bitset) {
	for i, x := range other {
		b[i] |= x
	}
}

func (b Bitset) OnesCount() (cnt int) {
	for _, x := range b {
		cnt += bits.OnesCount(x)
	}
	return
}
