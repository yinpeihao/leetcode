package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	for i := range nums {
		res = append(res, dfs([]int{}, nums[i:])...)
	}
	return res
}

func dfs(current []int, arr []int) [][]int {
	result := make([][]int, 0)
	if len(arr) == 0 {
		if len(current) >= 2 {
			n := make([]int, len(current))
			copy(n, current)
			result = append(result, n)
		}
		return result
	}
	current = append(current, arr[0])
	for i := 1; i <= len(arr); i++ {
		if i < len(arr) && current[len(current)-1] == arr[i] {
			result = append(result, dfs(current, arr[i:])...)
			break
		} else if i == len(arr) || current[len(current)-1] < arr[i] {
			result = append(result, dfs(current, arr[i:])...)
		}
	}
	current = current[:len(current)-1]
	return result
}

