package tfgame

import (
	"sort"
	"testing"
)

const syms = "+-*/"

func permuteUnique(cards []int) (res [][]int) {
	// 先排序，让相同的元素靠在一起
	sort.Ints(cards)
	path := make([]int, 0, 4)
	used := make([]bool, len(cards))
	var dfs func()
	dfs = func() {
		//base case
		if len(path) == len(cards) {
			tmp := make([]int, len(cards))
			copy(tmp, path)
			res = append(res, tmp)
			//fmt.Println(tmp)
			return
		}
		for i := 0; i < len(cards); i++ {
			if used[i] {
				continue
			}
			// 剪枝逻辑，固定相同的元素在排列中的相对位置
			if i > 0 && cards[i] == cards[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, cards[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}
func judgePoint24(cards []int) bool {
	per_cards := permuteUnique(cards)
	path := make([]byte, 0, 7)
	vis := make([]bool, 4)
	flag := false

	var dfs func(num int, nums []int)
	dfs = func(num int, nums []int) {
		if num == 24 && len(path) == 7 {

			//fmt.Println(string(path))

			flag = true
			return
		} else if flag || len(path) == 7 {
			return
		}
		for i := 0; i < len(nums); i++ {
			if !vis[i] {
				for j := 0; j < len(syms); j++ {
					switch syms[j] {
					case '+':
						vis[i] = true
						path = append(path, '+')
						path = append(path, byteitoa(nums[i]))
						dfs(num+nums[i], nums)
						vis[i] = false
						path = path[:len(path)-2]
					case '-':
						vis[i] = true
						path = append(path, '-')
						path = append(path, byteitoa(nums[i]))
						dfs(num-nums[i], nums)
						vis[i] = false
						path = path[:len(path)-2]
					case '*':
						vis[i] = true
						path = append(path, '*')
						path = append(path, byteitoa(nums[i]))
						dfs(num*nums[i], nums)
						vis[i] = false
						path = path[:len(path)-2]
					case '/':
						if num%nums[i] == 0 {
							vis[i] = true
							path = append(path, '/')
							path = append(path, byteitoa(nums[i]))
							dfs(num/nums[i], nums)
							vis[i] = false
							path = path[:len(path)-2]
						}
					}
				}
			}
		}
	}

	for i := range per_cards {
		numOne := per_cards[i][0]
		nums := make([]int, 0, 3)
		nums = append(nums, per_cards[i][1:]...)
		path = append(path, byteitoa(numOne))
		dfs(numOne, nums)
		path = path[:len(path)-1]
	}

	return flag
}
func byteitoa(n int) byte {
	switch n {
	case 0:
		return '0'
	case 1:
		return '1'
	case 2:
		return '2'
	case 3:
		return '3'
	case 4:
		return '4'
	case 5:
		return '5'
	case 6:
		return '6'
	case 7:
		return '7'
	case 8:
		return '8'
	case 9:
		return '9'
	}
	return '#'
}

func Test_tf_game(t *testing.T) {
	cards := []int{1, 1, 6, 8}
	t.Log(judgePoint24(cards))
}
