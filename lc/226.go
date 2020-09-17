package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	var dfs func(node *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		cur.Left, cur.Right = cur.Right, cur.Left
		dfs(cur.Left)
		dfs(cur.Right)
	}
	dfs(root)
	return root
}
