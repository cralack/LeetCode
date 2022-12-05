package deliveringboxesfromstoragetoports

import "testing"

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)
	//箱子重量前缀和
	weiSum := make([]int, n+1)
	//码头间行程数
	cost := make([]int, n)
	for i, box := range boxes {
		tarPort, weight := box[0], box[1]
		weiSum[i+1] = weiSum[i] + weight
		if i < n-1 {
			t := 0
			if tarPort != boxes[i+1][0] {
				t++
			}
			cost[i+1] = cost[i] + t
		}
	}
	//前 i 个箱子运送到位的最小行程数
	dp := make([]int, n+1)
	// { //O(n^2)
	// 	for i := 1; i <= n; i++ {
	// 		dp[i] = 1 << 30
	// 		//在 [i-maxBoxes,..i-1]这个窗口内找到一个 j，使得 dp[j] - cost[j]的值最小
	// 		for j := max(0, i-maxBoxes); j < i; j++ {
	// 			if weiSum[i]-weiSum[j] <= maxWeight {
	// 				dp[i] = min(dp[i], dp[j]+cost[i-1]-cost[j]+2)
	// 			}
	// 		}
	// 	}
	// }
	//箱子idx 单调栈
	que := &Queue{0}
	for i := 1; i <= n; i++ {
		//队列内箱子数或箱子总重超限
		for !que.IsEmpty() && (i-que.Head() > maxBoxes || weiSum[i]-weiSum[que.Head()] > maxWeight) {
			que.PopHead()
		}
		if !que.IsEmpty() {
			dp[i] = cost[i-1] + dp[que.Head()] - cost[que.Head()] + 2
		}
		if i < n {
			for !que.IsEmpty() && dp[que.Tail()]-cost[que.Tail()] >= dp[i]-cost[i] {
				que.PopTail()
			}
			que.Push(i)
		}
	}
	return dp[n]
}

type Queue []int

func (q Queue) Head() int      { return q[0] }
func (q Queue) Tail() int      { return q[len(q)-1] }
func (q *Queue) Push(x int)    { *q = append(*q, x) }
func (q *Queue) IsEmpty() bool { return len(*q) == 0 }
func (q *Queue) PopHead() int {
	head := q.Head()
	old := *q
	*q = old[1:]
	return head
}
func (q *Queue) PopTail() int {
	tail := q.Tail()
	old := *q
	*q = old[:len(*q)-1]
	return tail
}

func Test_delivering_boxes_from_storage_to_ports(t *testing.T) {
	boxes := [][]int{{1, 1}, {2, 1}, {1, 1}}
	portsCount, maxBoxes, maxWeight := 2, 3, 3
	t.Log(boxDelivering(boxes, portsCount, maxBoxes, maxWeight))

	boxes = [][]int{{1, 2}, {3, 3}, {3, 1}, {3, 1}, {2, 4}}
	portsCount, maxBoxes, maxWeight = 3, 3, 6
	t.Log(boxDelivering(boxes, portsCount, maxBoxes, maxWeight))

	boxes = [][]int{{1, 2}, {3, 3}, {3, 1}, {3, 1}, {2, 4}}
	portsCount, maxBoxes, maxWeight = 3, 3, 6
	t.Log(boxDelivering(boxes, portsCount, maxBoxes, maxWeight))

	boxes = [][]int{{1, 4}, {1, 2}, {2, 1}, {2, 1}, {3, 2}, {3, 4}}
	portsCount, maxBoxes, maxWeight = 3, 6, 7
	t.Log(boxDelivering(boxes, portsCount, maxBoxes, maxWeight))

	boxes = [][]int{{2, 4}, {2, 5}, {3, 1}, {3, 2}, {3, 7}, {3, 1}, {4, 4}, {1, 3}, {5, 2}}
	portsCount, maxBoxes, maxWeight = 5, 5, 7
	t.Log(boxDelivering(boxes, portsCount, maxBoxes, maxWeight))
}
