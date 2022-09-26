package missingtwolccihard

import "testing"

func missingTwo(nums []int) []int {
	//设缺失数为 a,b(a<b)
	n := len(nums) + 2
	sum := n * (n + 1) / 2
	for _, n := range nums {
		sum -= n
	}
	avg := sum / 2             // (a+b)/2 ==avg
	cur := avg * (1 + avg) / 2 //Sum (1~avg)
	for _, n := range nums {
		if n <= avg {
			cur -= n
		}
	} //cur==a
	return []int{cur, sum - cur}
}
func Test_missing_two_lcci(t *testing.T) {
	nums := []int{1}
	t.Log(missingTwo(nums))
	nums = []int{2, 3}
	t.Log(missingTwo(nums))
	nums = []int{3, 1, 2, 9, 8, 5, 7}
	t.Log(missingTwo(nums))
}
