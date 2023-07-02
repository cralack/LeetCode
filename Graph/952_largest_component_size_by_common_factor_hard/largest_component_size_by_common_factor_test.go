package largestcomponentsizebycommonfactorhard

import (
	"testing"
)

// type UF struct {
// 	parent []int
// 	primes []int
// }

// var n int = 1e5 + 7

// // var n = 70

// func Constructor(nums []int) UF {
// 	uf := UF{
// 		parent: make([]int, n),
// 		primes: getPrimeArr(n)}

// 	for i := 0; i < n; i++ {
// 		uf.parent[i] = i
// 	}
// 	return uf
// }
// func (uf *UF) find(x int) int {
// 	if uf.parent[x] != x {
// 		uf.parent[x] = uf.find(uf.parent[x])
// 	}
// 	return uf.parent[x]
// }
// func (uf *UF) union(p, q int) {
// 	rootP := uf.find(p)
// 	rootQ := uf.find(q)
// 	if rootP != rootQ {
// 		uf.parent[rootP] = rootQ
// 	}
// }

// func getPrimeArr(n int) []int {
// 	isPrime := make([]bool, n+1)
// 	arr := make([]int, n+1)
// 	for i := 2; i <= n; i++ {
// 		isPrime[i] = true
// 	}
// 	cnt := 0
// 	for i := 2; i <= n; i++ {
// 		if isPrime[i] {
// 			arr[cnt] = i
// 			cnt++
// 		}
// 		for j := 0; arr[j]*i <= n; j++ {
// 			isPrime[arr[j]*i] = false
// 			if i%arr[j] == 0 {
// 				break
// 			}
// 		}
// 	}
// 	return arr[:cnt]
// }

// func largestComponentSize(nums []int) (ans int) {
// 	uf := Constructor(nums)

//		for _, num := range nums {
//			quotient := num
//			k := len(uf.primes)
//			for j := 0; j < k && uf.primes[j]*uf.primes[j] <= quotient; j++ {
//				if quotient%uf.primes[j] == 0 {
//					uf.union(num, uf.primes[j])
//				}
//				for quotient%uf.primes[j] == 0 {
//					quotient /= uf.primes[j]
//				}
//			}
//			if quotient > 1 {
//				uf.union(quotient, num)
//			}
//		}
//		cnt := make([]int, n)
//		ans = 0
//		for _, num := range nums {
//			cnt[uf.find(num)]++
//			if ans < cnt[uf.find(num)] {
//				ans = cnt[uf.find(num)]
//			}
//		}
//		return
//	}
const mx int = 80

var pf [mx + 1][]int32

func init() {
	for i := 2; i <= mx; i++ {
		if pf[i] == nil {
			for j := i; j <= mx; j += i {
				pf[j] = append(pf[j], int32(i))
			}
		}
	}
}

var fa [mx + 1]int32

func largestComponentSize(nums []int) int {
	for i := range fa {
		fa[i] = int32(i)
	}
	var find func(int32) int32
	find = func(x int32) int32 {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	for _, v := range nums {
		for _, p := range pf[v] {
			fa[find(p)] = find(int32(v))
		}
	}
	cnt := [mx + 1]int16{}
	ans := int16(0)
	for _, v := range nums {
		v := find(int32(v))
		cnt[v]++
		if cnt[v] > ans {
			ans = cnt[v]
		}
	}
	return int(ans)
}
func Test_largest_component_size_by_common_factor(t *testing.T) {
	// t.Log(getPrimeArr(30))
	nums := []int{4, 6, 15, 35}
	t.Log(largestComponentSize(nums))
	nums = []int{20, 50, 9, 63}
	t.Log(largestComponentSize(nums))
	nums = []int{2, 3, 6, 7, 4, 12, 21, 39}
	t.Log(largestComponentSize(nums))
}
