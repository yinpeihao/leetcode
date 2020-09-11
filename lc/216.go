package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
}

var (
	ans   [][]int
	gk    int
)

func combinationSum3(k int, n int) [][]int {
	gk = k
	ans = ans[:0]
	dfs(make([]int, 0, n), n, 1)
	return ans
}

func dfs(current []int, num int, idx int) {
	if num == 0 && len(current) == gk {
		tmp := make([]int, len(current))
		copy(tmp, current)
		ans = append(ans, tmp)
		return
	}
	if num <= 0 || len(current) >= gk {
		return
	}
	for i := idx; i <= 9; i++ {
		dfs(append(current, i), num-i, i+1)
	}
}
