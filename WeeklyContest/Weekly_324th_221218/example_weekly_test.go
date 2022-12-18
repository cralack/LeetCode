package weekly_contest

import (
	"testing"
)

/************ 1st test************/
func similarPairs(words []string) (ans int) {
	cnt := make(map[int]int)
	for _, word := range words {
		mask := word2mask(word)
		//先把cnt[mask] 加到答案中
		ans += cnt[mask]
		//然后把mask 的出现次数加一
		cnt[mask]++
	}
	return
}

// 使用位运算将出现过的字母映射在数字mask上
func word2mask(word string) int {
	mask := 0
	for _, c := range word {
		mask |= 1 << (c - 'a')
	}
	return mask
}
func Test_1st(t *testing.T) {
	words := []string{"aba", "aabb", "abcd", "bac", "aabc"}
	t.Log(similarPairs(words))
	words = []string{"aabb", "ab", "ba"}
	t.Log(similarPairs(words))
	words = []string{"nba", "cba", "dba"}
	t.Log(similarPairs(words))

}

/************ 2nd test************/
func smallestValue(n int) int {
	notPrime := make([]bool, n+1)
	for i := 2; i*i <= n; i++ { //只需要遍历 [2,sqrt(n)]
		if !notPrime[i] {
			for j := i * i; j <= n; j += i { //from i*i
				notPrime[j] = true
			}
		}
	}
	factor := func(x int) int {
		sum := 0
		for i := 2; i*i <= n; i++ {
			for x%i == 0 {
				sum += i
				x /= i
			}
		}
		if x > 1 {
			sum += x
		}
		return sum
	}
	for notPrime[n] {
		sum := factor(n)
		if sum == n {
			return sum
		} else {
			n = sum
		}
	}
	return n
}

func Test_2nd(t *testing.T) {
	n := 15
	t.Log(smallestValue(n))
	n = 3
	t.Log(smallestValue(n))
	n = 4
	t.Log(smallestValue(n))
}

/************ 3rd test************/
func isPossible(n int, edges [][]int) bool {
	type pair struct{ from, to int }
	//from到to两点之间是否存在边
	has := map[pair]bool{}
	deg := make([]int, n+1)
	for _, e := range edges {
		from, to := e[0], e[1]
		has[pair{from, to}] = true
		has[pair{to, from}] = true
		deg[from]++
		deg[to]++
	}
	odd := []int{}
	for i, d := range deg {
		if d%2 > 0 {
			odd = append(odd, i)
		}
	}
	oddCnt := len(odd)
	if oddCnt == 0 {
		return true
	}
	if oddCnt == 2 {
		//奇点有2个:设这两个奇点为a和b。若a和b之间无连边，则把它们连起来；
		x, y := odd[0], odd[1]
		if !has[pair{x, y}] {
			return true
		}
		//否则我们需要找另一个点c分别连接a和b，枚举c即可
		for i := 1; i <= n; i++ {
			if i != x && i != y && !has[pair{i, x}] && !has[pair{i, y}] {
				return true
			}
		}
		return false
	}
	if oddCnt == 4 {
		//奇点有4个:只能用两条边连接这四个点才有可能符合要求。
		//因此枚举第一个点和哪个点连接，并检查是否合法即可
		a, b, c, d := odd[0], odd[1], odd[2], odd[3]
		return !has[pair{a, b}] && !has[pair{c, d}] ||
			!has[pair{a, c}] && !has[pair{b, d}] ||
			!has[pair{a, d}] && !has[pair{b, c}]
	}
	return false
}

func Test_3rd(t *testing.T) {
	edges, n := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 4}, {2, 5}}, 5
	t.Log(isPossible(n, edges))
	edges, n = [][]int{{1, 2}, {3, 4}}, 4
	t.Log(isPossible(n, edges))
	edges, n = [][]int{{1, 2}, {1, 3}, {1, 4}}, 4
	t.Log(isPossible(n, edges))

}

/************ 4th test************/
func cycleLengthQueries(n int, queries [][]int) (ans []int) {
	// 设 LCA 为a和b的最近公共祖先，那么环长等于dis(LCA->a)+dis(LCA->b)+1
	ans = make([]int, len(queries))
	for i, q := range queries {
		res := 1
		//不断循环，每次循环比较a和b的大小
		for a, b := q[0], q[1]; a != b; res++ {
			//如果a>b，说明a的深度大于等于b的深度
			if a > b {
				//那么把a移动到其父节点
				a /= 2
			} else { //否则相反
				b /= 2
			}
		}
		ans[i] = res
	}
	return
}
func Test_4th(t *testing.T) {
	queries, n := [][]int{{5, 3}, {4, 7}, {2, 3}}, 3
	t.Log(cycleLengthQueries(n, queries))
	queries, n = [][]int{{1, 2}}, 2
	t.Log(cycleLengthQueries(n, queries))
}
