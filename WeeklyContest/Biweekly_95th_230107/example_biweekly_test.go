package biweekly_contest

import (
	"sort"
	"testing"
)

/************ 1st test************/
func categorizeBox(length int, width int, height int, mass int) string {
	var bulky, heavy bool
	if mass >= 100 {
		heavy = true
	}
	v := 1
	for _, x := range []int{length, width, height} {
		if x >= 1e4 {
			bulky = true
			break
		} else {
			v *= x
			if v >= 1e9 {
				bulky = true
				break
			}
		}
	}
	if bulky && heavy {
		return "Both"
	} else if bulky {
		return "Bulky"
	} else if heavy {
		return "Heavy"
	}
	return "Neither"
}
func Test_1st(t *testing.T) {
	length := 1000
	width := 35
	height := 700
	mass := 300
	t.Log(categorizeBox(length, width, height, mass))
	length = 200
	width = 50
	height = 800
	mass = 5
	t.Log(categorizeBox(length, width, height, mass))
	length = 2000
	width = 2000
	height = 2000
	mass = 5
	t.Log(categorizeBox(length, width, height, mass))
	length = 2000
	width = 2000
	height = 2000
	mass = 900
	t.Log(categorizeBox(length, width, height, mass))
}

/************ 2nd test************/
type DataStream struct {
	value, k int
	que      []int
	idx      int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		value: value,
		k:     k,
		que:   make([]int, k),
		idx:   0,
	}
}

func (this *DataStream) Consec(num int) bool {
	this.que[this.idx] = num
	this.idx = (this.idx + 1) % this.k
	for _, n := range this.que {
		if n != this.value {
			return false
		}
	}
	return true
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */
func Test_2nd(t *testing.T) {
	sol := Constructor(4, 3)
	t.Log(sol.Consec(4))
	t.Log(sol.Consec(4))
	t.Log(sol.Consec(4))
	t.Log(sol.Consec(3))
}

/************ 3rd test************/
func xorBeauty(nums []int) (ans int) {
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return
}
func Test_3rd(t *testing.T) {
	nums := []int{1, 4}
	t.Log(xorBeauty(nums))
	nums = []int{15, 45, 20, 2, 34, 35, 5, 44, 32, 30}
	t.Log(xorBeauty(nums))
}

/************ 4th test************/
func maxPower(stations []int, r int, k int) (ans int64) {
	n := len(stations)
	check := func(lim int64) bool {
		arr := make([]int64, n)
		for i, x := range stations {
			arr[i] = int64(x)
		}
		// 初始化滑动窗口的和
		var sum int64 = 0
		for i := 0; i <= r && i < n; i++ {
			sum += arr[i]
		}
		// 表示还有几个供电站可以新建
		var rem int64 = int64(k)
		// 从左到右计算每个城市的电量，同时维护滑动窗口 [l, r]
		for i, left, right := 0, 0, r; ; i++ {
			if sum < lim {
				// 当前城市电量不足
				delta := lim - sum
				// 新供电站不够，返回 false
				if delta > rem {
					return false
				}
				// 新供电站足够，建在滑动窗口最右边
				rem -= delta
				arr[right] += delta
				sum += delta
			}
			if i == n-1 {
				break
			}
			// 滑动窗口向前移动一个城市
			if i >= r {
				sum -= arr[left]
				left++
			}
			if right != n-1 {
				sum += arr[right+1]
				right++
			}
		}
		return true
	}
	// 二分答案
	left, right := 0, int(2e10+1)
	for left < right {
		mid := left + (right-left)>>1
		if check(int64(mid)) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return int64(left) - 1
	// return int64(sort.Search(1e10+1, func(x int) bool {
	// 	x++
	// 	return !check(int64(x))
	// }))
}

// https://space.bilibili.com/206214
func maxPower_v2(stations []int, r int, k int) int64 {
	n := len(stations)
	sum := make([]int, n+1) // 前缀和
	for i, x := range stations {
		sum[i+1] = sum[i] + x
	}
	return int64(sort.Search(sum[n]+k, func(minPower int) bool {
		minPower++               // 改为二分最小的不满足要求的值，这样 sort.Search 返回的就是最大的满足要求的值
		diff := make([]int, n+1) // 差分数组
		sumD, need := 0, 0
		for i, d := range diff[:n] {
			sumD += d
			power := sum[min(i+r+1, n)] - sum[max(i-r, 0)]
			m := minPower - power - sumD
			if m > 0 { // 需要 m 个供电站
				need += m
				if need > k {
					return true // 不满足要求
				}
				sumD += m                  // 差分更新
				diff[min(i+r*2+1, n)] -= m // 差分更新
			}
		}
		return false
	}))
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func Test_4th(t *testing.T) {
	stations := []int{1, 2, 4, 5, 0}
	r, k := 1, 2
	t.Log(maxPower(stations, r, k))
	stations = []int{4, 4, 4, 4}
	r, k = 0, 3
	t.Log(maxPower(stations, r, k))
}
