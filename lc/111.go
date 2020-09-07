package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 1)
}

func bfs(cur *TreeNode) int {
	list := make([]*TreeNode, 0)
	var size, level int
	list = append(list, cur)
	for len(list) != 0 {
		level++
		size = len(list)
		for size != 0 {
			size--
			begin := list[0]
			list = list[1:]
			if begin.Left == nil && begin.Right == nil {
				return level
			} else if begin.Left == nil && begin.Right != nil {
				list = append(list, begin.Right)
			} else if begin.Right == nil && begin.Left != nil {
				list = append(list, begin.Left)
			} else {
				list = append(list, begin.Left)
				list = append(list, begin.Right)
			}
		}
	}
	return level
}

func dfs(cur *TreeNode, dep int) int {
	if cur.Right == nil && cur.Left == nil {
		return dep
	}
	if cur.Left == nil && cur.Right != nil {
		return dfs(cur.Right, dep+1)
	} else if cur.Left != nil && cur.Right == nil {
		return dfs(cur.Left, dep+1)
	}
	return minInt(dfs(cur.Left, dep+1), dfs(cur.Right, dep+1))
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
