package main

import "fmt"

func main() {
	fmt.Println(permute([]int{}))
	fmt.Println(permute([]int{1}))
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(len(permute([]int{1, 2, 3, 4})))
}

var (
	gnums  []int
	ans    [][]int
	mark   []bool
	before int
)

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	ansLen := 1
	for i := 2; i <= len(nums); i++ {
		ansLen *= i
	}
	if ansLen > before {
		ans = make([][]int, 0, ansLen)
	} else {
		ans = ans[:0]
	}
	before = ansLen
	gnums = nums
	mark = make([]bool, len(nums))
	dfs(make([]int, 0, len(gnums)))
	return ans
}

func dfs(current []int) {
	if len(current) == len(gnums) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		ans = append(ans, tmp)
		return
	}
	for i := 0; i < len(gnums); i++ {
		if mark[i] {
			continue
		}
		mark[i] = true
		dfs(append(current, gnums[i]))
		mark[i] = false
	}
}
