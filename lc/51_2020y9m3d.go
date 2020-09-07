package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solveNQueens(7))
}

var Global [][]Node

type Node struct {
	x, y int
}

func solveNQueens(n int) [][]string {
	line := strings.Repeat(".", n)
	solve([]Node{}, 0, n)
	//fmt.Println(Global)
	ans := make([][]string, 0)
	for _, s := range Global {
		result := make([]string, n)
		for i := range result {
			result[i] = line
		}
		for _, c := range s {
			tmp := []byte(result[c.x])
			tmp[c.y] = 'Q'
			result[c.x] = string(tmp)
		}

		ans = append(ans, result)
	}
	Global = [][]Node{}

	return ans
}

func solve(list []Node, index int, n int) {
	//list = append(list, Node{nextX, nextY})
	//fmt.Println(nextX, nextY)
	if len(list) == n {
		s := make([]Node, n)
		copy(s, list)
		Global = append(Global, s)
		//fmt.Println("list:", list)
	}

	for i := index; i < n; i++ {
	out:
		for j := 0; j < n; j++ {
			for _, curent := range list {
				if curent.x == i || curent.y == j || (curent.x+curent.y) == (i+j) || (curent.x-i) == curent.y-j {
					//fmt.Println("fail:", i, j)
					continue out
				}
			}
			solve(append(list, Node{i, j}), i+1, n)
		}
	}
	return
}
