package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int, len(nums))
	for _, num := range nums {
		mp[num] ++
	}
	type Node struct {
		val       int
		frequency int
	}
	n := make([]Node, 0)
	for k, v := range mp {
		n = append(n, Node{
			val:       k,
			frequency: v,
		})
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i].frequency > n[j].frequency
	})
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = n[i].val
	}
	return result
}
