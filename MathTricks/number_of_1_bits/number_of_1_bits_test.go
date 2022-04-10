package numberof1bits

import (
	"strconv"
	"testing"
)

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= num - 1
		cnt++
	}
	return cnt
}
func Test_number_of_1_bits(t *testing.T) {
	num, _ := strconv.ParseUint("00000000000000000000000000001011", 2, 64)
	t.Log(hammingWeight(uint32(num)))
	num, _ = strconv.ParseUint("00000000000000000000000010000000", 2, 64)
	t.Log(hammingWeight(uint32(num)))
	num, _ = strconv.ParseUint("00000000000000000000000010000000", 2, 64)
	t.Log(hammingWeight(uint32(num)))
}
