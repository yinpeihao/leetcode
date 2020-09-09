package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	result = result[:0]
	dfs(candidates, []int{}, target, 0)
	return result
}

func dfs(list []int, current []int, target int, index int) {
	if target == 0 {
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(list); i++ {
		dfs(list, append(current, list[i]), target-list[i], i)
	}
}
