package minimumgeneticmutation

import (
	"testing"
)

func minMutation(start string, end string, bank []string) int {
	valid := make(map[string]bool, 0)
	visited := make(map[string]bool, 0)
	for _, s := range bank {
		valid[s] = true
	}
	if _, ok := valid[end]; !ok {
		return -1
	}
	step := 0
	q1, q2 := make(map[string]bool, 0), make(map[string]bool, 0)
	q1[start] = true
	q2[end] = true
	for len(q1) > 0 && len(q2) > 0 {
		// 优先操作短队列
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}
		// tmp保存接下来要扩展的节点
		tmp := make(map[string]bool, 0)
		// 遍历q1
		for cur := range q1 {
			if visited[cur] == true {
				continue
			}
			visited[cur] = true
			// 两边有交集则停止
			if q2[cur] == true {
				return step
			}
			// 当前dna能突变的范围
			next := make([]string, 0, 3*8)
			for i := 0; i < 8; i++ {
				next = append(next, mutate(cur, i)...)
			}
			for _, mutant := range next {
				if _, ok := valid[mutant]; !ok {
					continue
				} else {
					tmp[mutant] = true
				}
			}
		}
		step++
		q1 = q2
		q2 = tmp
	}
	return -1
}

// 枚举tar上idx位置发生突变会出现的dna
func mutate(tar string, idx int) []string {
	res := make([]string, 0, 3)
	c := tar[idx]
	arr := []byte{'A', 'C', 'G', 'T'}
	for _, char := range arr {
		if c == char { // 跳过tar
			continue
		} else {
			tmp := make([]byte, len(tar))
			copy(tmp, []byte(tar))
			tmp[idx] = char
			res = append(res, string(tmp))
		}
	}
	return res
}
func Test_minimum_genetic_mutation(t *testing.T) {
	start := "AACCGGTT"
	end := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	t.Log(minMutation(start, end, bank))

	start = "AACCGGTT"
	end = "AAACGGTA"
	bank = []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	t.Log(minMutation(start, end, bank))

	start = "AAAAACCC"
	end = "AACCCCCC"
	bank = []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
	t.Log(minMutation(start, end, bank))

	start = "AAAAAAAA"
	end = "AAGTAAAA"
	bank = []string{"CAAAAAAA", "CCAAAAAA", "CCATAAAA", "CCGTAAAA", "CAGTAAAA", "AAGTAAAA"}
	t.Log(minMutation(start, end, bank))

	start = "AAAAAAAA"
	end = "CCCCCCCC"
	bank = []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC",
		"AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA"}
	t.Log(minMutation(start, end, bank))

	start = "AAAAAAAA"
	end = "CCCCCCCC"
	bank = []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC",
		"AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA", "CCCCCCCC"}
	t.Log(minMutation(start, end, bank))
}
