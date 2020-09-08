package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(4, 0))
}

var (
	gk, gn int
	mark   []bool
	result [][]int
)

func combine(n int, k int) [][]int {
	gk = k
	gn = n
	mark = make([]bool, n+1)
	result = result[:0]
	dfs77([]int{})
	return result
}

func dfs77(current []int) {
	if len(current) == gk {
		t := make([]int, len(current))
		copy(t, current)
		result = append(result, t)
		return
	}
	var begin = 1
	if len(current) != 0 {
		begin = current[len(current)-1]
	}
	for i := begin; i <= gn; i++ {
		if mark[i] {
			continue
		} else {
			mark[i] = true
			dfs77(append(current, i))
			mark[i] = false
		}
	}
}
