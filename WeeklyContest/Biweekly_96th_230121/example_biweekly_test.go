package biweekly_contest

import (
	"container/heap"
	"sort"
	"testing"
)

/************ 1st test************/
func getCommon(nums1 []int, nums2 []int) int {
	mp := make(map[int]struct{}, 0)
	for _, n := range nums1 {
		mp[n] = struct{}{}
	}
	for _, n := range nums2 {
		if _, has := mp[n]; has {
			return n
		}
	}
	return -1
}
func Test_1st(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4}
	t.Log(getCommon(nums1, nums2))
	nums1 = []int{1, 2, 3, 6}
	nums2 = []int{2, 3, 4, 5}
	t.Log(getCommon(nums1, nums2))
}

/************ 2nd test************/
func minOperations(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pos, neg := 0, 0
	for i := 0; i < n; i++ {
		delta := nums1[i] - nums2[i]
		if delta > 0 {
			pos += delta
		} else {
			neg += delta
		}
		if k != 0 && delta%k != 0 {
			return -1
		} else if k == 0 && delta != 0 {
			return -1
		}
	}

	if pos+neg != 0 {
		return -1
	}

	if pos == neg {
		return 0
	}

	return int64(pos / k)
}

func Test_2nd(t *testing.T) {
	nums1 := []int{4, 3, 1, 4}
	nums2 := []int{1, 3, 7, 1}
	k := 3
	t.Log(minOperations(nums1, nums2, k))
	nums1 = []int{3, 8, 5, 2}
	nums2 = []int{2, 4, 1, 6}
	k = 1
	t.Log(minOperations(nums1, nums2, k))
	nums1 = []int{10, 5, 15, 20}
	nums2 = []int{20, 10, 15, 5}
	k = 0
	t.Log(minOperations(nums1, nums2, k))
	nums1 = []int{5, 10, 15, 20}
	nums2 = []int{5, 10, 15, 20}
	k = 0
	t.Log(minOperations(nums1, nums2, k))
}

/************ 3rd test************/
type node struct {
	x, y int
}

// 优先级队列实现
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) (ans int64) {
	n := len(nums1)
	arr := make([]node, n)
	for i := 0; i < n; i++ {
		arr[i] = node{nums1[i], nums2[i]}
	}
	//nums1与nums2绑定后按nums2从大到小排序
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].y > arr[j].y
	})

	//先把排序后 nums1 前 K 个元素加入小根堆，并计算它们的和
	sum := 0
	pq := &IntHeap{}
	for i := 0; i < k; i++ {
		sum += arr[i].x
		heap.Push(pq, arr[i].x)
	}
	//计算初始答案
	ans = int64(sum * arr[k-1].y)

	//枚举 nums2 中的最小值
	for i := k; i < n; i++ {
		//小根堆新加入一个 nums1 中的数,如果当前元素大于堆顶
		//就替换堆顶,这样可以让堆中元素之和变大
		sum += arr[i].x
		heap.Push(pq, arr[i].x)
		//要去掉最小的数，使得小根堆的大小等于 K
		cur := heap.Pop(pq).(int)
		sum -= cur
		//更新当前答案
		ans = max(ans, int64(arr[i].y*sum))
	}
	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func Test_3rd(t *testing.T) {
	nums1 := []int{1, 3, 3, 2}
	nums2 := []int{2, 1, 3, 4}
	k := 3
	t.Log(maxScore(nums1, nums2, k))
	nums1 = []int{4, 2, 3, 1, 1}
	nums2 = []int{7, 4, 10, 9, 6}
	k = 1
	t.Log(maxScore(nums1, nums2, k))
	nums1 = []int{2, 1, 14, 12}
	nums2 = []int{11, 7, 13, 6}
	k = 1
	t.Log(maxScore(nums1, nums2, k))
}

/************ 4th test************/
func isReachable(targetX int, targetY int) bool {
	g := gcd(targetX, targetY)
	for g%2 == 0 {
		g /= 2
	}
	return g == 1
}
func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
func Test_4th(t *testing.T) {
	tarX, tarY := 6, 9
	t.Log(isReachable(tarX, tarY))
	tarX, tarY = 4, 7
	t.Log(isReachable(tarX, tarY))
	tarX, tarY = 3, 7
	t.Log(isReachable(tarX, tarY))
}
