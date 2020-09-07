package main

import "fmt"

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1,3},{3,0,1},{2},{0}}))
}
func canVisitAllRooms(rooms [][]int) bool {
	visitMap := make(map[int]bool, len(rooms))
	visitMap[0] = true
	begin := make([]int, len(rooms))
	copy(begin, rooms[0])
	return dfs(begin, visitMap, rooms)
}

func dfs(current []int, visitMap map[int]bool, rooms [][]int) bool {
	for _, next := range current {
		if visitMap[next] {
			continue
		}
		visitMap[next] = true
		dfs(rooms[next], visitMap, rooms)
	}
	for i := range rooms {
		if !visitMap[i] {
			return false
		}
	}
	return true
}
