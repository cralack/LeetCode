package numberofordersinthebacklog

import (
	"container/heap"
	"testing"
)

type order struct {
	price, amount int
}
type hp []order

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(order)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getNumberOfBacklogOrders(orders [][]int) (ans int) {
	//买单价需>=售单价 才能匹配一对订单
	sell := hp{} //小根堆
	buy := hp{}  //大根堆
	for _, cur := range orders {
		//curOrder
		curPrice, curAmount, curType := cur[0], cur[1], cur[2]
		//buy order
		if curType == 0 {
			//尽量匹配卖单价有<=当前买单价的订单
			for curAmount > 0 && len(sell) > 0 && sell[0].price <= curPrice {
				//卖价最小的单
				pair := heap.Pop(&sell).(order)
				//买单量足以匹配卖单
				if curAmount >= pair.amount {
					curAmount -= pair.amount
					//买单量不足以匹配卖单
				} else {
					//卖单抵扣数量后重新压回队列
					heap.Push(&sell, order{
						pair.price,
						pair.amount - curAmount})
					//买单量清零
					curAmount = 0
				}
			}
			//买单量有剩余
			if curAmount > 0 {
				heap.Push(&buy, order{-curPrice, curAmount})
			}
			//sell order
		} else {
			for curAmount > 0 && len(buy) > 0 && -buy[0].price >= curPrice {
				pair := heap.Pop(&buy).(order)
				if curAmount >= pair.amount {
					curAmount -= pair.amount
				} else {
					heap.Push(&buy, order{
						pair.price,
						pair.amount - curAmount})
					curAmount = 0
				}
			}
			if curAmount > 0 {
				heap.Push(&sell, order{curPrice, curAmount})
			}
		}
	}
	const mod int = 1e9 + 7
	//买单盈余
	for len(buy) > 0 {
		ans += heap.Pop(&buy).(order).amount
	}
	for len(sell) > 0 {
		ans += heap.Pop(&sell).(order).amount
	}
	return ans % mod
}
func Test_number_of_orders_in_the_backlog(t *testing.T) {
	orders := [][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}
	t.Log(getNumberOfBacklogOrders(orders))
	orders = [][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}}
	t.Log(getNumberOfBacklogOrders(orders))
}
