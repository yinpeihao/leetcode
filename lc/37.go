package main

var (
	lineMap map[int]map[byte]bool
	rowMap  map[int]map[byte]bool
	spanMap map[int]map[byte]bool
	gboard  [][]byte
)

func solveSudoku(board [][]byte) {
	gboard = board
	lineMap = make(map[int]map[byte]bool)
	rowMap = make(map[int]map[byte]bool)
	spanMap = make(map[int]map[byte]bool)
	for i := 0; i < 9; i++ {
		lineMap[i] = make(map[byte]bool)
		rowMap[i] = make(map[byte]bool)
		spanMap[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != ' ' {
				lineMap[i][board[i][j]] = true
				rowMap[j][board[i][j]] = true
				spanMap[i/3*3+j/3][board[i][j]] = true
			}
		}
	}
	dfs(0, 0)
}

func dfs(x, y int) bool {
	if x == 9 || y == 9 {
		return true
	}
	nextX, nextY := x, y+1
	if y+1 == 9 {
		nextX += 1
		nextY = 0
	}
	if gboard[x][y] != '.' {
		return dfs(nextX, nextY)
	}
	for i := byte(1); i <= 9; i++ {
		cur := i + '0'
		spanIndex := x/3*3 + y/3
		if !lineMap[x][cur] && !rowMap[y][cur] && !spanMap[spanIndex][cur] {
			gboard[x][y] = cur
			lineMap[x][cur] = true
			rowMap[y][cur] = true
			spanMap[spanIndex][cur] = true
			if dfs(nextX, nextY) {
				return true
			}
			gboard[x][y] = '.'
			lineMap[x][cur] = false
			rowMap[y][cur] = false
			spanMap[spanIndex][cur] = false
		}
	}
	return false
}
