package weeklycompetition

import (
	"fmt"
)

/************ 1st test************/
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := range ans {
		ans[i] = make([]int, n-2)
		for j := range ans[i] {
			ans[i][j] = getMax(i, j, grid)
		}
	}
	return ans
}

func getMax(x, y int, grid [][]int) (ans int) {
	ans = grid[x][y]
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if ans < grid[x+i][y+j] {
				ans = grid[x+i][y+j]
			}
		}
	}
	return
}
func Example_largest_local_values_in_a_matrix() {
	grid := [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}
	fmt.Println(largestLocal(grid))
	grid = [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}
	fmt.Println(largestLocal(grid))
	//Output:
	//[[9 9] [8 6]]
	//[[2 2 2] [2 2 2] [2 2 2]]
}

/************ 2nd test************/
func edgeScore(edges []int) (ans int) {
	scoreMap := make([]int, len(edges))
	for from, to := range edges {
		scoreMap[to] += from
		if scoreMap[to] > scoreMap[ans] || //取边积分最大,或
			(scoreMap[to] == scoreMap[ans] && to < ans) { //边积分相同取编号最小
			ans = to
		}
	}
	return ans
}
func Example_node_with_highest_edge_score() {
	edges := []int{1, 0, 0, 0, 0, 7, 7, 5}
	fmt.Println(edgeScore(edges))
	edges = []int{2, 0, 0, 2}
	fmt.Println(edgeScore(edges))
	edges = []int{3, 3, 3, 0}
	fmt.Println(edgeScore(edges))
	//Output:
	//7
	//0
	//0
}

/************ 3rd test************/
func smallestNumber(pattern string) (ans string) {
	used := make([]bool, 10)
	n := len(pattern) + 1
	str := make([]byte, n)
	ok := false

	var dfs func(int)
	dfs = func(idx int) {
		if ok {
			return
		}
		if idx > 1 {
			if pattern[idx-2] == 'I' && str[idx-1] < str[idx-2] {
				return
			} else if pattern[idx-2] == 'D' && str[idx-1] > str[idx-2] {
				return
			}
		}
		if idx == n {
			ans = string(str)
			ok = true
			return
		}
		for i := 1; i <= 9; i++ {
			if used[i] == true {
				continue
			}
			str[idx] = byte('0' + i)
			used[i] = true
			dfs(idx + 1)
			used[i] = false
			str[idx] = byte('0')
		}
	}
	dfs(0)
	return ans
}

func Example_construct_smallest_number_from_di_string() {
	pattern := "IIIDIDDD"
	fmt.Println(smallestNumber(pattern))
	pattern = "DDD"
	fmt.Println(smallestNumber(pattern))
	//Output:
	//123549876
	//4321
}

/************ 4th test************/
func countSpecialNumbers(n int) int {
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	n = len(digits)

	used := make([]int, 10)
	total := 0

	for i := 1; i < n; i++ {
		total += 9 * A(9, i-1)
	}

	for i := n - 1; i >= 0; i-- {
		num := digits[i]

		j := 0
		if i == n-1 {
			j = 1
		}
		for ; j < num; j++ {
			if used[j] != 0 {
				continue
			}
			total += A(10-(n-i), i)
		}

		used[num]++
		if used[num] > 1 {
			break
		}

		if i == 0 {
			total++
		}

	}
	return total
}
func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func A(m, n int) int {
	return fact(m) / fact(m-n)
}
func Example_count_special_integers() {
	n := 20
	fmt.Println(countSpecialNumbers(n))
	n = 5
	fmt.Println(countSpecialNumbers(n))
	n = 135
	fmt.Println(countSpecialNumbers(n))
	//Output:
	//19
	//5
	//110
}
