package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		curent := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if curent == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}
