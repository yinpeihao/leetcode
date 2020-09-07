package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	list1 := list.New()
	list1.PushBack(root)
	for list1.Len() != 0 {
		curArr := make([]int, 0)
		levelLen := list1.Len()
		for levelLen != 0 {
			levelLen--
			cur := list1.Front()
			list1.Remove(cur)
			curArr = append(curArr, cur.Value.(*TreeNode).Val)
			if cur.Value.(*TreeNode).Left != nil {
				list1.PushBack(cur.Value.(*TreeNode).Left)
			}
			if cur.Value.(*TreeNode).Right != nil {
				list1.PushBack(cur.Value.(*TreeNode).Right)
			}
		}
		result = append(result, curArr)
	}
	reverse := make([][]int, len(result))
	for i := 0; i < len(result); i++ {
		reverse[i] = result[len(result)-1-i]
	}
	return reverse
}
