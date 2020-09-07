package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}

func findSubsequences(nums []int) [][]int {
	return dfs(make([]int, 0, len(nums)), nums, map[string]struct{}{})
}

func dfs(current []int, arr []int, filter map[string]struct{}) [][]int {
	result := make([][]int, 0)
	if len(current) >= 2 {
		key := JoinString(current)
		if _, ok := filter[key]; !ok {
			filter[key] = struct{}{}
			n := make([]int, len(current))
			copy(n, current)
			result = append(result, n)
		}
	}
	for i := 0; i < len(arr); i++ {
		if len(current) == 0 || current[len(current)-1] <= arr[i] {
			result = append(result, dfs(append(current, arr[i]), arr[i+1:], filter)...)
		}
	}
	return result
}

func JoinString(arr []int) string {
	arrStr := make([]string, len(arr))
	for i, a := range arr {
		arrStr[i] = strconv.Itoa(a)
	}
	return strings.Join(arrStr, ",")
}
