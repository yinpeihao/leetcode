package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	return dfs("", root)
}

func dfs(path string, current *TreeNode) []string {
	path = fmt.Sprintf("%s%d->", path, current.Val)

	//fmt.Println(path,current)
	if current.Left == nil && current.Right == nil {
		if len(path) != 0 {
			path = path[:len(path)-2]
		}
		//fmt.Println(path)
		return []string{path}
	}
	result := []string{}
	if current.Left != nil {
		result = append(result, dfs(path, current.Left)...)
	}
	if current.Right != nil {
		result = append(result, dfs(path, current.Right)...)
	}
	return result
}
