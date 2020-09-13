package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}

var (
	gword  string
	gboard [][]byte
	mark   [][]bool
	ans    bool
)

func exist(board [][]byte, word string) bool {
	gword = word
	gboard = board
	ans = false
	mark = make([][]bool, len(board))
	for i := range mark {
		mark[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != word[0] {
				continue
			}
			mark[i][j] = true
			dfs(0, i, j)
			mark[i][j] = false
			if ans {
				return true
			}
		}
	}

	return false
}

func dfs(idx, x, y int) {
	if idx == len(gword)-1 {
		ans = true
		return
	}
	//四个方向走
	//上
	if x-1 >= 0 && !mark[x-1][y] && gboard[x-1][y] == gword[idx+1] {
		mark[x-1][y] = true
		dfs(idx+1, x-1, y)
		mark[x-1][y] = false
	}
	if ans {
		return
	}
	//右
	if y+1 < len(gboard[x]) && !mark[x][y+1] && gboard[x][y+1] == gword[idx+1] {
		mark[x][y+1] = true
		dfs(idx+1, x, y+1)
		mark[x][y+1] = false
	}
	if ans {
		return
	}
	//下
	if x+1 < len(gboard) && !mark[x+1][y] && gboard[x+1][y] == gword[idx+1] {
		mark[x+1][y] = true
		dfs(idx+1, x+1, y)
		mark[x+1][y] = false
	}
	if ans {
		return
	}
	//左
	if y-1 >= 0 && !mark[x][y-1] && gboard[x][y-1] == gword[idx+1] {
		mark[x][y-1] = true
		dfs(idx+1, x, y-1)
		mark[x][y-1] = false
	}
	if ans {
		return
	}
}
