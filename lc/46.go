package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

var (
	ans  [][]int
	nums []int
)

func combinationSum2(candidates []int, target int) [][]int {
	nums = candidates
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans = ans[:0]
	dfs([]int{}, 0, target)
	return ans
}

func dfs(current []int, idx, target int) {
	if target == 0 {
		tmp := make([]int, len(current))
		copy(tmp, current)
		ans = append(ans, tmp)
	}
	if target < 0 {
		return
	}
	for i := idx; i < len(nums); i++ {
		//这道题的重点在于如何去重。 一种方案是Set，但是由于Set是红黑树底层，其效率过低。 一种巧妙的方案是：
		//
		//            if(i>start&&candidates[i]==candidates[i-1])
		//            {
		//                continue;
		//            }
		//这样可以避免相同的情况筛选两次（一次原生For循环，一次是递归）。
		//很巧妙，在同一递归层次只选择相同的数字一次
		if i != idx && nums[i] == nums[i-1] {
			continue
		}
		dfs(append(current, nums[i]), i+1, target-nums[i])
	}
}
