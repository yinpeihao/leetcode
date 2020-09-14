package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法2：循环
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	ans := make([]int, 0)
	cur := root
	for len(stack) != 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}

//方法1：dfs
//func inorderTraversal(root *TreeNode) []int {
//	ans = ans[:0]
//	dfs(root)
//	return ans
//}
//
//func dfs(cur *TreeNode) {
//	if cur == nil {
//		return
//	}
//	dfs(cur.Left)
//	ans = append(ans, cur.Val)
//	dfs(cur.Right)
//}
