package goSort

import (
	"math/rand"
	"time"
)

func Knuth_shuffle(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	for i := len(arr) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
